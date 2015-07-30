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
package nyancat

import (
	"github.com/heia-fr/telecom-tower/color"
	"github.com/heia-fr/telecom-tower/ws2811"
)

const screenWidth = 128
const screenHeight = 8
const motifWidth = 32

var background color.Color
var pink1 color.Color
var pink2 color.Color
var salmon color.Color
var catHead color.Color
var catPaw color.Color
var cheek color.Color
var black color.Color
var white color.Color

var red color.Color
var orange color.Color
var yellow color.Color
var green color.Color
var blue1 color.Color
var blue2 color.Color

var Bitmap [2][screenHeight * (2*screenWidth + motifWidth)]color.Color

func Init(brightness int) {

	background = color.RGB(0, 51, 102)
	pink1 = color.RGB(253, 151, 253)
	pink2 = color.RGB(254, 48, 150)
	salmon = color.RGB(253, 202, 153)
	catHead = color.RGB(152, 152, 152)
	catPaw = color.RGB(162, 162, 162)
	cheek = color.RGB(253, 153, 153)
	black = color.RGB(0, 0, 0)
	white = color.RGB(255, 255, 255)

	red = color.RGB(253, 0, 0)
	orange = color.RGB(253, 153, 0)
	yellow = color.RGB(254, 254, 0)
	green = color.RGB(51, 253, 0)
	blue1 = color.RGB(0, 153, 253)
	blue2 = color.RGB(102, 253, 255)

	ws2811.Init(18, 1024, brightness)
	// init background color everywhere
	for i := 0; i < 2; i++ {
		for j := 0; j < screenHeight*(2*screenWidth+motifWidth); j++ {
			Bitmap[i][j] = background
		}
	}

	Set(2, 18, 0, salmon)
	Set(2, 19, 0, salmon)
	Set(2, 20, 0, salmon)
	Set(2, 21, 0, salmon)
	Set(2, 22, 0, salmon)
	Set(2, 23, 0, salmon)
	Set(2, 24, 0, catHead)
	Set(2, 25, 0, salmon)
	Set(2, 26, 0, salmon)
	Set(2, 28, 0, catHead)

	Set(2, 18, 1, salmon)
	Set(2, 19, 1, pink2)
	Set(2, 20, 1, pink1)
	Set(2, 21, 1, pink2)
	Set(2, 22, 1, pink1)
	Set(2, 23, 1, catHead)
	Set(2, 24, 1, catHead)
	Set(2, 25, 1, catHead)
	Set(2, 26, 1, salmon)
	Set(2, 27, 1, catHead)
	Set(2, 28, 1, catHead)
	Set(2, 29, 1, catHead)

	Set(2, 18, 2, salmon)
	Set(2, 19, 2, pink1)
	Set(2, 20, 2, pink2)
	Set(2, 21, 2, pink1)
	Set(2, 22, 2, pink1)
	Set(2, 23, 2, catHead)
	Set(2, 24, 2, white)
	Set(2, 25, 2, black)
	Set(2, 26, 2, catHead)
	Set(2, 27, 2, white)
	Set(2, 28, 2, black)
	Set(2, 29, 2, catHead)

	Set(2, 18, 3, salmon)
	Set(2, 19, 3, pink1)
	Set(2, 20, 3, pink1)
	Set(2, 21, 3, pink1)
	Set(2, 22, 3, catHead)
	Set(2, 23, 3, catHead)
	Set(2, 24, 3, black)
	Set(2, 25, 3, black)
	Set(2, 26, 3, catHead)
	Set(2, 27, 3, black)
	Set(2, 28, 3, black)
	Set(2, 29, 3, catHead)
	Set(2, 30, 3, catHead)

	Set(2, 18, 4, salmon)
	Set(2, 19, 4, pink2)
	Set(2, 20, 4, pink1)
	Set(2, 21, 4, pink2)
	Set(2, 22, 4, catHead)
	Set(2, 23, 4, cheek)
	Set(2, 24, 4, cheek)
	Set(2, 25, 4, catHead)
	Set(2, 26, 4, black)
	Set(2, 27, 4, catHead)
	Set(2, 28, 4, catHead)
	Set(2, 29, 4, cheek)
	Set(2, 30, 4, cheek)

	Set(2, 18, 5, salmon)
	Set(2, 19, 5, pink1)
	Set(2, 20, 5, pink2)
	Set(2, 21, 5, pink1)
	Set(2, 22, 5, pink1)
	Set(2, 23, 5, cheek)
	Set(2, 24, 5, cheek)
	Set(2, 25, 5, black)
	Set(2, 26, 5, black)
	Set(2, 27, 5, black)
	Set(2, 28, 5, catHead)
	Set(2, 29, 5, cheek)
	Set(2, 30, 5, cheek)

	Set(2, 18, 6, salmon)
	Set(2, 19, 6, salmon)
	Set(2, 20, 6, salmon)
	Set(2, 21, 6, salmon)
	Set(2, 22, 6, salmon)
	Set(2, 23, 6, salmon)
	Set(2, 24, 6, catHead)
	Set(2, 25, 6, catHead)
	Set(2, 26, 6, catHead)
	Set(2, 27, 6, catHead)
	Set(2, 28, 6, catHead)

	Set(0, 18, 7, catPaw)
	Set(0, 21, 7, catPaw)
	Set(0, 25, 7, catPaw)
	Set(0, 28, 7, catPaw)

	Set(1, 18, 6, catPaw)
	Set(1, 17, 7, catPaw)
	Set(1, 20, 7, catPaw)
	Set(1, 24, 7, catPaw)
	Set(1, 27, 7, catPaw)

	for i := 0; i < 18; i++ {
		var delta int
		if i%6 < 3 {
			delta = 0
		} else {
			delta = 1
		}
		Set(0, i, 0+delta, red)
		Set(0, i, 1+delta, orange)
		Set(0, i, 2+delta, yellow)
		Set(0, i, 3+delta, green)
		Set(0, i, 4+delta, blue1)
		Set(0, i, 5+delta, blue2)

		Set(1, i, 1-delta, red)
		Set(1, i, 2-delta, orange)
		Set(1, i, 3-delta, yellow)
		Set(1, i, 4-delta, green)
		Set(1, i, 5-delta, blue1)
		Set(1, i, 6-delta, blue2)
	}
}

func Set(frame int, col int, line int, color color.Color) {
	if frame == 0 || frame == 2 {
		if col%2 == 0 {
			Bitmap[0][(col+128)*8+line] = color
		} else {
			Bitmap[0][(col+128)*8+7-line] = color
		}
	}
	if frame == 1 || frame == 2 {
		if col%2 == 0 {
			Bitmap[1][(col+128)*8+7-line] = color
		} else {
			Bitmap[1][(col+128)*8+line] = color
		}
	}
}

func Roll() {
	for i := 0; i < (2*screenWidth+motifWidth)/2; i++ {
		ws2811.SetBitmap(Bitmap[0][i*16 : i*16+1024])
		ws2811.Render()
		ws2811.Wait()

		ws2811.SetBitmap(Bitmap[1][i*16+8 : i*16+8+1024])
		ws2811.Render()
		ws2811.Wait()
	}
}
