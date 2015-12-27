package main

import (
	"fmt"
	"github.com/heia-fr/telecom-tower/build-bitmap-font/cp850"
	"github.com/heia-fr/telecom-tower/build-bitmap-font/font"
	"os"
	"sort"
	"strings"
)

var test = map[rune]byte{
	'\u263A': 0x01,
	0x12:     0x02,
}

func main() {

	f, err := os.Create("res8.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var keys []int
	for k, _ := range cp850.Cp850 {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	fmt.Fprintln(f, "var Font8x8 = map[rune][]byte{")
	for _, k := range keys {
		comment := fmt.Sprintf("%c", rune(k))
		key := fmt.Sprintf("0x%04x", rune(k))

		var data []string
		for _, i := range font.F88.GetCharLSB(cp850.Cp850[rune(k)]) {
			data = append(data, fmt.Sprintf("0x%02x", i))
		}

		fmt.Fprintf(f, "\t%v: {%v}, // \"%v\"\n", key, strings.Join(data, ", "), comment)
	}
	fmt.Fprintln(f, "}")

	f, err = os.Create("res6.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "var Font6x8 = map[rune][]byte{")
	for _, k := range keys {
		comment := fmt.Sprintf("%c", rune(k))
		key := fmt.Sprintf("0x%04x", rune(k))

		var data []string
		for _, i := range font.F68.GetCharLSB(cp850.Cp850[rune(k)]) {
			data = append(data, fmt.Sprintf("0x%02x", i))
		}

		fmt.Fprintf(f, "\t%v: {%v}, // \"%v\"\n", key, strings.Join(data, ", "), comment)
	}
	fmt.Fprintln(f, "}")

}
