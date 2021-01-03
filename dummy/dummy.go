package main

import (
	"fmt"
	"github.com/racerxdl/gonx/svc"
)

func main() {
	fmt.Println("Hello world to debug console!")

	svc.Break(0, 0, 0)

	for {
	}
}
