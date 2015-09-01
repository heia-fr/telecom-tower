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
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/tower"
	"log"
	"os"
	"os/signal"
)

type Line struct {
	Text  string
	Font  int
	Color string
}

type Message []Line

var signalChannel chan os.Signal
var msgQueue chan Message
var writerQueue chan *ledmatrix.Writer

func compose(data Message, bg ledmatrix.Color) *ledmatrix.Writer {
	matrix := ledmatrix.NewMatrix()
	writer := ledmatrix.NewWriter(matrix)
	for _, line := range data {
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
	return writer
}

func displayBuilder() {
	bg := ledmatrix.RGB(0, 0, 0)
	for {
		message := <-msgQueue
		writerQueue <- compose(message, bg)
	}
}

func towerServer() {
	data := make(Message, 1)
	data[0] = Line{Text: "Booting tower... please wait. ", Font: 6, Color: "#3333FF"}
	currentDisplay := compose(data, ledmatrix.RGB(0, 0, 0))

	for {
		select {
		case m := <-writerQueue:
			currentDisplay = m
		case _ = <-signalChannel:
			log.Println("Shutting down tower now")
			tower.Shutdown()
			os.Exit(0)
		default:
			// continue
		}
		tower.Roll(currentDisplay)
	}
}

func main() {
	var jsonFile string
	flag.StringVar(&jsonFile, "message", "", "JSON File with data")
	flag.Parse()

	log.Println("Booting tower...")
	tower.Init(128)

	msgQueue = make(chan Message, 32)
	writerQueue = make(chan *ledmatrix.Writer, 32)

	go towerServer()
	go displayBuilder()

	// Read json file
	var data struct {
		Message Message
	}

	f, _ := os.Open(jsonFile)
	dec := json.NewDecoder(f)

	if err := dec.Decode(&data); err != nil {
		log.Println(err)
		return
	}

	msgQueue <- data.Message
	// log.Printf("Columns: %v\n", matrix.Columns())
	log.Println("Roll!")

	c := make(chan os.Signal, 1)
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	//go func() {
	s := <-c
	log.Println("Shutting down tower... please wait for message finish")
	signalChannel <- s
	//}()

}
