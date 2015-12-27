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
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/heia-fr/telecom-tower/cloud/tower"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/ledmatrix/font"
	"log"
	"net/http"
)

const (
	textMsqQueueSsize  = 32
	bitmapMsgQueueSize = 32
	defaultPort        = 8080
	defaultBrightness  = 32
	gpioPin            = 18
)

type Line struct {
	Text  string `json:"text"`
	Font  int    `json:"font"`
	Color string `json:"color"`
}

type TextMessage struct {
	Body         []Line `json:"body"`
	Introduction []Line `json:"introduction"`
	Conclusion   []Line `json:"conclusion"`
	Separator    []Line `json:"separator"`
	roll         bool
}

type BitmapMessage struct {
	stripes    []ledmatrix.Stripe
	preamble   int
	checkpoint int
	roll       bool
}

// textMsgQueue is the input queue for the displayBuilder goroutine.
var textMsgQueue chan TextMessage

// rollingBitmapMsgQueue is the input queue for the towerServer goroutine.
var rollingBitmapMsgQueue chan BitmapMessage

// renderdeTextMessage convert a text message to a LED bitmap matrix.
// This method is used by the displayBuilder gorouting.
func renderedTextMessage(message []Line, bg ledmatrix.Color) *ledmatrix.Matrix {
	matrix := ledmatrix.NewMatrix(tower.Rows, 0)
	writer := ledmatrix.NewWriter(matrix)
	writer.Spacer(matrix.Columns(), 0) // Blank bootstrap

	for _, line := range message {
		var r, g, b int
		var f font.Font

		if line.Font < 8 {
			f = font.Font6x8
		} else {
			f = font.Font8x8
		}

		fmt.Sscanf(line.Color, "#%02x%02x%02x", &r, &g, &b)
		writer.WriteText(line.Text, f, ledmatrix.RGB(r, g, b), bg)
	}
	return matrix
}

// blank generates a "space" from the given color
func blank(len int, bg ledmatrix.Color) *ledmatrix.Matrix {
	matrix := ledmatrix.NewMatrix(tower.Rows, tower.Columns)
	writer := ledmatrix.NewWriter(matrix)
	writer.Spacer(tower.Columns, bg) // Blank spacer
	return matrix
}

// displayBuilder is a goroutine that converts text messages to bitmaps. The data
// usually comes from a REST request and the bitmap is then sent to the
// towerServer
func displayBuilder() {
	bg := ledmatrix.RGB(0, 0, 0)
	preamble := blank(tower.Columns, bg)

	for { // Loop forever
		// Receive a text message
		message := <-textMsgQueue

		// If body is nil, then reset preamble. This is used when static frames
		// are sent to the tower and the next rolling message should start
		// with a blank.
		if message.Body == nil {
			preamble = blank(tower.Columns, bg)
			continue
		}

		// render all parts of the message
		body := renderedTextMessage(message.Body, bg)
		introduction := renderedTextMessage(message.Introduction, bg)
		conclusion := renderedTextMessage(message.Conclusion, bg)
		separator := renderedTextMessage(message.Separator, bg)

		if body.Columns() == 0 {
			log.Println("Skip empty message")
			continue
		}

		// we need at least a full display. So if the message is too short,
		// then append tbe separator and aother copy of the message.
		// We re-render the message and this is not optimal, but this
		// is only usd for shirt messages, so this is OK.
		for body.Columns() < tower.Columns {
			body.Append(separator, renderedTextMessage(message.Body, bg))
		}

		// Append the introduction to the preamble
		preamble.Append(introduction)

		// Compute the wrapper. The wrapper makes the link between the end and
		// the start of the message, allowing it to "roll" properly.
		wrapper := ledmatrix.Concat(separator, body.Slice(0, tower.Columns))

		// Send the bitmap message
		rollingBitmapMsgQueue <- BitmapMessage{
			stripes:    ledmatrix.Concat(preamble, body, wrapper).InterleavedStripes(),
			preamble:   preamble.Columns(),
			checkpoint: preamble.Columns() + body.Columns() - tower.Columns,
			roll:       message.roll,
		}

		// Compute the next preamble using the last frame of the body and the conclusion
		preamble = ledmatrix.Concat(body.Slice(body.Columns()-tower.Columns, body.Columns()), conclusion)
	}
}

func towerRoll(message BitmapMessage, low, high int) {
	for i := low; i < high; i++ {
		tower.SendFrame(
			message.stripes[i%2][i*tower.Rows : (i+tower.Columns)*tower.Rows])
	}
}

// towerServer is a the goroutine that receives bitmap messages from the displayBuilder
// and dispatch "frames" to the tower LEDs. The preamble is only sent once; at the
// checkpoint, the goroutine checks if a new message is available; if yes, it switches
// to this new message; if no, it finish the message and roll th same message again.
func towerServer() {
	var currentMessage BitmapMessage
	var roll chan int

	for { // Loop forever
		select {
		case m := <-rollingBitmapMsgQueue:
			log.Println("New rolling message...")
			currentMessage = m
			if m.roll {
				roll = make(chan int, 1)
				roll <- 0
			} else {
				roll = nil
			}
			// Display the message at least once
			towerRoll(currentMessage, 0, currentMessage.checkpoint)
		case r := <-roll:
			towerRoll(
				currentMessage,
				currentMessage.checkpoint,                               // from checkpoint
				len(currentMessage.stripes[0])/tower.Rows-tower.Columns) // to the last position
			towerRoll(currentMessage,
				currentMessage.preamble,   // from the preamble
				currentMessage.checkpoint) // to the checkpoint
			roll <- r
		}
	}
}

// PostMessage is the main entry point of the REST server. JSON messages
// posted here are sent to the displayBuilder goroutine through the textMsgQueue.
func PostMessage(w http.ResponseWriter, r *http.Request) {
	var message TextMessage
	dec := json.NewDecoder(r.Body)

	// decode the JSON message
	if err := dec.Decode(&message); err != nil {
		fmt.Fprintf(w, "err\n")
	} else {
		message.roll = true
		textMsgQueue <- message
		fmt.Fprintf(w, "OK\n")
	}
}

func ToastMessage(w http.ResponseWriter, r *http.Request) {
	// NYI
}

// tower-server starts a REST server and starts the towerServer goroutine and
// displayBuilder goroutine. The rest of the job is done in the PostMessage
// method.
func main() {
	var port = flag.Int("port", defaultPort, "port")
	var brightness = flag.Int(
		"brightness", defaultBrightness,
		"Brightness between 0 and 255.")

	flag.Parse()

	textMsgQueue = make(chan TextMessage, textMsqQueueSsize)
	rollingBitmapMsgQueue = make(chan BitmapMessage, bitmapMsgQueueSize)

	if err := tower.Init(gpioPin, *brightness); err != nil {
		log.Fatalf("Error: %s", err)
	}
	go towerServer()
	go displayBuilder()
	log.Println("Tower ready.")

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/message").HandlerFunc(PostMessage)
	router.Methods("POST").Path("/toast").HandlerFunc(ToastMessage)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
