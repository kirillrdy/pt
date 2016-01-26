package pt

import (
	"github.com/kirillrdy/pt/xlib"
)

//RenderEvent contains rendered pixel with X and Y coordinates
type pixelRenderResult struct {
	X     int
	Y     int
	Pixel Color
}

func RenderToWindow(pixelRenderResult chan pixelRenderResult) {

	xlib.CreateWindow(RenderConfig.Width, RenderConfig.Height)

	for event := range pixelRenderResult {
		xlib.SetPixel(event.X, event.Y, int(event.Pixel.R*65535), int(event.Pixel.G*65535), int(event.Pixel.B*65535))
	}
}
