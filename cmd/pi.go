
package main 

import (
	"github.com/stianeikeland/go-rpio/v4"

	"fmt"
	"time"
	"sync"
)

func openDoor () {
	fmt.Println ("opening door")

	pin := rpio.Pin(18)
	pin.Output()
	pin.High() // make it high
	time.Sleep(time.Second * 10)

	pin.Low()
	pin.PullOff()

	fmt.Println ("close door")
}

func monitorButton (wg *sync.WaitGroup, running *bool) {
	pin := rpio.Pin(24)
	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge) // look for it falling to ground

	triggered := false 
	for *running {
		if triggered == false && pin.EdgeDetected() {
			openDoor() // open saysme
			triggered = true // i think we're getting re-entry stuff...
		} else {
			triggered = false 
		}
		time.Sleep(time.Second / 2)
	}
}
