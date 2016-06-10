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
	"github.com/heia-fr/telecom-tower/ledmatrix/font"
)

type Writer struct {
	matrix *Matrix
	pos    int
}

func NewWriter(m *Matrix) *Writer {
	w := new(Writer)
	w.matrix = m
	w.pos = 0
	return w
}

func (w *Writer) SetPos(pos int) {
	w.pos = pos
}

func (w *Writer) Pos() int {
	return w.pos
}

func (w *Writer) WriteText(text string, fnt font.Font, color, bgColor uint32) {
	for _, c := range font.ExpandAlias(text) {
		for _, i := range fnt.Bitmap(c) {
			for k := 0; k < fnt.Height(); k++ {
				if i&(1<<uint(k)) != 0 {
					w.matrix.SetPixel(w.pos, k, color)
				} else {
					w.matrix.SetPixel(w.pos, k, bgColor)
				}
			}
			w.pos++
		}
	}
}

func (w *Writer) WriteBitmap(bitmap [][]uint32) {
	width := 0
	for y := 0; y < len(bitmap); y++ {
		if len(bitmap[y]) > width {
			width = len(bitmap[y])
		}
		for x := 0; x < len(bitmap[y]); x++ {
			w.matrix.SetPixel(w.pos+x, y, bitmap[y][x])
		}
	}
	w.pos += width
}

func (w *Writer) Spacer(width int, color uint32) {
	for y := 0; y < w.matrix.Rows; y++ {
		for x := 0; x < width; x++ {
			w.matrix.SetPixel(w.pos+x, y, color)
		}
	}
	w.pos += width
}
