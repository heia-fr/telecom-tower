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
// Package to send info to a NATS server
//
package tower

import (
	"errors"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/nats-io/nats"
	"log"
	"math"
	"time"
)

const (
	Rows    = 8
	Columns = 128
	subject = "tower.frames"
	fps     = 33.0 // frame per secons
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
	conn         *nats.EncodedConn
	displayQueue chan ledmatrix.Stripe
	closing      chan chan error
	initialized  bool
}

// Init initialize the tower and starts the Tower's Daemon. The daemon
// receives bitmap matrix frames and display them on the LEDs.
func Init(gpioPin int, brightness int) error {
	var nc *nats.Conn
	var err error
	if nc, err = nats.Connect(nats.DefaultURL); err != nil {
		return err
	}
	if tower.conn, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER); err != nil {
		return err
	}

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
		case frame := <-tower.displayQueue:
			if len(frame) == Columns*Rows {
				tower.conn.Publish(subject, frame)
				time.Sleep(time.Duration(math.Floor((1.0 / fps) * float64(time.Second))))
			} else {
				log.Fatalf("Invalid frame (%d instead of %d)", len(frame), Columns*Rows)
			}
		}
	}
}

// SendFrame sends the frame to the daemon through the displayQueue channel.
func SendFrame(frame []ledmatrix.Color) {
	tower.displayQueue <- frame
}

// Shutdown closes the daemon using the closing channel and shuts down
// the tower.
func Shutdown() {
	if tower.initialized {
		errc := make(chan error)
		tower.closing <- errc
		<-errc
		tower.conn.Close()
		tower.initialized = false
	}
}
