// This program is a minimal Nintendo Switch program that just outputs messages to
// emulator console
package main

func test2() {
	panic("test2")
}

func main() {
	println("Hello world!")
	println("ABC", 123, 343)
	//panic("TEST")

	x := 0

	for x < 100 {
		x++
		println("CYCLE")
		if x > 50 {
			test2()
		}
	}

	//
	//
	//for {
	//   println("Cycle!!")
	//   time.Sleep(time.Second)
	//}
	panic("TAS")
}
