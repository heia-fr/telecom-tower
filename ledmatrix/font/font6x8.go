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

// Bitmap font 6x8
package font

var Font6x8 = Font{
	width:  6,
	height: 8,
	bitmap: map[rune][]byte{
		0x0020: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // " "
		0x0021: {0x00, 0x00, 0x06, 0x5f, 0x06, 0x00}, // "!"
		0x0022: {0x00, 0x07, 0x03, 0x00, 0x07, 0x03}, // """
		0x0023: {0x00, 0x24, 0x7e, 0x24, 0x7e, 0x24}, // "#"
		0x0024: {0x00, 0x24, 0x2b, 0x6a, 0x12, 0x00}, // "$"
		0x0025: {0x00, 0x63, 0x13, 0x08, 0x64, 0x63}, // "%"
		0x0026: {0x00, 0x36, 0x49, 0x56, 0x20, 0x50}, // "&"
		0x0027: {0x00, 0x00, 0x07, 0x03, 0x00, 0x00}, // "'"
		0x0028: {0x00, 0x00, 0x3e, 0x41, 0x00, 0x00}, // "("
		0x0029: {0x00, 0x00, 0x41, 0x3e, 0x00, 0x00}, // ")"
		0x002a: {0x00, 0x08, 0x3e, 0x1c, 0x3e, 0x08}, // "*"
		0x002b: {0x00, 0x08, 0x08, 0x3e, 0x08, 0x08}, // "+"
		0x002c: {0x00, 0x00, 0xe0, 0x60, 0x00, 0x00}, // ","
		0x002d: {0x00, 0x08, 0x08, 0x08, 0x08, 0x08}, // "-"
		0x002e: {0x00, 0x00, 0x60, 0x60, 0x00, 0x00}, // "."
		0x002f: {0x00, 0x20, 0x10, 0x08, 0x04, 0x02}, // "/"
		0x0030: {0x00, 0x3e, 0x51, 0x49, 0x45, 0x3e}, // "0"
		0x0031: {0x00, 0x00, 0x42, 0x7f, 0x40, 0x00}, // "1"
		0x0032: {0x00, 0x62, 0x51, 0x49, 0x49, 0x46}, // "2"
		0x0033: {0x00, 0x22, 0x49, 0x49, 0x49, 0x36}, // "3"
		0x0034: {0x00, 0x18, 0x14, 0x12, 0x7f, 0x10}, // "4"
		0x0035: {0x00, 0x2f, 0x49, 0x49, 0x49, 0x31}, // "5"
		0x0036: {0x00, 0x3c, 0x4a, 0x49, 0x49, 0x30}, // "6"
		0x0037: {0x00, 0x01, 0x71, 0x09, 0x05, 0x03}, // "7"
		0x0038: {0x00, 0x36, 0x49, 0x49, 0x49, 0x36}, // "8"
		0x0039: {0x00, 0x06, 0x49, 0x49, 0x29, 0x1e}, // "9"
		0x003a: {0x00, 0x00, 0x6c, 0x6c, 0x00, 0x00}, // ":"
		0x003b: {0x00, 0x00, 0xec, 0x6c, 0x00, 0x00}, // ";"
		0x003c: {0x00, 0x08, 0x14, 0x22, 0x41, 0x00}, // "<"
		0x003d: {0x00, 0x24, 0x24, 0x24, 0x24, 0x24}, // "="
		0x003e: {0x00, 0x00, 0x41, 0x22, 0x14, 0x08}, // ">"
		0x003f: {0x00, 0x02, 0x01, 0x59, 0x09, 0x06}, // "?"
		0x0040: {0x00, 0x3e, 0x41, 0x5d, 0x55, 0x1e}, // "@"
		0x0041: {0x00, 0x7e, 0x11, 0x11, 0x11, 0x7e}, // "A"
		0x0042: {0x00, 0x7f, 0x49, 0x49, 0x49, 0x36}, // "B"
		0x0043: {0x00, 0x3e, 0x41, 0x41, 0x41, 0x22}, // "C"
		0x0044: {0x00, 0x7f, 0x41, 0x41, 0x41, 0x3e}, // "D"
		0x0045: {0x00, 0x7f, 0x49, 0x49, 0x49, 0x41}, // "E"
		0x0046: {0x00, 0x7f, 0x09, 0x09, 0x09, 0x01}, // "F"
		0x0047: {0x00, 0x3e, 0x41, 0x49, 0x49, 0x7a}, // "G"
		0x0048: {0x00, 0x7f, 0x08, 0x08, 0x08, 0x7f}, // "H"
		0x0049: {0x00, 0x00, 0x41, 0x7f, 0x41, 0x00}, // "I"
		0x004a: {0x00, 0x30, 0x40, 0x40, 0x40, 0x3f}, // "J"
		0x004b: {0x00, 0x7f, 0x08, 0x14, 0x22, 0x41}, // "K"
		0x004c: {0x00, 0x7f, 0x40, 0x40, 0x40, 0x40}, // "L"
		0x004d: {0x00, 0x7f, 0x02, 0x04, 0x02, 0x7f}, // "M"
		0x004e: {0x00, 0x7f, 0x02, 0x04, 0x08, 0x7f}, // "N"
		0x004f: {0x00, 0x3e, 0x41, 0x41, 0x41, 0x3e}, // "O"
		0x0050: {0x00, 0x7f, 0x09, 0x09, 0x09, 0x06}, // "P"
		0x0051: {0x00, 0x3e, 0x41, 0x51, 0x21, 0x5e}, // "Q"
		0x0052: {0x00, 0x7f, 0x09, 0x09, 0x19, 0x66}, // "R"
		0x0053: {0x00, 0x26, 0x49, 0x49, 0x49, 0x32}, // "S"
		0x0054: {0x00, 0x01, 0x01, 0x7f, 0x01, 0x01}, // "T"
		0x0055: {0x00, 0x3f, 0x40, 0x40, 0x40, 0x3f}, // "U"
		0x0056: {0x00, 0x1f, 0x20, 0x40, 0x20, 0x1f}, // "V"
		0x0057: {0x00, 0x3f, 0x40, 0x3c, 0x40, 0x3f}, // "W"
		0x0058: {0x00, 0x63, 0x14, 0x08, 0x14, 0x63}, // "X"
		0x0059: {0x00, 0x07, 0x08, 0x70, 0x08, 0x07}, // "Y"
		0x005a: {0x00, 0x71, 0x49, 0x45, 0x43, 0x00}, // "Z"
		0x005b: {0x00, 0x00, 0x7f, 0x41, 0x41, 0x00}, // "["
		0x005c: {0x00, 0x02, 0x04, 0x08, 0x10, 0x20}, // "\"
		0x005d: {0x00, 0x00, 0x41, 0x41, 0x7f, 0x00}, // "]"
		0x005e: {0x00, 0x04, 0x02, 0x01, 0x02, 0x04}, // "^"
		0x005f: {0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, // "_"
		0x0060: {0x00, 0x00, 0x03, 0x07, 0x00, 0x00}, // "`"
		0x0061: {0x00, 0x20, 0x54, 0x54, 0x54, 0x78}, // "a"
		0x0062: {0x00, 0x7f, 0x44, 0x44, 0x44, 0x38}, // "b"
		0x0063: {0x00, 0x38, 0x44, 0x44, 0x44, 0x28}, // "c"
		0x0064: {0x00, 0x38, 0x44, 0x44, 0x44, 0x7f}, // "d"
		0x0065: {0x00, 0x38, 0x54, 0x54, 0x54, 0x08}, // "e"
		0x0066: {0x00, 0x08, 0x7e, 0x09, 0x09, 0x00}, // "f"
		0x0067: {0x00, 0x18, 0xa4, 0xa4, 0xa4, 0x7c}, // "g"
		0x0068: {0x00, 0x7f, 0x04, 0x04, 0x78, 0x00}, // "h"
		0x0069: {0x00, 0x00, 0x00, 0x7d, 0x40, 0x00}, // "i"
		0x006a: {0x00, 0x40, 0x80, 0x84, 0x7d, 0x00}, // "j"
		0x006b: {0x00, 0x7f, 0x10, 0x28, 0x44, 0x00}, // "k"
		0x006c: {0x00, 0x00, 0x00, 0x7f, 0x40, 0x00}, // "l"
		0x006d: {0x00, 0x7c, 0x04, 0x18, 0x04, 0x78}, // "m"
		0x006e: {0x00, 0x7c, 0x04, 0x04, 0x78, 0x00}, // "n"
		0x006f: {0x00, 0x38, 0x44, 0x44, 0x44, 0x38}, // "o"
		0x0070: {0x00, 0xfc, 0x44, 0x44, 0x44, 0x38}, // "p"
		0x0071: {0x00, 0x38, 0x44, 0x44, 0x44, 0xfc}, // "q"
		0x0072: {0x00, 0x44, 0x78, 0x44, 0x04, 0x08}, // "r"
		0x0073: {0x00, 0x08, 0x54, 0x54, 0x54, 0x20}, // "s"
		0x0074: {0x00, 0x04, 0x3e, 0x44, 0x24, 0x00}, // "t"
		0x0075: {0x00, 0x3c, 0x40, 0x20, 0x7c, 0x00}, // "u"
		0x0076: {0x00, 0x1c, 0x20, 0x40, 0x20, 0x1c}, // "v"
		0x0077: {0x00, 0x3c, 0x60, 0x30, 0x60, 0x3c}, // "w"
		0x0078: {0x00, 0x6c, 0x10, 0x10, 0x6c, 0x00}, // "x"
		0x0079: {0x00, 0x9c, 0xa0, 0x60, 0x3c, 0x00}, // "y"
		0x007a: {0x00, 0x64, 0x54, 0x54, 0x4c, 0x00}, // "z"
		0x007b: {0x00, 0x08, 0x3e, 0x41, 0x41, 0x00}, // "{"
		0x007c: {0x00, 0x00, 0x00, 0x77, 0x00, 0x00}, // "|"
		0x007d: {0x00, 0x00, 0x41, 0x41, 0x3e, 0x08}, // "}"
		0x007e: {0x00, 0x02, 0x01, 0x02, 0x01, 0x00}, // "~"
		0x00a0: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // " "
		0x00a1: {0x00, 0x00, 0x30, 0x7d, 0x30, 0x00}, // "¡"
		0x00a2: {0x00, 0x18, 0x24, 0x66, 0x24, 0x00}, // "¢"
		0x00a3: {0x00, 0x48, 0x3e, 0x49, 0x49, 0x62}, // "£"
		0x00a4: {0x00, 0x5d, 0x22, 0x22, 0x22, 0x5d}, // "¤"
		0x00a5: {0x00, 0x29, 0x2a, 0x7c, 0x2a, 0x29}, // "¥"
		0x00a6: {0x00, 0x00, 0x00, 0x77, 0x00, 0x00}, // "¦"
		0x00a7: {0x00, 0x22, 0x4d, 0x55, 0x59, 0x22}, // "§"
		0x00a8: {0x00, 0x00, 0x08, 0x00, 0x08, 0x00}, // "¨"
		0x00a9: {0x3e, 0x41, 0x5d, 0x55, 0x41, 0x3e}, // "©"
		0x00aa: {0x00, 0x08, 0x55, 0x55, 0x55, 0x5e}, // "ª"
		0x00ab: {0x00, 0x08, 0x14, 0x00, 0x08, 0x14}, // "«"
		0x00ac: {0x04, 0x04, 0x04, 0x04, 0x04, 0x1c}, // "¬"
		0x00ad: {0x00, 0x00, 0x08, 0x08, 0x08, 0x00}, // "­"
		0x00ae: {0x3e, 0x41, 0x5d, 0x4b, 0x55, 0x3e}, // "®"
		0x00af: {0x00, 0x00, 0x02, 0x02, 0x02, 0x00}, // "¯"
		0x00b0: {0x00, 0x06, 0x09, 0x09, 0x06, 0x00}, // "°"
		0x00b1: {0x00, 0x00, 0x24, 0x2e, 0x24, 0x00}, // "±"
		0x00b2: {0x00, 0x09, 0x0d, 0x0a, 0x00, 0x00}, // "²"
		0x00b3: {0x00, 0x09, 0x0f, 0x05, 0x00, 0x00}, // "³"
		0x00b4: {0x00, 0x00, 0x07, 0x03, 0x00, 0x00}, // "´"
		0x00b5: {0x00, 0xfc, 0x20, 0x20, 0x1c, 0x00}, // "µ"
		0x00b6: {0x00, 0x06, 0x09, 0x7f, 0x01, 0x7f}, // "¶"
		0x00b7: {0x00, 0x00, 0x08, 0x00, 0x00, 0x00}, // "·"
		0x00b8: {0x00, 0x00, 0x08, 0x18, 0x18, 0x00}, // "¸"
		0x00b9: {0x00, 0x02, 0x0f, 0x00, 0x00, 0x00}, // "¹"
		0x00ba: {0x00, 0x4e, 0x51, 0x51, 0x4e, 0x00}, // "º"
		0x00bb: {0x00, 0x14, 0x08, 0x00, 0x14, 0x08}, // "»"
		0x00bc: {0x00, 0x17, 0x08, 0x34, 0x2a, 0x78}, // "¼"
		0x00bd: {0x00, 0x17, 0x08, 0x4c, 0x6a, 0x50}, // "½"
		0x00be: {0x05, 0x17, 0x0a, 0x34, 0x2a, 0x78}, // "¾"
		0x00bf: {0x00, 0x30, 0x48, 0x4d, 0x40, 0x20}, // "¿"
		0x00c0: {0x00, 0x70, 0x29, 0x25, 0x28, 0x70}, // "À"
		0x00c1: {0x00, 0x70, 0x28, 0x25, 0x29, 0x70}, // "Á"
		0x00c2: {0x00, 0x70, 0x29, 0x25, 0x29, 0x70}, // "Â"
		0x00c3: {0x00, 0x70, 0x2a, 0x25, 0x2a, 0x71}, // "Ã"
		0x00c4: {0x00, 0x70, 0x29, 0x24, 0x29, 0x70}, // "Ä"
		0x00c5: {0x00, 0x78, 0x2f, 0x25, 0x2f, 0x78}, // "Å"
		0x00c6: {0x00, 0x7e, 0x09, 0x7f, 0x49, 0x49}, // "Æ"
		0x00c7: {0x00, 0x1e, 0xa1, 0xe1, 0x21, 0x12}, // "Ç"
		0x00c8: {0x00, 0x7c, 0x55, 0x55, 0x54, 0x44}, // "È"
		0x00c9: {0x00, 0x7c, 0x54, 0x54, 0x55, 0x45}, // "É"
		0x00ca: {0x00, 0x7c, 0x55, 0x55, 0x55, 0x44}, // "Ê"
		0x00cb: {0x00, 0x7c, 0x55, 0x54, 0x55, 0x44}, // "Ë"
		0x00cc: {0x00, 0x00, 0x45, 0x7d, 0x44, 0x00}, // "Ì"
		0x00cd: {0x00, 0x00, 0x44, 0x7d, 0x45, 0x00}, // "Í"
		0x00ce: {0x00, 0x00, 0x45, 0x7d, 0x45, 0x00}, // "Î"
		0x00cf: {0x00, 0x00, 0x45, 0x7c, 0x45, 0x00}, // "Ï"
		0x00d0: {0x00, 0x08, 0x7f, 0x49, 0x41, 0x3e}, // "Ð"
		0x00d1: {0x00, 0x7a, 0x11, 0x22, 0x79, 0x00}, // "Ñ"
		0x00d2: {0x00, 0x3d, 0x43, 0x42, 0x3c, 0x00}, // "Ò"
		0x00d3: {0x00, 0x3c, 0x42, 0x43, 0x3d, 0x00}, // "Ó"
		0x00d4: {0x00, 0x3c, 0x43, 0x43, 0x3d, 0x00}, // "Ô"
		0x00d5: {0x00, 0x3a, 0x45, 0x46, 0x39, 0x00}, // "Õ"
		0x00d6: {0x00, 0x3d, 0x42, 0x42, 0x3d, 0x00}, // "Ö"
		0x00d7: {0x00, 0x22, 0x14, 0x08, 0x14, 0x22}, // "×"
		0x00d8: {0x00, 0x7e, 0x61, 0x5d, 0x43, 0x3f}, // "Ø"
		0x00d9: {0x00, 0x3d, 0x41, 0x40, 0x3c, 0x00}, // "Ù"
		0x00da: {0x00, 0x3c, 0x40, 0x41, 0x3d, 0x00}, // "Ú"
		0x00db: {0x00, 0x3c, 0x41, 0x41, 0x3d, 0x00}, // "Û"
		0x00dc: {0x00, 0x3c, 0x41, 0x40, 0x3d, 0x00}, // "Ü"
		0x00dd: {0x00, 0x04, 0x08, 0x71, 0x09, 0x04}, // "Ý"
		0x00de: {0x00, 0xff, 0xa5, 0x24, 0x18, 0x00}, // "Þ"
		0x00df: {0x00, 0xfe, 0x4a, 0x4a, 0x34, 0x00}, // "ß"
		0x00e0: {0x00, 0x20, 0x55, 0x55, 0x54, 0x78}, // "à"
		0x00e1: {0x00, 0x20, 0x54, 0x55, 0x55, 0x78}, // "á"
		0x00e2: {0x00, 0x20, 0x55, 0x55, 0x55, 0x78}, // "â"
		0x00e3: {0x00, 0x20, 0x56, 0x55, 0x56, 0x79}, // "ã"
		0x00e4: {0x00, 0x20, 0x55, 0x54, 0x55, 0x78}, // "ä"
		0x00e5: {0x00, 0x20, 0x57, 0x55, 0x57, 0x78}, // "å"
		0x00e6: {0x00, 0x34, 0x54, 0x7c, 0x54, 0x58}, // "æ"
		0x00e7: {0x00, 0x1c, 0xa2, 0xe2, 0x22, 0x14}, // "ç"
		0x00e8: {0x00, 0x38, 0x55, 0x55, 0x54, 0x08}, // "è"
		0x00e9: {0x00, 0x38, 0x54, 0x54, 0x55, 0x09}, // "é"
		0x00ea: {0x00, 0x38, 0x55, 0x55, 0x55, 0x08}, // "ê"
		0x00eb: {0x00, 0x38, 0x55, 0x54, 0x55, 0x08}, // "ë"
		0x00ec: {0x00, 0x00, 0x01, 0x7c, 0x40, 0x00}, // "ì"
		0x00ed: {0x00, 0x00, 0x00, 0x7d, 0x41, 0x00}, // "í"
		0x00ee: {0x00, 0x00, 0x01, 0x7d, 0x41, 0x00}, // "î"
		0x00ef: {0x00, 0x00, 0x01, 0x7c, 0x41, 0x00}, // "ï"
		0x00f0: {0x00, 0x22, 0x55, 0x59, 0x30, 0x00}, // "ð"
		0x00f1: {0x00, 0x7a, 0x09, 0x0a, 0x71, 0x00}, // "ñ"
		0x00f2: {0x00, 0x39, 0x45, 0x44, 0x38, 0x00}, // "ò"
		0x00f3: {0x00, 0x38, 0x44, 0x45, 0x39, 0x00}, // "ó"
		0x00f4: {0x00, 0x38, 0x45, 0x45, 0x39, 0x00}, // "ô"
		0x00f5: {0x00, 0x32, 0x49, 0x4a, 0x31, 0x00}, // "õ"
		0x00f6: {0x00, 0x38, 0x45, 0x44, 0x39, 0x00}, // "ö"
		0x00f7: {0x00, 0x08, 0x08, 0x2a, 0x08, 0x08}, // "÷"
		0x00f8: {0x80, 0x70, 0x68, 0x58, 0x38, 0x04}, // "ø"
		0x00f9: {0x00, 0x3d, 0x41, 0x20, 0x7c, 0x00}, // "ù"
		0x00fa: {0x00, 0x3c, 0x40, 0x21, 0x7d, 0x00}, // "ú"
		0x00fb: {0x00, 0x3c, 0x41, 0x21, 0x7d, 0x00}, // "û"
		0x00fc: {0x00, 0x3d, 0x40, 0x20, 0x7d, 0x00}, // "ü"
		0x00fd: {0x00, 0x9c, 0xa0, 0x61, 0x3d, 0x00}, // "ý"
		0x00fe: {0x00, 0xfe, 0xaa, 0x28, 0x10, 0x00}, // "þ"
		0x00ff: {0x00, 0x9c, 0xa1, 0x60, 0x3d, 0x00}, // "ÿ"
		0x0131: {0x00, 0x00, 0x00, 0x07, 0x00, 0x00}, // "ı"
		0x0192: {0x00, 0x40, 0x88, 0x7e, 0x09, 0x02}, // "ƒ"
		0x2017: {0x00, 0x24, 0x24, 0x24, 0x24, 0x24}, // "‗"
		0x2022: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "•"
		0x203c: {0x00, 0x00, 0x5f, 0x00, 0x5f, 0x00}, // "‼"
		0x2190: {0x00, 0x08, 0x1c, 0x3e, 0x08, 0x08}, // "←"
		0x2191: {0x00, 0x04, 0x06, 0x7f, 0x06, 0x04}, // "↑"
		0x2192: {0x00, 0x08, 0x08, 0x3e, 0x1c, 0x08}, // "→"
		0x2193: {0x00, 0x10, 0x30, 0x7f, 0x30, 0x10}, // "↓"
		0x2194: {0x00, 0x08, 0x3e, 0x08, 0x3e, 0x08}, // "↔"
		0x2195: {0x00, 0x14, 0x36, 0x7f, 0x36, 0x14}, // "↕"
		0x21a8: {0x00, 0x14, 0xb6, 0xff, 0xb6, 0x14}, // "↨"
		0x221f: {0x00, 0x78, 0x40, 0x40, 0x40, 0x40}, // "∟"
		0x2302: {0x00, 0x3c, 0x26, 0x23, 0x26, 0x3c}, // "⌂"
		0x2500: {0x08, 0x08, 0x08, 0x08, 0x08, 0x08}, // "─"
		0x2502: {0x00, 0x00, 0x00, 0xff, 0x00, 0x00}, // "│"
		0x250c: {0x00, 0x00, 0x00, 0xf8, 0x08, 0x08}, // "┌"
		0x2510: {0x08, 0x08, 0x08, 0xf8, 0x00, 0x00}, // "┐"
		0x2514: {0x00, 0x00, 0x00, 0x0f, 0x08, 0x08}, // "└"
		0x2518: {0x08, 0x08, 0x08, 0x0f, 0x00, 0x00}, // "┘"
		0x251c: {0x00, 0x00, 0x00, 0xff, 0x08, 0x08}, // "├"
		0x2524: {0x08, 0x08, 0x08, 0xff, 0x00, 0x00}, // "┤"
		0x252c: {0x08, 0x08, 0x08, 0xf8, 0x08, 0x08}, // "┬"
		0x2534: {0x08, 0x08, 0x08, 0x0f, 0x08, 0x08}, // "┴"
		0x253c: {0x08, 0x08, 0x08, 0xff, 0x08, 0x08}, // "┼"
		0x2550: {0x0a, 0x0a, 0x0a, 0x0a, 0x0a, 0x0a}, // "═"
		0x2551: {0x00, 0xff, 0x00, 0xff, 0x00, 0x00}, // "║"
		0x2554: {0x00, 0xfe, 0x02, 0xfa, 0x0a, 0x0a}, // "╔"
		0x2557: {0x0a, 0xfa, 0x02, 0xfe, 0x00, 0x00}, // "╗"
		0x255a: {0x00, 0x0f, 0x08, 0x0b, 0x0a, 0x0a}, // "╚"
		0x255d: {0x0a, 0x0b, 0x08, 0x0f, 0x00, 0x00}, // "╝"
		0x2560: {0x00, 0xff, 0x00, 0xfb, 0x0a, 0x0a}, // "╠"
		0x2563: {0x0a, 0xfb, 0x00, 0xff, 0x00, 0x00}, // "╣"
		0x2566: {0x0a, 0xfa, 0x02, 0xfa, 0x0a, 0x0a}, // "╦"
		0x2569: {0x0a, 0x0b, 0x08, 0x0b, 0x0a, 0x0a}, // "╩"
		0x256c: {0x0a, 0xfb, 0x00, 0xfb, 0x0a, 0x0a}, // "╬"
		0x2580: {0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f}, // "▀"
		0x2584: {0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0}, // "▄"
		0x2588: {0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // "█"
		0x2591: {0x44, 0x11, 0x44, 0x11, 0x44, 0x11}, // "░"
		0x2592: {0xaa, 0x55, 0xaa, 0x55, 0xaa, 0x55}, // "▒"
		0x2593: {0xbb, 0xee, 0xbb, 0xee, 0xbb, 0xee}, // "▓"
		0x25a0: {0x00, 0x3c, 0x3c, 0x3c, 0x3c, 0x00}, // "■"
		0x25ac: {0x00, 0x60, 0x60, 0x60, 0x60, 0x00}, // "▬"
		0x25b2: {0x00, 0x30, 0x3c, 0x3f, 0x3c, 0x30}, // "▲"
		0x25ba: {0x00, 0x00, 0x7f, 0x3e, 0x1c, 0x08}, // "►"
		0x25bc: {0x00, 0x03, 0x0f, 0x3f, 0x0f, 0x03}, // "▼"
		0x25c4: {0x00, 0x08, 0x1c, 0x3e, 0x7f, 0x00}, // "◄"
		0x25cb: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "○"
		0x25d8: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "◘"
		0x25d9: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "◙"
		0x263a: {0x00, 0x3e, 0x45, 0x51, 0x45, 0x3e}, // "☺"
		0x263b: {0x00, 0x3e, 0x6b, 0x6f, 0x6b, 0x3e}, // "☻"
		0x263c: {0x00, 0x2a, 0x1c, 0x36, 0x1c, 0x2a}, // "☼"
		0x2640: {0x00, 0x06, 0x29, 0x79, 0x29, 0x06}, // "♀"
		0x2642: {0x00, 0x30, 0x48, 0x4a, 0x36, 0x0e}, // "♂"
		0x2660: {0x00, 0x18, 0x5c, 0x7e, 0x5c, 0x18}, // "♠"
		0x2663: {0x00, 0x30, 0x36, 0x7f, 0x36, 0x30}, // "♣"
		0x2665: {0x00, 0x1c, 0x3e, 0x7c, 0x3e, 0x1c}, // "♥"
		0x2666: {0x00, 0x18, 0x3c, 0x7e, 0x3c, 0x18}, // "♦"
		0x266a: {0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // "♪"
		0x266b: {0x00, 0x60, 0x7e, 0x0a, 0x35, 0x3f}, // "♫"
	},
}