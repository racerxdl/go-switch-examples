package main

import (
	"github.com/racerxdl/gonx/nx"
	"github.com/racerxdl/gonx/nx/graphics"
	"github.com/racerxdl/gonx/nx/system"
)

func main() {
	system.SvcOutputDebugString("Getting window")
	window := graphics.GetDefaultWindow()

	system.SvcOutputDebugString("Making FB")
	fb, err := window.MakeFramebuffer(1280, 720, 2, graphics.PixelFormatRgba8888)

	if err != nil {
		panic(err)
	}

	system.SvcOutputDebugString("Making FB Linear")
	fb.MakeLinear()
	i := 0

	for nx.AppletMainLoop() {
		nx.HidScanInput()

		keysDown := nx.HidKeysDown(nx.ControllerP1Auto)

		if keysDown.IsKeyDown(nx.KeyPlus) {
			break
		}

		img, err := fb.StartFrameAsRGBA()
		if err != nil {
			panic(err)
		}

		i++
		if i > 60 {
			i = 0
		}

		shade := uint8((i * 256) / 60)

		shade32 := uint32(shade)
		shade32 = 0xFF<<24 + shade32<<16 + shade32<<8 + shade32

		for y := 0; y < 720; y++ {
			for x := 0; x < 1280; x++ {
				img.SetRGBA32(x, y, shade32)
			}
		}

		img.End()
	}

	system.SvcOutputDebugString("Framebuffer Close")
	fb.Close()
}
