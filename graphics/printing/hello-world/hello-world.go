// Hello World example
// Same as https://github.com/switchbrew/switch-examples/tree/master/graphics/printing/hello-world
package main

import (
	"fmt"
	"github.com/racerxdl/gonx/nx"
)

func main() {
	nx.ConsoleInit(nil)

	fmt.Println("\x1b[16;20HHello World!")

	for nx.AppletMainLoop() {
		nx.HidScanInput()

		keysDown := nx.HidKeysDown(nx.ControllerP1Auto)

		if keysDown.IsKeyDown(nx.KeyPlus) {
			break
		}

		nx.ConsoleUpdate(nil)
	}

	nx.ConsoleExit(nil)
}
