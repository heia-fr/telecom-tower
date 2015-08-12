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
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	// "github.com/heia-fr/telecom-tower/sprite"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/heia-fr/telecom-tower/tower"
	"log"
	"os"
	"os/signal"
)

func main() {
	var jsonFile string
	var roll bool
	flag.StringVar(&jsonFile, "message", "", "JSON File with data")
	flag.BoolVar(&roll, "roll", true, "n/a")

	flag.Parse()

	log.Println("Booting tower...")
	matrix := ledmatrix.NewMatrix()
	writer := ledmatrix.NewWriter(matrix)

	tower.Init(128)

	bg := ledmatrix.RGB(0, 0, 0)
	// tux := sprite.NewSpriteFromImage("tux.png")

	var data struct {
		Message []struct {
			Text  string
			Font  int
			Color string
		}
	}

	f, _ := os.Open(jsonFile)
	dec := json.NewDecoder(f)

	if err := dec.Decode(&data); err != nil {
		log.Println(err)
		return
	}

	for _, line := range data.Message {
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

	log.Printf("Columns: %v\n", matrix.Columns())
	log.Println("Roll!")

	c := make(chan os.Signal, 1)
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		s := <-c
		log.Println("Shutting down tower... please wait for message finish")
		signalChannel <- s
	}()

	if roll {
		for true {
			tower.Roll(writer)
			select {
			case _ = <-signalChannel:
				log.Println("Shutting down tower now")
				tower.Shutdown()
				os.Exit(0)
			default:
				// continue
			}
		}
	} else {
		tower.Roll(writer)
		tower.Shutdown()
	}
}
