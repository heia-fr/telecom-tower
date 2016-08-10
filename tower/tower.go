// Copyright 2016 Jacques Supcik, HEIA-FR
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
// 2015-11-18 | JS | Version 1.0
// 2016-06-15 | JS | Version 2.0 with new ws2811
// 2016-08-10 | JS | Remove "virtual" version

//
// Package to display info on the Telecom Tower
//

package tower

import (
	"errors"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/supcik-go/ws2811"
	"log"
)

const (
	Columns = 128
	Rows    = 8

	// DisplayQueue capacity. If this value is small, the tower will be more
	// reactive and can decide to change the display quicker. If the value is
	// large, the display will be more smooth and the sender can have some
	// delays in the processing of the information. A value of between 4 and
	// 16 seems reasonable here.
	queueSize = 16
)

var tower struct {
	// DisplayQueue is the communication channel to the Tower's Daemon. The queue
	// will be created by the Init method and will have a capacity of "queueSize"
	ws           *ws2811.WS2811
	displayQueue chan ledmatrix.Stripe
	closing      chan chan error
	initialized  bool
}

// Init initialize the tower and starts the Tower's Daemon. The daemon
// receives bitmap matrix frames and display them on the LEDs.
func Init(gpioPin int, brightness int) error {
	ws, err := ws2811.MakeWS2811(ws2811.TargetFreq, ws2811.DefaultDmaNum)
	if err != nil {
		return err
	}
	ws.SetChannel(0, gpioPin, Rows*Columns, brightness, 0, ws2811.StripGRB)
	err = ws.Init()
	if err != nil {
		return err
	}
	tower.ws = ws
	tower.displayQueue = make(chan ledmatrix.Stripe, queueSize)
	tower.initialized = true
	go daemon()
	return nil
}

// Phantom of the Tower. Goroutine that implements the main loop of the tower.
// The goroutine communicates with the channels "closing" and "displayQueue".
func daemon() {
	for {
		select {
		case errc := <-tower.closing:
			errc <- errors.New("OK")
			return
		case req := <-tower.displayQueue:
			if len(req) == Columns*Rows {
				tower.ws.SetBitmap(req)
				tower.ws.Render()
				tower.ws.Wait()
			} else {
				log.Fatalf("Invalid frame (%d instead of %d)", len(req), Columns*Rows)
			}
		}
	}
}

// Write sends the frame to the daemon through the displayQueue channel.
func SendFrame(frame []uint32) {
	tower.displayQueue <- frame
}

// Shutdown closes the daemon using the closing channel and shuts down
// the tower.
func Shutdown() {
	if tower.initialized {
		errc := make(chan error)
		tower.closing <- errc
		<-errc
		tower.ws.Wait()
		tower.ws.Clear()
		tower.ws.Render()
		tower.ws.Wait()
		tower.ws.Fini()
		tower.initialized = false
	}
}
