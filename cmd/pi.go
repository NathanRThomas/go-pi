
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
	time.Sleep(time.Second * 30)

	pin.Low()
	pin.PullOff()

	fmt.Println ("close door")
}

func monitorButton (wg *sync.WaitGroup, running *bool) {
	pin := rpio.Pin(24)
	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge) // look for it falling to ground

	for *running {
		if pin.EdgeDetected() {
			openDoor() // open saysme
		}
		time.Sleep(time.Second / 2)
	}
}
