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

//
// package to render bitmap fonts
//
package font

type Font struct {
	Width  int
	Bitmap [][]byte
}

func (f Font) GetCharMSB(c byte) []byte {
	return f.Bitmap[c]
}

func (f Font) GetCharLSB(c byte) []byte {
	res := make([]byte, len(f.Bitmap[c]))
	copy(res, f.Bitmap[c])
	for i := 0; i < len(res); i++ {
		res[i] = bitReverse[res[i]]
	}
	return res
}

func (f Font) GetCharMLSB(c byte) []byte {
	res := make([]byte, len(f.Bitmap[c]))
	copy(res, f.Bitmap[c])
	for i := 0; i < len(res); i += 2 {
		res[i] = bitReverse[res[i]]
	}
	return res
}

func (f Font) GetCharLMSB(c byte) []byte {
	res := make([]byte, len(f.Bitmap[c]))
	copy(res, f.Bitmap[c])
	for i := 1; i < len(res); i += 2 {
		res[i] = bitReverse[res[i]]
	}
	return res
}
