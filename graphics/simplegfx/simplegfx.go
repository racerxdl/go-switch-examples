// Hello World example
// Same as https://github.com/switchbrew/switch-examples/tree/master/graphics/printing/hello-world
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
	//
	if err != nil {
		panic(err)
	}
	//

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

		for x := 0; x < 1280; x++ {
			for y := 0; y < 720; y++ {
				off := y*img.Stride + x*4
				img.Pix[off] = shade
				img.Pix[off+1] = shade
				img.Pix[off+2] = shade
				img.Pix[off+3] = 255
			}
		}

		img.End()
	}

	system.SvcOutputDebugString("Framebuffer Close")
	fb.Close()
}
