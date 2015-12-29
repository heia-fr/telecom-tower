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

// 2015-07-29 | JS | First version
// 2015-11-18 | JS | Latest version

//
// Telecom Tower REST server
//
package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/sprite"
	"github.com/heia-fr/telecom-tower/tower"
	"log"
	"net/http"
)

const (
	TEXT_MSG_QUEUE_SIZE   = 32
	BITMAP_MSG_QUEUE_SIZE = 32
	DEFAULT_PORT          = 8080
	DEFAULT_BRIGHTNESS    = 32
)

type Line struct {
	Text  string
	Font  int
	Color string
}

type TextMessage struct {
	Body         []Line `json:"body"`
	Introduction []Line `json:"introduction"`
	Conclusion   []Line `json:"conclusion"`
	Separator    []Line `json:"separator"`
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

// blank generates a "space" from the given color
func blank(len int, bg ledmatrix.Color) *ledmatrix.Matrix {
	writer := ledmatrix.NewWriter(ledmatrix.NewMatrix(tower.Rows, tower.Columns))
	writer.Spacer(tower.Columns, bg) // Blank spacer
	return writer.Matrix
}

// towerServer is a the goroutine that receives bitmap messages from the displayBuilder
// and dispatch "frames" to the tower LEDs. The preamble is only sent once; at the
// checkpoint, the goroutine checks if a new message is available; if yes, it switches
// to this new message; if no, it finish the message and roll th same message again.
func towerServer() {
	currentMessage := BitmapMessage{
		matrix:     blank(tower.Columns, 0),
		preamble:   0,
		checkpoint: 0,
	}

	start := 0
	for { // Loop forever
		// Roll until checkpoint!
		for i := start; i < currentMessage.checkpoint; i++ {
			tower.Write(currentMessage.matrix.BitmapSliceAt(i, tower.Columns))
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
				tower.Write(currentMessage.matrix.BitmapSliceAt(i, tower.Columns))
			}
			// skip preamble for the next run
			start = currentMessage.preamble
		}
	}
}

// PostMessage is the main entry point of the REST server. JSON messages
// posted here are sent to the displayBuilder goroutine through the textMsgQueue.
func PostMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func sendMessage() {
	bg := ledmatrix.RGB(0, 0, 0)
	preamble := blank(tower.Columns, bg)

	spr := sprite.NewSpriteFromImage("tower-sep.png")

	matrix := ledmatrix.NewMatrix(tower.Rows, 0)
	writer := ledmatrix.NewWriter(matrix)
	writer.Spacer(matrix.Columns, 0) // Blank bootstrap
	writer.WriteText("Filière Télécommunications ", bitmapfont.F68, ledmatrix.RGB(248, 179, 30), bg)
	writer.WriteText("// ", bitmapfont.F68, ledmatrix.RGB(30, 30, 250), bg)
	writer.WriteText("Filière Télécommunications ", bitmapfont.F68, ledmatrix.RGB(248, 179, 30), bg)

	writer.WriteBitmap(spr.Bitmap)
	writer.WriteText(" Fribourg is awesome #TelecomTower ", bitmapfont.F68, ledmatrix.RGB(100, 50, 250), bg)
	writer.WriteBitmap(spr.Bitmap)
	writer.WriteText(" Fribourg is awesome #TelecomTower ", bitmapfont.F68, ledmatrix.RGB(100, 50, 250), bg)

	body := writer.Matrix

	introduction := renderedTextMessage([]Line{Line{"", 8, "#FF0000"}}, bg)
	separator := renderedTextMessage([]Line{Line{"        ", 8, "#FF0000"}}, bg)

	// Append the introduction to the preamble
	preamble.Append(introduction)

	// Compute the wrapper. The wrapper makes the link between the end and
	// the start of the message, allowing it to "roll" properly.
	wrapper := ledmatrix.Concat(separator, body.Slice(0, tower.Columns))

	// Send the bitmap message
	bitmapMsgQueue <- BitmapMessage{
		matrix:     ledmatrix.Concat(preamble, body, wrapper),
		preamble:   preamble.Columns,
		checkpoint: preamble.Columns + body.Columns - tower.Columns,
	}

}

// tower-server starts a REST server and starts the towerServer goroutine and
// displayBuilder goroutine. The rest of the job is done in the PostMessage
// method.
func main() {
	var port = flag.Int("port", DEFAULT_PORT, "port")
	var brightness = flag.Int(
		"brightness", DEFAULT_BRIGHTNESS,
		"Brightness between 0 and 255.")

	flag.Parse()

	tower.Init(*brightness)
	textMsgQueue = make(chan TextMessage, TEXT_MSG_QUEUE_SIZE)
	bitmapMsgQueue = make(chan BitmapMessage, BITMAP_MSG_QUEUE_SIZE)

	go towerServer()
	log.Println("Tower ready.")

	sendMessage()

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/message").HandlerFunc(PostMessage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
