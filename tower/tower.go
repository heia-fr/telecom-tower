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
)

const (
	queueSize = 8
	gpioPin   = 18
)

var Queue chan []ledmatrix.Color
var initialized = false

func Init(brightness int) {
	if initialized {
		panic("Tower already initialized!")
	}
	initialized = true
	ws2811.Init(gpioPin, ledmatrix.Rows*ledmatrix.Columns, brightness)
	Queue = make(chan []ledmatrix.Color, queueSize)
	go func() { // Phantom of the Tower
		for {
			req := <-Queue
			ws2811.SetBitmap(req)
			ws2811.Render()
			ws2811.Wait()
		}
	}()
}

func Shutdown() {
	ws2811.Wait()
	ws2811.Clear()
	ws2811.Render()
	ws2811.Wait()
	// ws2811.Fini()
}

func Roll(w *ledmatrix.Writer) {
	w.ExtendCircular()
	for i := 0; i < w.Pos(); i++ {
		Queue <- w.Matrix.SliceAt(i)
	}
}
