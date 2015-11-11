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

// Package to display info on the Telecom Tower
package tower

import (
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/ws2811"
	"log"
)

const (
	// DisplayQueue capacity. If this value is small, the tower will be more
	// reactive and can decide to change the display quicker. If the value is
	// large, the display will be more smooth and the sender can have some
	// delays in the processing of the information. A value of between 4 and
	// 16 seems reasonable here.
	queueSize = 16
	// gpioPin is the pin where the LEDs matrix control wire is connected to
	// For the current Telecom Tower, the electronic board uses the pin 18.
	gpioPin = 18
	// Columns is the display width. For the current Telecom Tower this is 128.
	Columns = 128
	// Rows is the display height. For the current Telecom Tower this is 8.
	Rows = 8
)

// DisplayQueue is the communication channel to the Tower's Daemon. The queue
// will be created by the Init method and will have a capacity of "queueSize"
var DisplayQueue chan []ledmatrix.Color

// initialized is a global variable used to prevent a double initialisation
// of the tower. This variable is used and set in the Init and Shutdown methods.
var initialized = false

// Init initialize the tower and starts the Tower's Daemon. The daemon
// receives bitmap matrix frames and display them on the LEDs.
func Init(brightness int) {
	if !initialized {
		initialized = true
		ws2811.Init(gpioPin, Rows*Columns, brightness)
		DisplayQueue = make(chan []ledmatrix.Color, queueSize)
		go func() { // Phantom of the Tower
			for {
				req := <-DisplayQueue
				if len(req) == Columns*Rows {
					ws2811.SetBitmap(req)
					ws2811.Render()
					ws2811.Wait()
				} else {
					log.Fatalf("Invalid frame (%d instead of %d)", len(req), Columns*Rows)
				}
			}
		}()
	}
}

func Shutdown() {
	if initialized {
		ws2811.Wait()
		ws2811.Clear()
		ws2811.Render()
		ws2811.Wait()
		ws2811.Fini()
		initialized = false
	}
}
