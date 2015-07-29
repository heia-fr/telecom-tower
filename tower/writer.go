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

package tower

import (
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/cp850"
	"github.com/heia-fr/telecom-tower/ws2811"
)

type Writer struct {
	Bitmap [2][]ws2811.Color
	pos    int
}

func NewWriter() *Writer {
	w := new(Writer)
	w.pos = 0
	w.Bitmap[0] = make([]ws2811.Color, screenHeight*screenWidth)
	w.Bitmap[1] = make([]ws2811.Color, screenHeight*screenWidth)
	return w
}

func (w *Writer) checkSize(size int) {
	for i := 0; i < len(w.Bitmap); i++ {
		if cap(w.Bitmap[i]) < size {
			t := make([]ws2811.Color, size)
			copy(t, w.Bitmap[i])
			w.Bitmap[i] = t
		}
	}
}

func (w *Writer) Write(text string, font bitmapfont.Font, color ws2811.Color) {
	t := cp850.StringToCp850(text)

	delta := len(t) * font.Width * screenHeight
	w.checkSize(w.pos + delta)

	for i := 0; i < len(t); i++ {
		a := font.GetCharMLSB(t[i])
		pos := w.pos
		for j := 0; j < font.Width; j++ {
			for k := 0; k < screenHeight; k++ {
				if a[j]&(1<<uint(k)) != 0 {
					w.Bitmap[0][pos] = color
				} else {
					w.Bitmap[0][pos] = 0
				}
				pos++
			}
		}

		a = font.GetCharLMSB(t[i])
		pos = w.pos
		for j := 0; j < font.Width; j++ {
			for k := 0; k < screenHeight; k++ {
				if a[j]&(1<<uint(k)) != 0 {
					w.Bitmap[1][pos] = color
				} else {
					w.Bitmap[1][pos] = 0
				}
				pos++
			}
		}

		w.pos = pos
	}
}

func (w *Writer) Finish() {
	delta := screenHeight * screenWidth
	w.checkSize(w.pos + delta)

	for b := 0; b < len(w.Bitmap); b++ {
		for src, dst := 0, w.pos; dst < w.pos+delta; src, dst = src+1, dst+1 {
			w.Bitmap[b][dst] = w.Bitmap[b][src]
		}
	}
}

func (w *Writer) Columns() int {
	return w.pos / screenHeight
}
