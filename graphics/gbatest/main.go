// Based on https://github.com/conejoninja/gbatest
// but ported to nintendoswitch
package main

// Draw a red square on the GameBoy Advance screen.

import (
	"fmt"
	"github.com/racerxdl/gonx/nintendoswitch"
	"github.com/racerxdl/gonx/services/display"
	"github.com/racerxdl/gonx/svc"
	"image/color"
	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	err := nintendoswitch.Configure()
	if err != nil {
		println("error initializing nintendoswitch: ", err)
		svc.Break(0x4442, 0, 0)
		return
	}
	defer nintendoswitch.Cleanup()

	fmt.Println("Fetching surface")
	surface, err := display.OpenLayer()
	if err != nil {
		fmt.Printf("Error getting surface: %s\n", err)
		return
	}

	fmt.Println("Fetching Frame")
	frame, err := surface.GetFrame()
	if err != nil {
		fmt.Printf("Error getting frame: %s\n", err)
		return
	}

	frames := 0

	for {
		if frames%8 == 0 {
			update(frame)
		}

		frames++

		err = frame.Display()
		if err != nil {
			fmt.Printf("Error displaying data: %s\n", err)
			return
		}

		err = frame.WaitVSync()
		if err != nil {
			fmt.Println("Error waiting vsync", err)
			return
		}
		//time.Sleep(time.Millisecond * 10)
	}
}

var (
	krgb = uint8(0)
)

func drawLogo(display drivers.Displayer, x, y int16) {
	tinyfont.DrawChar(display, &freesans.Bold24pt7b, x+36, y+60, 'T', getRainbowRGB(krgb))
	tinyfont.DrawChar(display, &freesans.Bold24pt7b, x+66, y+60, 'i', getRainbowRGB(1+krgb))
	tinyfont.DrawChar(display, &freesans.Bold24pt7b, x+79, y+60, 'n', getRainbowRGB(2+krgb))
	tinyfont.DrawChar(display, &freesans.Bold24pt7b, x+108, y+60, 'y', getRainbowRGB(3+krgb))
	tinyfont.DrawChar(display, &freesans.Bold24pt7b, x+134, y+60, 'G', getRainbowRGB(4+krgb))
	tinyfont.DrawChar(display, &freesans.Bold24pt7b, x+170, y+60, 'o', getRainbowRGB(5+krgb))
}

func drawFosdem(display drivers.Displayer, x, y int16) {
	fosdemXOffset := int16(100)
	tinyfont.WriteLine(display, &freemono.Bold9pt7b, x+70, y+90, "Go compiler for small places", rescale(color.RGBA{R: 10, G: 30, B: 30, A: 255}))
	tinydraw.Triangle(display, x+60+fosdemXOffset, y+110, x+60+fosdemXOffset, y+126, x+174+fosdemXOffset, y+110, rescale(color.RGBA{R: 20, G: 20, A: 255}))
	tinydraw.Triangle(display, x+72+fosdemXOffset, y+136, x+184+fosdemXOffset, y+136, x+184+fosdemXOffset, y+120, rescale(color.RGBA{R: 20, B: 20, A: 255}))
	tinyfont.WriteLine(display, &freesans.Bold9pt7b, x+68+fosdemXOffset, y+130, "FOSDEM '20", rescale(color.RGBA{R: 10, G: 30, B: 30, A: 255}))
}

func drawGophers(display drivers.Displayer, x, y int16) {
	if krgb >= 40 {
		krgb = 0
	}
	if krgb%2 == 0 {
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+14, y+140, 'N', rescale(color.RGBA{A: 255}))
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+230, y+140, 'G', rescale(color.RGBA{A: 255}))
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+230, y+140, 'N', getRainbowRGB(krgb))
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+14, y+140, 'G', getRainbowRGB(krgb))
	} else {
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+230, y+140, 'N', rescale(color.RGBA{A: 255}))
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+14, y+140, 'G', rescale(color.RGBA{A: 255}))
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+14, y+140, 'N', getRainbowRGB(krgb))
		tinyfont.DrawChar(display, &gophers.Regular121pt, x+230, y+140, 'G', getRainbowRGB(krgb))
	}
}

func update(display drivers.Displayer) {
	drawLogo(display, 514, 200)
	drawFosdem(display, 412, 200)
	drawGophers(display, 480, 250)
	krgb++
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 10 {
		return rescale(color.RGBA{R: i * 3, G: 30 - i*3, A: 255})
	} else if i < 20 {
		i -= 10
		return rescale(color.RGBA{R: 30 - i*3, B: i * 3, A: 255})
	}
	i -= 20
	return rescale(color.RGBA{G: i * 3, B: 30 - i*3, A: 255})
}

func rescale(c color.RGBA) color.RGBA {
	return color.RGBA{R: c.R * 4, G: c.G * 4, B: c.B * 4, A: c.A}
}
