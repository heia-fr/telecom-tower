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

// Credit: Thank you Benedikt K. for the original bitmap font files.
//         http://www.mikrocontroller.net/topic/54860

// Bitmap font 8x8
package font

var Font8x8 = Font{
	width:  8,
	height: 8,
	bitmap: map[rune][]byte{
		0x0020: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // " "
		0x0021: {0x00, 0x06, 0x5f, 0x5f, 0x06, 0x00, 0x00, 0x00}, // "!"
		0x0022: {0x00, 0x07, 0x07, 0x00, 0x07, 0x07, 0x00, 0x00}, // """
		0x0023: {0x14, 0x7f, 0x7f, 0x14, 0x7f, 0x7f, 0x14, 0x00}, // "#"
		0x0024: {0x24, 0x2e, 0x6b, 0x6b, 0x3a, 0x12, 0x00, 0x00}, // "$"
		0x0025: {0x46, 0x66, 0x30, 0x18, 0x0c, 0x66, 0x62, 0x00}, // "%"
		0x0026: {0x30, 0x7a, 0x4f, 0x5d, 0x37, 0x7a, 0x48, 0x00}, // "&"
		0x0027: {0x04, 0x07, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00}, // "'"
		0x0028: {0x00, 0x1c, 0x3e, 0x63, 0x41, 0x00, 0x00, 0x00}, // "("
		0x0029: {0x00, 0x41, 0x63, 0x3e, 0x1c, 0x00, 0x00, 0x00}, // ")"
		0x002a: {0x08, 0x2a, 0x3e, 0x1c, 0x1c, 0x3e, 0x2a, 0x08}, // "*"
		0x002b: {0x08, 0x08, 0x3e, 0x3e, 0x08, 0x08, 0x00, 0x00}, // "+"
		0x002c: {0x00, 0xa0, 0xe0, 0x60, 0x00, 0x00, 0x00, 0x00}, // ","
		0x002d: {0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x00, 0x00}, // "-"
		0x002e: {0x00, 0x00, 0x60, 0x60, 0x00, 0x00, 0x00, 0x00}, // "."
		0x002f: {0x60, 0x30, 0x18, 0x0c, 0x06, 0x03, 0x01, 0x00}, // "/"
		0x0030: {0x3e, 0x7f, 0x59, 0x4d, 0x7f, 0x3e, 0x00, 0x00}, // "0"
		0x0031: {0x42, 0x42, 0x7f, 0x7f, 0x40, 0x40, 0x00, 0x00}, // "1"
		0x0032: {0x62, 0x73, 0x59, 0x49, 0x6f, 0x66, 0x00, 0x00}, // "2"
		0x0033: {0x22, 0x63, 0x49, 0x49, 0x7f, 0x36, 0x00, 0x00}, // "3"
		0x0034: {0x18, 0x1c, 0x16, 0x13, 0x7f, 0x7f, 0x10, 0x00}, // "4"
		0x0035: {0x27, 0x67, 0x45, 0x45, 0x7d, 0x39, 0x00, 0x00}, // "5"
		0x0036: {0x3c, 0x7e, 0x4b, 0x49, 0x79, 0x30, 0x00, 0x00}, // "6"
		0x0037: {0x03, 0x63, 0x71, 0x19, 0x0f, 0x07, 0x00, 0x00}, // "7"
		0x0038: {0x36, 0x7f, 0x49, 0x49, 0x7f, 0x36, 0x00, 0x00}, // "8"
		0x0039: {0x06, 0x4f, 0x49, 0x69, 0x3f, 0x1e, 0x00, 0x00}, // "9"
		0x003a: {0x00, 0x00, 0x6c, 0x6c, 0x00, 0x00, 0x00, 0x00}, // ":"
		0x003b: {0x00, 0xa0, 0xec, 0x6c, 0x00, 0x00, 0x00, 0x00}, // ";"
		0x003c: {0x08, 0x1c, 0x36, 0x63, 0x41, 0x00, 0x00, 0x00}, // "<"
		0x003d: {0x14, 0x14, 0x14, 0x14, 0x14, 0x14, 0x00, 0x00}, // "="
		0x003e: {0x00, 0x41, 0x63, 0x36, 0x1c, 0x08, 0x00, 0x00}, // ">"
		0x003f: {0x02, 0x03, 0x51, 0x59, 0x0f, 0x06, 0x00, 0x00}, // "?"
		0x0040: {0x3e, 0x7f, 0x41, 0x5d, 0x5d, 0x1f, 0x1e, 0x00}, // "@"
		0x0041: {0x7c, 0x7e, 0x13, 0x13, 0x7e, 0x7c, 0x00, 0x00}, // "A"
		0x0042: {0x41, 0x7f, 0x7f, 0x49, 0x49, 0x7f, 0x36, 0x00}, // "B"
		0x0043: {0x1c, 0x3e, 0x63, 0x41, 0x41, 0x63, 0x22, 0x00}, // "C"
		0x0044: {0x41, 0x7f, 0x7f, 0x41, 0x63, 0x7f, 0x1c, 0x00}, // "D"
		0x0045: {0x41, 0x7f, 0x7f, 0x49, 0x5d, 0x41, 0x63, 0x00}, // "E"
		0x0046: {0x41, 0x7f, 0x7f, 0x49, 0x1d, 0x01, 0x03, 0x00}, // "F"
		0x0047: {0x1c, 0x3e, 0x63, 0x41, 0x51, 0x73, 0x72, 0x00}, // "G"
		0x0048: {0x7f, 0x7f, 0x08, 0x08, 0x7f, 0x7f, 0x00, 0x00}, // "H"
		0x0049: {0x00, 0x41, 0x7f, 0x7f, 0x41, 0x00, 0x00, 0x00}, // "I"
		0x004a: {0x30, 0x70, 0x40, 0x41, 0x7f, 0x3f, 0x01, 0x00}, // "J"
		0x004b: {0x41, 0x7f, 0x7f, 0x08, 0x1c, 0x77, 0x63, 0x00}, // "K"
		0x004c: {0x41, 0x7f, 0x7f, 0x41, 0x40, 0x60, 0x70, 0x00}, // "L"
		0x004d: {0x7f, 0x7f, 0x06, 0x0c, 0x06, 0x7f, 0x7f, 0x00}, // "M"
		0x004e: {0x7f, 0x7f, 0x06, 0x0c, 0x18, 0x7f, 0x7f, 0x00}, // "N"
		0x004f: {0x1c, 0x3e, 0x63, 0x41, 0x63, 0x3e, 0x1c, 0x00}, // "O"
		0x0050: {0x41, 0x7f, 0x7f, 0x49, 0x09, 0x0f, 0x06, 0x00}, // "P"
		0x0051: {0x1e, 0x3f, 0x21, 0x71, 0x7f, 0x5e, 0x00, 0x00}, // "Q"
		0x0052: {0x41, 0x7f, 0x7f, 0x19, 0x39, 0x6f, 0x46, 0x00}, // "R"
		0x0053: {0x26, 0x67, 0x4d, 0x59, 0x7b, 0x32, 0x00, 0x00}, // "S"
		0x0054: {0x03, 0x41, 0x7f, 0x7f, 0x41, 0x03, 0x00, 0x00}, // "T"
		0x0055: {0x7f, 0x7f, 0x40, 0x40, 0x7f, 0x7f, 0x00, 0x00}, // "U"
		0x0056: {0x1f, 0x3f, 0x60, 0x60, 0x3f, 0x1f, 0x00, 0x00}, // "V"
		0x0057: {0x7f, 0x7f, 0x30, 0x18, 0x30, 0x7f, 0x7f, 0x00}, // "W"
		0x0058: {0x63, 0x77, 0x1c, 0x08, 0x1c, 0x77, 0x63, 0x00}, // "X"
		0x0059: {0x07, 0x4f, 0x78, 0x78, 0x4f, 0x07, 0x00, 0x00}, // "Y"
		0x005a: {0x67, 0x73, 0x59, 0x4d, 0x47, 0x63, 0x71, 0x00}, // "Z"
		0x005b: {0x00, 0x7f, 0x7f, 0x41, 0x41, 0x00, 0x00, 0x00}, // "["
		0x005c: {0x01, 0x03, 0x06, 0x0c, 0x18, 0x30, 0x60, 0x00}, // "\"
		0x005d: {0x00, 0x41, 0x41, 0x7f, 0x7f, 0x00, 0x00, 0x00}, // "]"
		0x005e: {0x08, 0x0c, 0x06, 0x03, 0x06, 0x0c, 0x08, 0x00}, // "^"
		0x005f: {0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, // "_"
		0x0060: {0x00, 0x00, 0x03, 0x07, 0x04, 0x00, 0x00, 0x00}, // "`"
		0x0061: {0x20, 0x74, 0x54, 0x54, 0x3c, 0x78, 0x40, 0x00}, // "a"
		0x0062: {0x41, 0x3f, 0x7f, 0x44, 0x44, 0x7c, 0x38, 0x00}, // "b"
		0x0063: {0x38, 0x7c, 0x44, 0x44, 0x6c, 0x28, 0x00, 0x00}, // "c"
		0x0064: {0x30, 0x78, 0x48, 0x49, 0x3f, 0x7f, 0x40, 0x00}, // "d"
		0x0065: {0x38, 0x7c, 0x54, 0x54, 0x5c, 0x18, 0x00, 0x00}, // "e"
		0x0066: {0x48, 0x7e, 0x7f, 0x49, 0x03, 0x02, 0x00, 0x00}, // "f"
		0x0067: {0x98, 0xbc, 0xa4, 0xa4, 0xf8, 0x7c, 0x04, 0x00}, // "g"
		0x0068: {0x41, 0x7f, 0x7f, 0x08, 0x04, 0x7c, 0x78, 0x00}, // "h"
		0x0069: {0x00, 0x44, 0x7d, 0x7d, 0x40, 0x00, 0x00, 0x00}, // "i"
		0x006a: {0x40, 0xc4, 0x84, 0xfd, 0x7d, 0x00, 0x00, 0x00}, // "j"
		0x006b: {0x41, 0x7f, 0x7f, 0x10, 0x38, 0x6c, 0x44, 0x00}, // "k"
		0x006c: {0x00, 0x41, 0x7f, 0x7f, 0x40, 0x00, 0x00, 0x00}, // "l"
		0x006d: {0x7c, 0x7c, 0x0c, 0x18, 0x0c, 0x7c, 0x78, 0x00}, // "m"
		0x006e: {0x7c, 0x7c, 0x04, 0x04, 0x7c, 0x78, 0x00, 0x00}, // "n"
		0x006f: {0x38, 0x7c, 0x44, 0x44, 0x7c, 0x38, 0x00, 0x00}, // "o"
		0x0070: {0x84, 0xfc, 0xf8, 0xa4, 0x24, 0x3c, 0x18, 0x00}, // "p"
		0x0071: {0x18, 0x3c, 0x24, 0xa4, 0xf8, 0xfc, 0x84, 0x00}, // "q"
		0x0072: {0x44, 0x7c, 0x78, 0x44, 0x1c, 0x18, 0x00, 0x00}, // "r"
		0x0073: {0x48, 0x5c, 0x54, 0x54, 0x74, 0x24, 0x00, 0x00}, // "s"
		0x0074: {0x00, 0x04, 0x3e, 0x7f, 0x44, 0x24, 0x00, 0x00}, // "t"
		0x0075: {0x3c, 0x7c, 0x40, 0x40, 0x3c, 0x7c, 0x40, 0x00}, // "u"
		0x0076: {0x1c, 0x3c, 0x60, 0x60, 0x3c, 0x1c, 0x00, 0x00}, // "v"
		0x0077: {0x3c, 0x7c, 0x60, 0x30, 0x60, 0x7c, 0x3c, 0x00}, // "w"
		0x0078: {0x44, 0x6c, 0x38, 0x10, 0x38, 0x6c, 0x44, 0x00}, // "x"
		0x0079: {0x9c, 0xbc, 0xa0, 0xa0, 0xfc, 0x7c, 0x00, 0x00}, // "y"
		0x007a: {0x4c, 0x64, 0x74, 0x5c, 0x4c, 0x64, 0x00, 0x00}, // "z"
		0x007b: {0x08, 0x08, 0x3e, 0x77, 0x41, 0x41, 0x00, 0x00}, // "{"
		0x007c: {0x00, 0x00, 0x00, 0x77, 0x77, 0x00, 0x00, 0x00}, // "|"
		0x007d: {0x41, 0x41, 0x77, 0x3e, 0x08, 0x08, 0x00, 0x00}, // "}"
		0x007e: {0x02, 0x03, 0x01, 0x03, 0x02, 0x03, 0x01, 0x00}, // "~"
		0x00a0: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // " "
		0x00a1: {0x00, 0x00, 0x60, 0xfa, 0xfa, 0x60, 0x00, 0x00}, // "¡"
		0x00a2: {0x18, 0x3c, 0x24, 0xe7, 0xe7, 0x24, 0x24, 0x00}, // "¢"
		0x00a3: {0x68, 0x7e, 0x7f, 0x49, 0x43, 0x66, 0x20, 0x00}, // "£"
		0x00a4: {0x66, 0x3c, 0x3c, 0x24, 0x3c, 0x3c, 0x66, 0x00}, // "¤"
		0x00a5: {0x2b, 0x2f, 0xfc, 0xfc, 0x2f, 0x2b, 0x00, 0x00}, // "¥"
		0x00a6: {0x00, 0x00, 0x00, 0x77, 0x77, 0x00, 0x00, 0x00}, // "¦"
		0x00a7: {0xda, 0xbf, 0xa5, 0xa5, 0xfd, 0x59, 0x03, 0x02}, // "§"
		0x00a8: {0x02, 0x02, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00}, // "¨"
		0x00a9: {0x1c, 0x22, 0x5d, 0x55, 0x55, 0x41, 0x22, 0x1c}, // "©"
		0x00aa: {0x00, 0x26, 0x2f, 0x29, 0x2f, 0x2f, 0x28, 0x00}, // "ª"
		0x00ab: {0x08, 0x1c, 0x36, 0x22, 0x08, 0x1c, 0x36, 0x22}, // "«"
		0x00ac: {0x08, 0x08, 0x08, 0x08, 0x38, 0x38, 0x00, 0x00}, // "¬"
		0x00ad: {0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x00, 0x00}, // "­"
		0x00ae: {0x1c, 0x22, 0x7d, 0x4b, 0x5b, 0x65, 0x22, 0x1c}, // "®"
		0x00af: {0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x00, 0x00}, // "¯"
		0x00b0: {0x00, 0x06, 0x0f, 0x09, 0x0f, 0x06, 0x00, 0x00}, // "°"
		0x00b1: {0x44, 0x44, 0x5f, 0x5f, 0x44, 0x44, 0x00, 0x00}, // "±"
		0x00b2: {0x00, 0x19, 0x1d, 0x15, 0x17, 0x12, 0x00, 0x00}, // "²"
		0x00b3: {0x00, 0x11, 0x15, 0x15, 0x1f, 0x1f, 0x0a, 0x00}, // "³"
		0x00b4: {0x00, 0x00, 0x02, 0x03, 0x01, 0x00, 0x00, 0x00}, // "´"
		0x00b5: {0x80, 0xfe, 0x7e, 0x20, 0x20, 0x3e, 0x1e, 0x00}, // "µ"
		0x00b6: {0x06, 0x0f, 0x09, 0x7f, 0x7f, 0x01, 0x7f, 0x7f}, // "¶"
		0x00b7: {0x00, 0x00, 0x00, 0x10, 0x10, 0x00, 0x00, 0x00}, // "·"
		0x00b8: {0x00, 0x80, 0xc0, 0x40, 0x00, 0x00, 0x00, 0x00}, // "¸"
		0x00b9: {0x00, 0x12, 0x13, 0x1f, 0x1f, 0x10, 0x10, 0x00}, // "¹"
		0x00ba: {0x00, 0x26, 0x2f, 0x29, 0x29, 0x2f, 0x26, 0x00}, // "º"
		0x00bb: {0x22, 0x36, 0x1c, 0x08, 0x22, 0x36, 0x1c, 0x08}, // "»"
		0x00bc: {0x61, 0x3f, 0x1f, 0x4c, 0x66, 0x73, 0xd9, 0xf8}, // "¼"
		0x00bd: {0x61, 0x3f, 0x1f, 0xcc, 0xee, 0xab, 0xb9, 0x90}, // "½"
		0x00be: {0x71, 0x35, 0x1f, 0x4c, 0x66, 0x73, 0xd9, 0xf8}, // "¾"
		0x00bf: {0x30, 0x78, 0x4d, 0x45, 0x60, 0x20, 0x00, 0x00}, // "¿"
		0x00c0: {0x71, 0x7b, 0x2e, 0x2c, 0x78, 0x70, 0x00, 0x00}, // "À"
		0x00c1: {0x70, 0x78, 0x2c, 0x2e, 0x7b, 0x71, 0x00, 0x00}, // "Á"
		0x00c2: {0x72, 0x79, 0x2d, 0x2d, 0x79, 0x72, 0x00, 0x00}, // "Â"
		0x00c3: {0x72, 0x7b, 0x2d, 0x2f, 0x7a, 0x73, 0x01, 0x00}, // "Ã"
		0x00c4: {0x79, 0x7d, 0x26, 0x26, 0x7d, 0x79, 0x00, 0x00}, // "Ä"
		0x00c5: {0x70, 0x7a, 0x2d, 0x2d, 0x7a, 0x70, 0x00, 0x00}, // "Å"
		0x00c6: {0x7c, 0x7e, 0x0b, 0x09, 0x7f, 0x7f, 0x49, 0x00}, // "Æ"
		0x00c7: {0x1e, 0xbf, 0xe1, 0x61, 0x33, 0x12, 0x00, 0x00}, // "Ç"
		0x00c8: {0x44, 0x7d, 0x7f, 0x56, 0x54, 0x44, 0x00, 0x00}, // "È"
		0x00c9: {0x44, 0x7c, 0x7e, 0x57, 0x55, 0x44, 0x00, 0x00}, // "É"
		0x00ca: {0x46, 0x7d, 0x7d, 0x55, 0x55, 0x46, 0x00, 0x00}, // "Ê"
		0x00cb: {0x45, 0x7d, 0x7c, 0x54, 0x55, 0x45, 0x00, 0x00}, // "Ë"
		0x00cc: {0x00, 0x45, 0x7f, 0x7e, 0x44, 0x00, 0x00, 0x00}, // "Ì"
		0x00cd: {0x00, 0x44, 0x7e, 0x7f, 0x45, 0x00, 0x00, 0x00}, // "Í"
		0x00ce: {0x02, 0x45, 0x7d, 0x7d, 0x45, 0x02, 0x00, 0x00}, // "Î"
		0x00cf: {0x01, 0x45, 0x7c, 0x7c, 0x45, 0x01, 0x00, 0x00}, // "Ï"
		0x00d0: {0x49, 0x7f, 0x7f, 0x49, 0x63, 0x7f, 0x1c, 0x00}, // "Ð"
		0x00d1: {0x7a, 0x7b, 0x19, 0x33, 0x7a, 0x7b, 0x01, 0x00}, // "Ñ"
		0x00d2: {0x38, 0x7c, 0x45, 0x47, 0x46, 0x7c, 0x38, 0x00}, // "Ò"
		0x00d3: {0x38, 0x7c, 0x46, 0x47, 0x45, 0x7c, 0x38, 0x00}, // "Ó"
		0x00d4: {0x3a, 0x7d, 0x45, 0x45, 0x45, 0x7d, 0x3a, 0x00}, // "Ô"
		0x00d5: {0x3a, 0x7f, 0x45, 0x47, 0x46, 0x7f, 0x39, 0x00}, // "Õ"
		0x00d6: {0x39, 0x7d, 0x44, 0x44, 0x44, 0x7d, 0x39, 0x00}, // "Ö"
		0x00d7: {0x44, 0x6c, 0x38, 0x38, 0x6c, 0x44, 0x00, 0x00}, // "×"
		0x00d8: {0x5c, 0x3e, 0x73, 0x49, 0x67, 0x3e, 0x1d, 0x00}, // "Ø"
		0x00d9: {0x3c, 0x7d, 0x43, 0x42, 0x7c, 0x3c, 0x00, 0x00}, // "Ù"
		0x00da: {0x3c, 0x7c, 0x42, 0x43, 0x7d, 0x3c, 0x00, 0x00}, // "Ú"
		0x00db: {0x3a, 0x79, 0x41, 0x41, 0x79, 0x3a, 0x00, 0x00}, // "Û"
		0x00dc: {0x3d, 0x7d, 0x40, 0x40, 0x7d, 0x3d, 0x00, 0x00}, // "Ü"
		0x00dd: {0x0c, 0x5c, 0x72, 0x73, 0x5d, 0x0c, 0x00, 0x00}, // "Ý"
		0x00de: {0x41, 0x7f, 0x7f, 0x55, 0x14, 0x1c, 0x08, 0x00}, // "Þ"
		0x00df: {0xfc, 0xfe, 0x2a, 0x2a, 0x3e, 0x14, 0x00, 0x00}, // "ß"
		0x00e0: {0x20, 0x75, 0x57, 0x56, 0x7c, 0x78, 0x40, 0x00}, // "à"
		0x00e1: {0x20, 0x74, 0x56, 0x57, 0x7d, 0x78, 0x40, 0x00}, // "á"
		0x00e2: {0x02, 0x23, 0x75, 0x55, 0x55, 0x7d, 0x7b, 0x42}, // "â"
		0x00e3: {0x22, 0x77, 0x55, 0x57, 0x7e, 0x7b, 0x41, 0x00}, // "ã"
		0x00e4: {0x21, 0x75, 0x54, 0x54, 0x7d, 0x79, 0x40, 0x00}, // "ä"
		0x00e5: {0x00, 0x22, 0x77, 0x55, 0x55, 0x7f, 0x7a, 0x40}, // "å"
		0x00e6: {0x20, 0x74, 0x54, 0x54, 0x7c, 0x7c, 0x54, 0x54}, // "æ"
		0x00e7: {0x1c, 0xbe, 0xe2, 0x62, 0x36, 0x14, 0x00, 0x00}, // "ç"
		0x00e8: {0x38, 0x7d, 0x57, 0x56, 0x5c, 0x18, 0x00, 0x00}, // "è"
		0x00e9: {0x38, 0x7c, 0x56, 0x57, 0x5d, 0x18, 0x00, 0x00}, // "é"
		0x00ea: {0x02, 0x3b, 0x7d, 0x55, 0x55, 0x5d, 0x1b, 0x02}, // "ê"
		0x00eb: {0x39, 0x7d, 0x54, 0x54, 0x5d, 0x19, 0x00, 0x00}, // "ë"
		0x00ec: {0x00, 0x45, 0x7f, 0x7e, 0x40, 0x00, 0x00, 0x00}, // "ì"
		0x00ed: {0x00, 0x44, 0x7e, 0x7f, 0x41, 0x00, 0x00, 0x00}, // "í"
		0x00ee: {0x02, 0x03, 0x45, 0x7d, 0x7d, 0x43, 0x02, 0x00}, // "î"
		0x00ef: {0x01, 0x45, 0x7c, 0x7c, 0x41, 0x01, 0x00, 0x00}, // "ï"
		0x00f0: {0x05, 0x27, 0x72, 0x57, 0x7d, 0x38, 0x00, 0x00}, // "ð"
		0x00f1: {0x7a, 0x7b, 0x09, 0x0b, 0x7a, 0x73, 0x01, 0x00}, // "ñ"
		0x00f2: {0x30, 0x79, 0x4b, 0x4a, 0x78, 0x30, 0x00, 0x00}, // "ò"
		0x00f3: {0x30, 0x78, 0x48, 0x4a, 0x7b, 0x31, 0x00, 0x00}, // "ó"
		0x00f4: {0x32, 0x7b, 0x49, 0x49, 0x7b, 0x32, 0x00, 0x00}, // "ô"
		0x00f5: {0x32, 0x7b, 0x49, 0x4b, 0x7a, 0x33, 0x01, 0x00}, // "õ"
		0x00f6: {0x32, 0x7a, 0x48, 0x48, 0x7a, 0x32, 0x00, 0x00}, // "ö"
		0x00f7: {0x08, 0x08, 0x6b, 0x6b, 0x08, 0x08, 0x00, 0x00}, // "÷"
		0x00f8: {0x38, 0x7c, 0x64, 0x54, 0x4c, 0x7c, 0x38, 0x00}, // "ø"
		0x00f9: {0x38, 0x79, 0x43, 0x42, 0x78, 0x78, 0x40, 0x00}, // "ù"
		0x00fa: {0x38, 0x78, 0x40, 0x42, 0x7b, 0x79, 0x40, 0x00}, // "ú"
		0x00fb: {0x3a, 0x7b, 0x41, 0x41, 0x7b, 0x7a, 0x40, 0x00}, // "û"
		0x00fc: {0x3a, 0x7a, 0x40, 0x40, 0x7a, 0x7a, 0x40, 0x00}, // "ü"
		0x00fd: {0xb8, 0xb8, 0xa2, 0xa3, 0xf9, 0x78, 0x00, 0x00}, // "ý"
		0x00fe: {0x42, 0x7e, 0x7e, 0x54, 0x1c, 0x08, 0x00, 0x00}, // "þ"
		0x00ff: {0xba, 0xba, 0xa0, 0xa0, 0xfa, 0x7a, 0x00, 0x00}, // "ÿ"
		0x0131: {0x0a, 0x0e, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00}, // "ı"
		0x0192: {0x40, 0xc8, 0x88, 0xfe, 0x7f, 0x09, 0x0b, 0x02}, // "ƒ"
		0x2017: {0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x00, 0x00}, // "‗"
		0x2022: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "•"
		0x203c: {0x00, 0x5f, 0x5f, 0x00, 0x00, 0x5f, 0x5f, 0x00}, // "‼"
		0x2190: {0x08, 0x1c, 0x3e, 0x2a, 0x08, 0x08, 0x08, 0x00}, // "←"
		0x2191: {0x00, 0x04, 0x06, 0x7f, 0x7f, 0x06, 0x04, 0x00}, // "↑"
		0x2192: {0x08, 0x08, 0x08, 0x2a, 0x3e, 0x1c, 0x08, 0x00}, // "→"
		0x2193: {0x00, 0x10, 0x30, 0x7f, 0x7f, 0x30, 0x10, 0x00}, // "↓"
		0x2194: {0x08, 0x1c, 0x3e, 0x08, 0x08, 0x3e, 0x1c, 0x08}, // "↔"
		0x2195: {0x00, 0x24, 0x66, 0xff, 0xff, 0x66, 0x24, 0x00}, // "↕"
		0x21a8: {0x80, 0x94, 0xb6, 0xff, 0xff, 0xb6, 0x94, 0x80}, // "↨"
		0x221f: {0x3c, 0x3c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00}, // "∟"
		0x2302: {0x78, 0x7c, 0x46, 0x43, 0x46, 0x7c, 0x78, 0x00}, // "⌂"
		0x2500: {0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10}, // "─"
		0x2502: {0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0x00}, // "│"
		0x250c: {0x00, 0x00, 0x00, 0xf0, 0xf0, 0x10, 0x10, 0x10}, // "┌"
		0x2510: {0x10, 0x10, 0x10, 0xf0, 0xf0, 0x00, 0x00, 0x00}, // "┐"
		0x2514: {0x00, 0x00, 0x00, 0x1f, 0x1f, 0x10, 0x10, 0x10}, // "└"
		0x2518: {0x10, 0x10, 0x10, 0x1f, 0x1f, 0x00, 0x00, 0x00}, // "┘"
		0x251c: {0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x10, 0x10}, // "├"
		0x2524: {0x10, 0x10, 0x10, 0xff, 0xff, 0x00, 0x00, 0x00}, // "┤"
		0x252c: {0x10, 0x10, 0x10, 0xf0, 0xf0, 0x10, 0x10, 0x10}, // "┬"
		0x2534: {0x10, 0x10, 0x10, 0x1f, 0x1f, 0x10, 0x10, 0x10}, // "┴"
		0x253c: {0x10, 0x10, 0x10, 0xff, 0xff, 0x10, 0x10, 0x10}, // "┼"
		0x2550: {0x14, 0x14, 0x14, 0x14, 0x14, 0x14, 0x14, 0x14}, // "═"
		0x2551: {0x00, 0x00, 0xff, 0xff, 0x00, 0xff, 0xff, 0x00}, // "║"
		0x2554: {0x00, 0x00, 0xfc, 0xfc, 0x04, 0xf4, 0xf4, 0x14}, // "╔"
		0x2557: {0x14, 0x14, 0xf4, 0xf4, 0x04, 0xfc, 0xfc, 0x00}, // "╗"
		0x255a: {0x00, 0x00, 0x1f, 0x1f, 0x10, 0x17, 0x17, 0x14}, // "╚"
		0x255d: {0x14, 0x14, 0x17, 0x17, 0x10, 0x1f, 0x1f, 0x00}, // "╝"
		0x2560: {0x00, 0x00, 0xff, 0xff, 0x00, 0xf7, 0xf7, 0x14}, // "╠"
		0x2563: {0x14, 0x14, 0xf7, 0xf7, 0x00, 0xff, 0xff, 0x00}, // "╣"
		0x2566: {0x14, 0x14, 0xf4, 0xf4, 0x04, 0xf4, 0xf4, 0x14}, // "╦"
		0x2569: {0x14, 0x14, 0x17, 0x17, 0x10, 0x17, 0x17, 0x14}, // "╩"
		0x256c: {0x14, 0x14, 0xf7, 0xf7, 0x00, 0xf7, 0xf7, 0x14}, // "╬"
		0x2580: {0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f}, // "▀"
		0x2584: {0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0}, // "▄"
		0x2588: {0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // "█"
		0x2591: {0xaa, 0x00, 0x55, 0x00, 0xaa, 0x00, 0x55, 0x00}, // "░"
		0x2592: {0xaa, 0x55, 0xaa, 0x55, 0xaa, 0x55, 0xaa, 0x55}, // "▒"
		0x2593: {0x55, 0xff, 0xaa, 0xff, 0x55, 0xff, 0xaa, 0xff}, // "▓"
		0x25a0: {0x00, 0x00, 0x3c, 0x3c, 0x3c, 0x3c, 0x00, 0x00}, // "■"
		0x25ac: {0x00, 0x70, 0x70, 0x70, 0x70, 0x70, 0x70, 0x00}, // "▬"
		0x25b2: {0x30, 0x38, 0x3c, 0x3e, 0x3e, 0x3c, 0x38, 0x30}, // "▲"
		0x25ba: {0x7f, 0x3e, 0x3e, 0x1c, 0x1c, 0x08, 0x08, 0x00}, // "►"
		0x25bc: {0x06, 0x0e, 0x1e, 0x3e, 0x3e, 0x1e, 0x0e, 0x06}, // "▼"
		0x25c4: {0x08, 0x08, 0x1c, 0x1c, 0x3e, 0x3e, 0x7f, 0x00}, // "◄"
		0x25cb: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "○"
		0x25d8: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "◘"
		0x25d9: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "◙"
		0x263a: {0x7e, 0x81, 0x95, 0xb1, 0xb1, 0x95, 0x81, 0x7e}, // "☺"
		0x263b: {0x7e, 0xff, 0xeb, 0xcf, 0xcf, 0xeb, 0xff, 0x7e}, // "☻"
		0x263c: {0x99, 0x5a, 0x3c, 0xe7, 0xe7, 0x3c, 0x5a, 0x99}, // "☼"
		0x2640: {0x00, 0x4e, 0x5f, 0xf1, 0xf1, 0x5f, 0x4e, 0x00}, // "♀"
		0x2642: {0x70, 0xf8, 0x88, 0x88, 0xfd, 0x7f, 0x07, 0x0f}, // "♂"
		0x2660: {0x10, 0x38, 0xbc, 0xff, 0xbc, 0x38, 0x10, 0x00}, // "♠"
		0x2663: {0x38, 0x3a, 0x9f, 0xff, 0x9f, 0x3a, 0x38, 0x00}, // "♣"
		0x2665: {0x0e, 0x1f, 0x3f, 0x7e, 0x3f, 0x1f, 0x0e, 0x00}, // "♥"
		0x2666: {0x08, 0x1c, 0x3e, 0x7f, 0x3e, 0x1c, 0x08, 0x00}, // "♦"
		0x266a: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "♪"
		0x266b: {0xc0, 0xff, 0x7f, 0x05, 0x05, 0x65, 0x7f, 0x3f}, // "♫"
	},
}
