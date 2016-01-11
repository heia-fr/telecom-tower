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

// Bitmap fonts
package font

import (
	"bytes"
)

var alias = map[rune]string{
	0x2764:     "\u2665", // ❤
	0x0001F601: ":|",     // 😁
	0x0001F602: ":)",     // 😂
	0x0001F603: ":D",     // 😃
}

type Font struct {
	width  int
	height int
	bitmap map[rune][]byte
}

func ExpandAlias(text string) string {
	var f func(b *bytes.Buffer, s string)
	f = func(b *bytes.Buffer, s string) {
		for _, c := range s {
			m, ok := alias[c]
			if ok {
				f(b, m)
			} else {
				b.WriteRune(c)
			}
		}
	}
	b := new(bytes.Buffer)
	f(b, text)
	return b.String()
}

func (f *Font) Height() int {
	return f.height
}

func (f *Font) Width() int {
	return f.width
}

func (f *Font) Bitmap(c rune) []byte {
	return f.bitmap[c]
}
