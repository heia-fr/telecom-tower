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

type Color uint32

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
