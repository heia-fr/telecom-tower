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

package cp850

import (
	"bytes"
	"testing"
)

func TestBasic(t *testing.T) {

	if !bytes.Equal(StringToCp850("Tchô técoles"), []byte{84, 99, 104, 147, 32, 116, 130, 99, 111, 108, 101, 115}) {
		t.Error("Fail basic tests")
	}

}
