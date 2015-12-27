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
	rows    int
	columns int
	bitmap  Stripe
}

func NewMatrix(rows, columns int) *Matrix {
	m := new(Matrix)
	m.rows = rows
	m.columns = columns
	m.bitmap = make(Stripe, m.rows*m.columns)
	return m
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Columns() int {
	return m.columns
}

func (m *Matrix) SetPixel(x, y int, color Color) {
	// Check if the coordinate (x,y) is valid in the Matrix. Extend the
	// Matrix if there is not enough space for x
	if y < 0 || y >= m.rows {
		panic("y out of bound")
	}
	if x < 0 {
		panic("x out of bound")
	}
	if x >= m.columns { // resize
		requiredSize := (x + 1) * m.rows
		if cap(m.bitmap) < requiredSize {
			t := make(Stripe, requiredSize)
			copy(t, m.bitmap)
			m.bitmap = t
		} else {
			m.bitmap = m.bitmap[:requiredSize]
		}
		m.columns = x + 1
	}

	m.bitmap[x*m.rows+y] = color
}

func (m *Matrix) GetPixel(x, y int) Color {
	if y < 0 || y >= m.rows {
		panic("y out of bound")
	}
	if x < 0 || x >= m.columns {
		panic("x out of bound")
	}
	return m.bitmap[x*m.rows+y]
}

func (m *Matrix) Slice(low, high int) *Matrix {
	res := new(Matrix)
	res.bitmap = m.bitmap[low*m.rows : high*m.rows]
	res.rows = m.rows
	res.columns = high - low
	return res
}

func (m *Matrix) Append(x ...*Matrix) {
	for _, i := range x {
		if m.rows != i.rows {
			panic("Error Matrix Operation")
		}
		m.bitmap = append(m.bitmap, i.bitmap...)
		m.columns += i.columns
	}
}

func Concat(a *Matrix, b ...*Matrix) *Matrix {
	res := new(Matrix)
	res.columns = a.columns
	res.rows = a.rows
	res.bitmap = append(res.bitmap, a.bitmap...)
	res.Append(b...)
	return res
}

func (m *Matrix) InterleavedStripes() []Stripe {
	res := make([]Stripe, 2)
	res[0] = make(Stripe, m.rows*m.columns)
	res[1] = make(Stripe, m.rows*m.columns)
	for x := 0; x < m.columns; x++ {
		for y := 0; y < m.rows; y++ {
			if x%2 == 0 { // Even Column
				res[0][x*m.rows+y] = m.bitmap[x*m.rows+y]
				res[1][x*m.rows+y] = m.bitmap[x*m.rows+m.rows-1-y]
			} else { // Odd Column
				res[0][x*m.rows+y] = m.bitmap[x*m.rows+m.rows-1-y]
				res[1][x*m.rows+y] = m.bitmap[x*m.rows+y]
			}
		}
	}
	return res
}
