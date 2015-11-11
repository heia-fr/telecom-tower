// Copyright 2015 Damien Goetschi, HEIA-FR
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
package lyrics

import (
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/ws2811"
)

var W *ledmatrix.Writer

func Init(brightness int) {

	ws2811.Init(18, ledmatrix.Rows*ledmatrix.Columns, brightness)
}

func Shutdown() {
	ws2811.Wait()
	ws2811.Clear()
	ws2811.Render()
	ws2811.Wait()
	// ws2811.Fini()
}

func WriteText(text string, font bitmapfont.Font, color, bgColor ledmatrix.Color) {
	W.WriteText(text, font, color, bgColor)
}

func Roll(w *ledmatrix.Writer) {
	for i := 0; i < w.Pos(); i+=2 {
		ws2811.SetBitmap(w.Matrix.SliceAt(i))
		ws2811.Render()
		ws2811.Wait()
	}
}
