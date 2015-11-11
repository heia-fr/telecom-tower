// Copyright 2015 Jacques Supcik, HEIA-FR
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Test programm for Telecom Tower
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/tower"
	"html"
	"log"
	"net/http"
)

type Line struct {
	Text  string
	Font  int
	Color string
}

type TextMessage struct {
	Body     []Line `json:"body"`
	EntrySep Line   `json:"entry-sep"`
	ExitSep  Line   `json:"exit-sep"`
	LoopSep  Line   `json:"loop-sep"`
}

type BitmapMessage struct {
	matrix     *ledmatrix.Matrix
	preamble   int
	checkpoint int
}

// textMsgQueue is the input queue for the displayBuilder goroutine.
var textMsgQueue chan TextMessage

// bitmapMsgQueue is the input queue for the towerServer goroutine.
var bitmapMsgQueue chan BitmapMessage

// renderdeTextMessage convert a text message to a LED bitmap matrix.
// This method is used by the displayBuilder gorouting.
func renderedTextMessage(message []Line, bg ledmatrix.Color) *ledmatrix.Matrix {
	matrix := ledmatrix.NewMatrix(tower.Rows, 0)
	writer := ledmatrix.NewWriter(matrix)
	writer.Spacer(matrix.Columns, 0) // Blank bootstrap

	for _, line := range message {
		var r, g, b int
		var f bitmapfont.Font

		if line.Font < 8 {
			f = bitmapfont.F68
		} else {
			f = bitmapfont.F88
		}

		fmt.Sscanf(line.Color, "#%02x%02x%02x", &r, &g, &b)
		writer.WriteText(line.Text, f, ledmatrix.RGB(r, g, b), bg)
	}
	return writer.Matrix
}

func blank(len int, bg ledmatrix.Color) *ledmatrix.Matrix {
	writer := ledmatrix.NewWriter(ledmatrix.NewMatrix(tower.Rows, tower.Columns))
	writer.Spacer(tower.Columns, bg) // Blank spacer
	return writer.Matrix
}

func displayBuilder() {
	bg := ledmatrix.RGB(0, 0, 0)
	preamble := blank(tower.Columns, bg)

	for {
		message := <-textMsgQueue
		body := renderedTextMessage(message.Body, bg)
		entrySep := renderedTextMessage([]Line{message.EntrySep}, bg)
		exitSep := renderedTextMessage([]Line{message.ExitSep}, bg)
		loopSep := renderedTextMessage([]Line{message.LoopSep}, bg)

		if body.Columns == 0 {
			log.Println("Skip empty message")
			continue
		}

		for body.Columns < tower.Columns {
			body.Append(loopSep, renderedTextMessage(message.Body, bg))
		}

		preamble.Append(entrySep)

		wrapper := ledmatrix.Concat(loopSep, body.Slice(0, tower.Columns))

		bitmapMsgQueue <- BitmapMessage{
			matrix:     ledmatrix.Concat(preamble, body, wrapper),
			preamble:   preamble.Columns,
			checkpoint: preamble.Columns + body.Columns - tower.Columns,
		}
		preamble = ledmatrix.Concat(body.Slice(body.Columns-tower.Columns, body.Columns), exitSep)
	}
}

func towerServer() {
	currentMessage := BitmapMessage{
		matrix:     blank(tower.Columns, 0),
		preamble:   0,
		checkpoint: 0,
	}

	start := 0
	for {
		// Roll until checkpoint!
		for i := start; i < currentMessage.checkpoint; i++ {
			tower.DisplayQueue <- currentMessage.matrix.BitmapSliceAt(i, tower.Columns)
		}
		// Checkpoint
		select {
		case m := <-bitmapMsgQueue:
			log.Println("New Message...")
			currentMessage = m
			start = 0

		default:
			// continue
			for i := currentMessage.checkpoint; i < currentMessage.matrix.Columns-tower.Columns; i++ {
				tower.DisplayQueue <- currentMessage.matrix.BitmapSliceAt(i, tower.Columns)
			}
			// skip preamble for the next run
			start = currentMessage.preamble
		}

	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	var message TextMessage
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&message); err != nil {
		fmt.Fprintf(w, "err\n")
	} else {
		textMsgQueue <- message
		fmt.Fprintf(w, "OK\n")
	}

}

func main() {
	var jsonFile string
	flag.StringVar(&jsonFile, "message", "", "JSON File with data")
	flag.Parse()

	log.Println("Booting tower...")
	tower.Init(32)

	textMsgQueue = make(chan TextMessage, 32)
	bitmapMsgQueue = make(chan BitmapMessage, 32)

	go towerServer()
	go displayBuilder()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.Methods("POST").Path("/message").HandlerFunc(PostMessage)
	log.Fatal(http.ListenAndServe(":8080", router))
}
