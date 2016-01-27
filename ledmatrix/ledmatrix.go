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

type Stripe []Color

type Matrix struct {
	Rows    int
	Columns int
	Bitmap  Stripe
}

func NewMatrix(rows, columns int) *Matrix {
	m := new(Matrix)
	m.Rows = rows
	m.Columns = columns
	m.Bitmap = make(Stripe, m.Rows*m.Columns)
	return m
}

func (m *Matrix) SetPixel(x, y int, color Color) {
	// Check if the coordinate (x,y) is valid in the Matrix. Extend the
	// Matrix if there is not enough space for x
	if y < 0 || y >= m.Rows {
		panic("y out of bound")
	}
	if x < 0 {
		panic("x out of bound")
	}
	if x >= m.Columns { // resize
		requiredSize := (x + 1) * m.Rows
		if cap(m.Bitmap) < requiredSize {
			t := make(Stripe, requiredSize)
			copy(t, m.Bitmap)
			m.Bitmap = t
		} else {
			m.Bitmap = m.Bitmap[:requiredSize]
		}
		m.Columns = x + 1
	}

	m.Bitmap[x*m.Rows+y] = color
}

func (m *Matrix) GetPixel(x, y int) Color {
	if y < 0 || y >= m.Rows {
		panic("y out of bound")
	}
	if x < 0 || x >= m.Columns {
		panic("x out of bound")
	}
	return m.Bitmap[x*m.Rows+y]
}

func (m *Matrix) Slice(low, high int) *Matrix {
	res := new(Matrix)
	res.Bitmap = m.Bitmap[low*m.Rows : high*m.Rows]
	res.Rows = m.Rows
	res.Columns = high - low
	return res
}

func (m *Matrix) Append(x ...*Matrix) {
	for _, i := range x {
		if m.Rows != i.Rows {
			panic("Error Matrix Operation")
		}
		m.Bitmap = append(m.Bitmap, i.Bitmap...)
		m.Columns += i.Columns
	}
}

func Concat(a *Matrix, b ...*Matrix) *Matrix {
	res := new(Matrix)
	res.Columns = a.Columns
	res.Rows = a.Rows
	res.Bitmap = append(res.Bitmap, a.Bitmap...)
	res.Append(b...)
	return res
}

func (m *Matrix) InterleavedStripes() []Stripe {
	res := make([]Stripe, 2)
	res[0] = make(Stripe, m.Rows*m.Columns)
	res[1] = make(Stripe, m.Rows*m.Columns)
	for x := 0; x < m.Columns; x++ {
		for y := 0; y < m.Rows; y++ {
			if x%2 == 0 { // Even Column
				res[0][x*m.Rows+y] = m.Bitmap[x*m.Rows+y]
				res[1][x*m.Rows+y] = m.Bitmap[x*m.Rows+m.Rows-1-y]
			} else { // Odd Column
				res[0][x*m.Rows+y] = m.Bitmap[x*m.Rows+m.Rows-1-y]
				res[1][x*m.Rows+y] = m.Bitmap[x*m.Rows+y]
			}
		}
	}
	return res
}
