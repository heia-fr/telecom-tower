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

const Rows = 8
const Columns = 128

type Color uint32

type Matrix struct {
	Bitmap [2][]Color
}

func NewMatrix() *Matrix {
	m := new(Matrix)
	for i := 0; i < len(m.Bitmap); i++ {
		m.Bitmap[i] = make([]Color, Rows*Columns)
	}
	return m
}

func RGB(r, g, b int) Color {
	if r < 0 || r > 255 {
		panic("Red component must be between 0 and 255")
	}
	if g < 0 || g > 255 {
		panic("Green component must be between 0 and 255")
	}
	if b < 0 || b > 255 {
		panic("Blue component must be between 0 and 255")
	}
	return Color(r)<<16 + Color(g)<<8 + Color(b)
}

func (m *Matrix) checkXY(x, y int) {
	if y < 0 || y >= Rows {
		panic("y out of bound")
	}
	requiredSize := (x + 1) * Rows
	if len(m.Bitmap[0]) < requiredSize {
		if cap(m.Bitmap[0]) < requiredSize {
			for i := 0; i < len(m.Bitmap); i++ {
				t := make([]Color, requiredSize)
				copy(t, m.Bitmap[i])
				m.Bitmap[i] = t
			}
		} else {
			for i := 0; i < len(m.Bitmap); i++ {
				m.Bitmap[i] = m.Bitmap[i][:requiredSize]
			}
		}
	}
}

func (m *Matrix) SetPixel(x, y int, color Color) {
	m.checkXY(x, y)
	if x%2 == 0 {
		m.Bitmap[0][x*Rows+y] = color
		m.Bitmap[1][x*Rows+(Rows-1-y)] = color
	} else {
		m.Bitmap[0][x*Rows+(Rows-1-y)] = color
		m.Bitmap[1][x*Rows+y] = color
	}
}

func (m *Matrix) GetPixel(x, y int) Color {
	m.checkXY(x, y)
	if x%2 == 0 {
		return m.Bitmap[0][x*Rows+y]
	} else {
		return m.Bitmap[0][x*Rows+(Rows-1-y)]
	}
}

func (m *Matrix) Columns() int {
	return len(m.Bitmap[0]) / Rows
}

func (m *Matrix) SliceAt(column int) []Color {
	index := column % 2
	return m.Bitmap[index][column*Rows:column*Rows+Rows*Columns]

}
