package main

import (
	"syscall/js"
	"time"
)

func main() {
	loc := time.FixedZone("America/Los_Angeles", -8*60*60)
	day := js.Global().Get("document").Call("getElementById", "day")
	t := js.Global().Get("document").Call("getElementById", "time")
	for {
		now := time.Now().In(loc)
		day.Set("textContent", now.Format("2006 01/02"))
		t.Set("textContent", now.Format("15:04:05"))
		time.Sleep(time.Millisecond * 100)
	}
}
