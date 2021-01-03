// This program is a minimal Nintendo Switch program that just outputs messages to
// emulator console
package main

import (
	"fmt"
	"github.com/racerxdl/gonx/font"
	"github.com/racerxdl/gonx/nintendoswitch"
	"github.com/racerxdl/gonx/services/display"
	"github.com/racerxdl/gonx/svc"
	"unsafe"
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
		svc.Break(0x4443, 0, 0)
		return
	}
	//defer surface.CloseLayer()

	fmt.Println("Got surface!")

	fmt.Println("Fetching Frame")
	frame, err := surface.GetFrame()
	if err != nil {
		fmt.Printf("Error getting frame: %s\n", err)
		svc.Break(0x4444, 0, 0)
		*(*uint64)(unsafe.Pointer(uintptr(1))) = uint64(0xDEAD155500000002)
		return
	}

	frame.Clear(Red)

	fmt.Println("Drawing Red Square")
	for y := int16(0); y < 64; y++ {
		for x := int16(0); x < 64; x++ {
			frame.SetPixel(128+x, 128+y, Red)
		}
	}

	fmt.Println("Displaying image")
	err = frame.Display()
	if err != nil {
		fmt.Printf("Error displaying data: %s\n", err)
		svc.Break(0x4445, 0, 0)
		return
	}

	drawfont := font.GetFontByName("nixedsys_normal")

	fmt.Println("DRAWN!")
	err = frame.WaitVSync()
	if err != nil {
		fmt.Println("Error waiting vsync", err)
		svc.Break(0x4446, 0, 0)
		return
	}
	i := int16(0)
	totalFrameCount := 0
	for {
		frame.Clear(Black)
		for y := int16(0); y < 64; y++ {
			for x := int16(0); x < 64; x++ {
				frame.SetPixel(i+x, i+y, Red)
			}
		}
		i++
		if i > 256 {
			i = 0
		}

		frame.DrawStringAt(800, 100, "Hello world!!ãõç1;/|/\nHUEBR2\n1234", White, drawfont)

		err = frame.Display()
		if err != nil {
			fmt.Printf("Error displaying data: %s\n", err)
			*(*uint64)(unsafe.Pointer(uintptr(1))) = uint64(0xDEAD155500000004)
			return
		}
		err = frame.WaitVSync()
		if err != nil {
			fmt.Println("Error waiting vsync", err)
			*(*uint64)(unsafe.Pointer(uintptr(1))) = uint64(0xDEAD155500000005)
			return
		}
		totalFrameCount++
		if totalFrameCount > 100 {
			break
		}
	}
}
