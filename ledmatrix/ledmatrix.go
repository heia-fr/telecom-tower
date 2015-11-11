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
	"fmt"
)

type Matrix struct {
	Rows    int
	Columns int
	Bitmap  [2][]Color
}

func NewMatrix(rows, columns int) *Matrix {
	m := new(Matrix)
	m.Rows = rows
	m.Columns = columns
	for i := 0; i < len(m.Bitmap); i++ {
		m.Bitmap[i] = make([]Color, m.Rows*m.Columns)
	}
	return m
}

// Check if the bitmaps sizes are consistent with the dimension of the matrix
func (m *Matrix) check() {
	size := m.Rows * m.Columns
	for i := 0; i < len(m.Bitmap); i++ {
		if len(m.Bitmap[i]) != size {
			panic(fmt.Sprintf("Bad bitmap %d. Size shoud be %d and is %d", i, size, len(m.Bitmap[i])))
		}
	}
}

// Check if the coordinate (x,y) is valid in the Matrix. Extend the
// Matrix if there is not enough space for x
func (m *Matrix) checkXY(x, y int) {
	if y < 0 || y >= m.Rows {
		panic("y out of bound")
	}

	if x >= m.Columns {
		requiredSize := (x + 1) * m.Rows
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
		m.Columns = x + 1
	}
}

func (m *Matrix) SetPixel(x, y int, color Color) {
	m.checkXY(x, y)
	if x%2 == 0 {
		m.Bitmap[0][x*m.Rows+y] = color
		m.Bitmap[1][x*m.Rows+(m.Rows-1-y)] = color
	} else {
		m.Bitmap[0][x*m.Rows+(m.Rows-1-y)] = color
		m.Bitmap[1][x*m.Rows+y] = color
	}
	m.check()
}

func (m *Matrix) GetPixel(x, y int) Color {
	m.checkXY(x, y)
	if x%2 == 0 {
		return m.Bitmap[0][x*m.Rows+y]
	} else {
		return m.Bitmap[0][x*m.Rows+(m.Rows-1-y)]
	}
}

func (m *Matrix) BitmapSliceAt(column, size int) []Color {
	index := column % 2
	return m.Bitmap[index][column*m.Rows : (column+size)*m.Rows]
}

func (m *Matrix) Slice(low, high int) *Matrix {
	res := new(Matrix)
	if low%2 == 0 {
		res.Bitmap = [2][]Color{
			m.Bitmap[0][low*m.Rows : high*m.Rows],
			m.Bitmap[1][low*m.Rows : high*m.Rows]}
	} else {
		res.Bitmap = [2][]Color{
			m.Bitmap[1][low*m.Rows : high*m.Rows],
			m.Bitmap[0][low*m.Rows : high*m.Rows]}
	}

	res.Rows = m.Rows
	res.Columns = high - low
	res.check()
	return res
}

func (m *Matrix) Append(x ...*Matrix) {
	for _, i := range x {
		if m.Rows != i.Rows {
			panic("Error Matrix Operation")
		}
		if m.Columns%2 == 0 {
			m.Bitmap = [2][]Color{
				append(m.Bitmap[0], i.Bitmap[0]...),
				append(m.Bitmap[1], i.Bitmap[1]...)}
		} else {
			m.Bitmap = [2][]Color{
				append(m.Bitmap[0], i.Bitmap[1]...),
				append(m.Bitmap[1], i.Bitmap[0]...)}
		}
		m.Columns += i.Columns
	}
	m.check()
}

func Concat(a *Matrix, b ...*Matrix) *Matrix {
	res := new(Matrix)
	res.Columns = a.Columns
	res.Rows = a.Rows
	res.Bitmap = [2][]Color{
		append(res.Bitmap[0], a.Bitmap[0]...),
		append(res.Bitmap[1], a.Bitmap[1]...)}

	res.Append(b...)
	return res
}
