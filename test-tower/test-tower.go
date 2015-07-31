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
	"github.com/heia-fr/telecom-tower/sprite"
	"github.com/heia-fr/telecom-tower/tower"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Println("Booting tower...")
	matrix := ledmatrix.NewMatrix()
	writer := ledmatrix.NewWriter(matrix)

	tower.Init(128)

	bg := ledmatrix.RGB(0, 0, 0)
	tux := sprite.NewSpriteFromImage("tux.png")

	writer.WriteText(
		"Haute école d'ingénierie et d'architecture Fribourg",
		bitmapfont.F88,
		ledmatrix.RGB(150, 200, 255),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Informatique",
		bitmapfont.F68,
		ledmatrix.RGB(223, 0, 30),
		bg,
	)

	/*
		writer.WriteText(
			" \u2665 ",
			bitmapfont.F88,
			ledmatrix.RGB(255, 0, 0),
			bg,
		)
	*/

	writer.Spacer(8, 0)
	writer.WriteBitmap(tux.Bitmap)
	writer.Spacer(8, 0)

	writer.WriteText(
		"Télécommunications",
		bitmapfont.F68,
		ledmatrix.RGB(248, 179, 30),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Génie électrique",
		bitmapfont.F68,
		ledmatrix.RGB(222, 180, 8),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Génie mécanique",
		bitmapfont.F68,
		ledmatrix.RGB(100, 60, 4),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Chimie",
		bitmapfont.F68,
		ledmatrix.RGB(0, 200, 30),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Architecture",
		bitmapfont.F68,
		ledmatrix.RGB(110, 40, 140),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Génie civil",
		bitmapfont.F68,
		ledmatrix.RGB(90, 90, 180),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Ecole technique de la construction",
		bitmapfont.F68,
		ledmatrix.RGB(90, 90, 90),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Hochschule für Technik und Architektur Freiburg",
		bitmapfont.F88,
		ledmatrix.RGB(150, 200, 255),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Informatik",
		bitmapfont.F68,
		ledmatrix.RGB(223, 0, 30),
		bg,
	)
	/*
		writer.WriteText(
			" \u2665 ",
			bitmapfont.F88,
			ledmatrix.RGB(255, 0, 0),
			bg,
		)
	*/

	writer.Spacer(8, 0)
	writer.WriteBitmap(tux.Bitmap)
	writer.Spacer(8, 0)

	writer.WriteText(
		"Telekommunikation",
		bitmapfont.F68,
		ledmatrix.RGB(248, 179, 30),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Elektrotechnik",
		bitmapfont.F68,
		ledmatrix.RGB(222, 180, 8),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Maschinentechnik",
		bitmapfont.F68,
		ledmatrix.RGB(100, 60, 4),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Chemie",
		bitmapfont.F68,
		ledmatrix.RGB(0, 200, 30),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Architektur",
		bitmapfont.F68,
		ledmatrix.RGB(110, 40, 140),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Bauingenieurwesen",
		bitmapfont.F68,
		ledmatrix.RGB(90, 90, 180),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	writer.WriteText(
		"Bautechnische Schule",
		bitmapfont.F68,
		ledmatrix.RGB(90, 90, 90),
		bg,
	)
	writer.WriteText(
		" \u2665 ",
		bitmapfont.F88,
		ledmatrix.RGB(255, 0, 0),
		bg,
	)

	log.Println("Roll!")

	c := make(chan os.Signal, 1)
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		s := <-c
		log.Println("Shutting down tower... please wait for message finish")
		signalChannel <- s
	}()

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
}
