package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

// Suggestions from golang-nuts
// http://play.golang.org/p/Ctg3_AQisl

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func helloworld(t time.Time) {
	fmt.Printf("%v: Hello, World!\n", t)
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}

// func getLocation() {

// 	x, y := robotgo.GetMousePos()
// 	fmt.Println("pos:", x, y)

// }

func main() {

	doEvery(100*time.Millisecond, helloworld)
	robotgo.ScrollMouse(10, "up")
	robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 1.5)
}
