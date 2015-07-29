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

// Package cp850 provides functions to convert unicode string to
// "Code Page 850". This package was done to use simple bitmap fonts
package cp850

func StringToCp850(s string) []byte {
	sa := []rune(s)
	res := make([]byte, len(sa))
	for i := 0; i < len(sa); i++ {
		res[i] = cp850[sa[i]]
	}
	return res
}
