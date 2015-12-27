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

package font

import (
	"fmt"
	"testing"
)

func TestFontSize(t *testing.T) {
	h := Font6x8.Height()
	if h != 8 {
		t.Errorf("Height if font 6x8 is %d and should be 8", h)
	}
	w := Font6x8.Width()
	if w != 6 {
		t.Errorf("Width if font 6x8 is %d and should be 6", w)
	}

	h = Font8x8.Height()
	if h != 8 {
		t.Errorf("Height if font 8x8 is %d and should be 8", h)
	}
	w = Font8x8.Width()
	if w != 8 {
		t.Errorf("Width if font 8x8 is %d and should be 8", w)
	}
}

func TestBitmaps(t *testing.T) {
	for _, f := range []Font{Font6x8, Font8x8} {
		b := f.Bitmap('A')
		if len(b) != f.Width() {
			t.Errorf("bitmap size mismatch. %d instead of %d", len(b), f.Width())
		}
	}
}

func ExampleA() {
	f := Font8x8
	c := f.Bitmap('A')
	for i := 0; i < f.Height(); i++ {
		for j := 0; j < f.Width(); j++ {
			if c[j]&(1<<uint(i)) != 0 {
				fmt.Print("*")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
	// Output:
	// __**____
	// _****___
	// **__**__
	// **__**__
	// ******__
	// **__**__
	// **__**__
	// ________
}
