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
	"testing"
)

func TestColors(t *testing.T) {
	c := Red
	if c != 0xff0000 {
		t.Error("Bad red")
	}
	c = Green
	if c != 0x00ff00 {
		t.Error("Bad green")
	}
	c = Blue
	if c != 0x0000ff {
		t.Error("Bad blue")
	}
	c = White
	if c != 0xffffff {
		t.Error("Bad white")
	}
	c = Black
	if c != 0x000000 {
		t.Error("Bad black")
	}
}

func testException(t *testing.T, r, g, b int) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Nothing to recover")
		}
	}()
	RGB(r, g, b)
	t.Error("Exception not catched")
}

func TestException(t *testing.T) {
	testException(t, 0, 0, 256)
	testException(t, 0, 256, 0)
	testException(t, 256, 0, 0)
	testException(t, 0, 0, -1)
	testException(t, 0, -1, 0)
	testException(t, -1, 0, 0)
}
