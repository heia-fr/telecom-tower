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

package sprite

import (
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type Sprite struct {
	Width, Height int
	Bitmap        [][]ledmatrix.Color
}

func NewSprite(width, height int) *Sprite {
	s := new(Sprite)
	s.Width = width
	s.Height = height
	s.Bitmap = make([][]ledmatrix.Color, height)
	for i := 0; i < height; i++ {
		s.Bitmap[i] = make([]ledmatrix.Color, width)
	}
	return s
}

func NewSpriteFromImage(fileName string) *Sprite {
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	if bounds.Max.Y-bounds.Min.Y != 8 {
		log.Fatal("Invalid image size")
	}

	s := NewSprite(bounds.Max.X-bounds.Min.X, bounds.Max.Y-bounds.Min.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			s.SetPixel(x-bounds.Min.X, y-bounds.Min.Y, ledmatrix.RGB(int(r>>8), int(g>>8), int(b>>8)))
		}
	}

	return s
}

func (s *Sprite) SetPixel(x, y int, c ledmatrix.Color) {
	s.Bitmap[y][x] = c
}
