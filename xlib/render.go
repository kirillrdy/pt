package xlib

// #cgo LDFLAGS: -lX11
// #include "main.h"
import "C"

func CreateWindow(width, height int) {
	C.create_window(C.int(width), C.int(height))
}

func SetPixel(x, y, r, g, b int) {
	C.set_pixel(C.int(x), C.int(y), C.int(r), C.int(g), C.int(b))
}

func Flush() {
	C.flush()
}
