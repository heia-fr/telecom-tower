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

package ledmatrix

import (
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/cp850"
)

type Writer struct {
	Matrix *Matrix
	pos    int
}

func NewWriter(m *Matrix) *Writer {
	w := new(Writer)
	w.Matrix = m
	w.pos = 0
	return w
}

func (w *Writer) SetPos(pos int) {
	w.pos = pos
}

func (w *Writer) Pos() int {
	return w.pos
}

func (w *Writer) WriteText(text string, font bitmapfont.Font, color, bgColor Color) {
	t := cp850.StringToCp850(text)

	for i := 0; i < len(t); i++ {
		a := font.GetCharMSB(t[i])
		for j := 0; j < font.Width; j++ {
			for k := 0; k < Rows; k++ {
				if a[j]&(1<<uint(k)) != 0 {
					w.Matrix.SetPixel(w.pos+j, Rows-1-k, color)
				} else {
					w.Matrix.SetPixel(w.pos+j, Rows-1-k, bgColor)
				}
			}
		}
		w.pos += font.Width
	}
}

func (w *Writer) WriteBitmap(bitmap [][]Color) {
	width := 0
	for y := 0; y < len(bitmap); y++ {
		if len(bitmap[y]) > width {
			width = len(bitmap[y])
		}
		for x := 0; x < len(bitmap[y]); x++ {
			w.Matrix.SetPixel(w.pos+x, y, bitmap[y][x])
		}
	}
	w.pos += width
}

func (w *Writer) Spacer(width int, color Color) {
	for y := 0; y < Rows; y++ {
		for x := 0; x < width; x++ {
			w.Matrix.SetPixel(w.pos+x, y, color)
		}
	}
	w.pos += width
}

func (w *Writer) ExtendCircular(from int) {
	for i := 0; i < Columns; i++ {
		for j := 0; j < Rows; j++ {
			w.Matrix.SetPixel(w.pos+i, j, w.Matrix.GetPixel(from + i, j))
		}
	}
}

func (w *Writer) ExtendClear(bgColor Color) {
	for i := 0; i < Columns; i++ {
		for j := 0; j < Rows; j++ {
			w.Matrix.SetPixel(w.pos+i, j, bgColor)
		}
	}
}
