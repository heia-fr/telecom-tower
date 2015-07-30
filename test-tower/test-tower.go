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

// Test programm for Telecom Tower
package main

import (
	"github.com/heia-fr/telecom-tower/bitmapfont"
	"github.com/heia-fr/telecom-tower/color"
	"github.com/heia-fr/telecom-tower/tower"
)

func main() {
	tower.Init(128)
	defer tower.Shutdown()

	tower.Write("Haute école d'ingénierie et d'architecture Fribourg", bitmapfont.F88, color.RGB(150, 200, 255))
	tower.Write(" \u2665 ", bitmapfont.F88, color.RGB(255, 0, 0))
	tower.Write("Informatique", bitmapfont.F68, color.RGB(223, 0, 30))
	tower.Write(" \u2665 ", bitmapfont.F88, color.RGB(255, 0, 0))
	tower.Write("Télécommunications", bitmapfont.F68, color.RGB(248, 179, 30))
	tower.Write(" \u2665 ", bitmapfont.F88, color.RGB(255, 0, 0))

	for true {
		tower.Roll()
	}
}
