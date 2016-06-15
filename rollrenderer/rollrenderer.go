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

// 2015-07-29 | JS | First version
// 2015-11-18 | JS | Latest version

//
// Telecom Tower rendering server
//
package rollrenderer

import (
	"fmt"
	"github.com/heia-fr/telecom-tower/ledmatrix"
	"github.com/heia-fr/telecom-tower/ledmatrix/font"
	"github.com/heia-fr/telecom-tower/tower"
)

type Line struct {
	Text  string
	Font  int
	Color string
}

type TextMessage struct {
	Body         []Line
	Introduction []Line
	Conclusion   []Line
	Separator    []Line
}

type BitmapMessage struct {
	Matrix     *ledmatrix.Matrix
	Preamble   int
	Checkpoint int
}

// renderdeTextMessage convert a text to a LED bitmap matrix.
// This method is used by the displayBuilder gorouting.
func textToBitmap(text []Line, bg uint32) *ledmatrix.Matrix {
	matrix := ledmatrix.NewMatrix(tower.Rows, 0)
	writer := ledmatrix.NewWriter(matrix)
	writer.Spacer(matrix.Columns, 0) // Blank bootstrap

	for _, line := range text {
		var r, g, b int
		var f font.Font

		if line.Font < 8 {
			f = font.Font6x8
		} else {
			f = font.Font8x8
		}

		fmt.Sscanf(line.Color, "#%02x%02x%02x", &r, &g, &b)
		writer.WriteText(line.Text, f, ledmatrix.RGB(r, g, b), bg)
	}
	return matrix
}

// blank generates a "space" from the given color
func blank(len int, bg uint32) *ledmatrix.Matrix {
	matrix := ledmatrix.NewMatrix(tower.Rows, tower.Columns)
	writer := ledmatrix.NewWriter(matrix)
	writer.Spacer(tower.Columns, bg) // Blank spacer
	return matrix
}

var bg = ledmatrix.RGB(0, 0, 0)
var preamble = blank(tower.Columns, bg)

func RenderMessage(message *TextMessage) *BitmapMessage {
	// If body is nil, then reset preamble. This is used when static
	// frames are sent to the tower and the next rolling message should
	// start with a blank.
	if message.Body == nil {
		preamble = blank(tower.Columns, bg)
		return nil
	}

	// render all parts of the message
	body := textToBitmap(message.Body, bg)
	introduction := textToBitmap(message.Introduction, bg)
	conclusion := textToBitmap(message.Conclusion, bg)
	separator := textToBitmap(message.Separator, bg)

	if body.Columns == 0 { // Skip empty message
		return nil
	}

	// we need at least a full display. So if the message is too short,
	// then append tbe separator and aother copy of the message.
	// We re-render the message and this is not optimal, but this
	// is only usd for shirt messages, so this is OK.
	for body.Columns < tower.Columns {
		body.Append(separator, textToBitmap(message.Body, bg))
	}

	// Append the introduction to the preamble
	preamble.Append(introduction)

	// Compute the wrapper. The wrapper makes the link between the end
	// and the start of the message, allowing it to "roll" properly.
	wrapper := ledmatrix.Concat(separator, body.Slice(0, tower.Columns))

	// Send the bitmap message
	result := &BitmapMessage{
		Matrix:     ledmatrix.Concat(preamble, body, wrapper),
		Preamble:   preamble.Columns,
		Checkpoint: preamble.Columns + body.Columns - tower.Columns,
	}

	// Compute the next preamble using the last frame of the body and
	// the conclusion
	preamble = ledmatrix.Concat(
		body.Slice(body.Columns-tower.Columns,
			body.Columns), conclusion)
	return result
}
