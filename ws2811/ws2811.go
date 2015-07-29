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

// Interface to ws2811 chip (neopixel driver). Make sure that you have
// ws2811.h and pwm.h in a GCC include path (e.g. /usr/local/include) and
// libws2811.a in a GCC library path (e.g. /usr/local/lib).
// See https://github.com/jgarff/rpi_ws281x for instructions
package ws2811

/*
#cgo CFLAGS: -std=c99
#cgo LDFLAGS: -lws2811
#include <stdint.h>
#include <string.h>
#include <ws2811.h>

ws2811_t ledstring = {
   .freq = 800000,
   .dmanum = 5,
   .channel = {
	   [0] = {
		   .gpionum = 18,
		   .count = 256,
		   .invert = 0,
		   .brightness = 32,
	   },
	   [1] = {
		   .gpionum = 0,
		   .count = 0,
		   .invert = 0,
		   .brightness = 0,
	   },
   },
};

void ws2811_set_led(ws2811_t *ws2811, int index, uint32_t value) {
	ws2811->channel[0].leds[index] = value;
}

void ws2811_clear(ws2811_t *ws2811) {
	for (int chan = 0; chan < RPI_PWM_CHANNELS; chan++) {
		ws2811_channel_t *channel = &ws2811->channel[chan];
		memset(channel->leds, 0, sizeof(ws2811_led_t) * channel->count);
	}
}

void ws2811_set_bitmap(ws2811_t *ws2811, void* a, int len) {
	memcpy(ws2811->channel[0].leds, a, len);
}

*/
import "C"
import "unsafe"

type Color uint32

func Init(gpioPin int, ledCount int, brightness int) int {
	C.ledstring.channel[0].gpionum = C.int(gpioPin)
	C.ledstring.channel[0].count = C.int(ledCount)
	C.ledstring.channel[0].brightness = C.int(brightness)
	res := int(C.ws2811_init(&C.ledstring))
	return res
}

func Fini() {
	C.ws2811_fini(&C.ledstring)
}

func Render() int {
	res := int(C.ws2811_render(&C.ledstring))
	return res
}

func Wait() int {
	res := int(C.ws2811_wait(&C.ledstring))
	return res
}

func SetLed(index int, value Color) {
	C.ws2811_set_led(&C.ledstring, C.int(index), C.uint32_t(value))
}

func Clear() {
	C.ws2811_clear(&C.ledstring)
}

func SetBitmap(a []Color) {
	C.ws2811_set_bitmap(&C.ledstring, unsafe.Pointer(&a[0]), C.int(len(a)*4))
}

func RGB(r int, g int, b int) Color {
	return Color(r)<<16 + Color(g)<<8 + Color(b)
}
