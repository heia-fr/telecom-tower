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

// Package to display text on the Telecom Tower
package tower

import (
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/color"
	"github.com/heia-fr/telecom-tower/ws2811"
)

const screenWidth = 128
const screenHeight = 8

var W *Writer

func Init(brightness int) {
	W = NewWriter()
	ws2811.Init(18, 1024, brightness)
}

func Shutdown() {
	ws2811.Clear()
	ws2811.Render()
	ws2811.Wait()
	ws2811.Fini()
}

func Write(text string, font bitmapfont.Font, color color.Color) {
	W.Write(text, font, color)
}

func Roll() {
	W.Finish()
	for i := 0; i < W.Columns()/2; i++ {
		ws2811.SetBitmap(W.Bitmap[0][i*16 : i*16+1024])
		ws2811.Render()
		ws2811.Wait()

		ws2811.SetBitmap(W.Bitmap[1][i*16+8 : i*16+8+1024])
		ws2811.Render()
		ws2811.Wait()
	}
}
