
package main 

import (
	"github.com/stianeikeland/go-rpio/v4"

	"fmt"
	"log"
	"time"
)

func main () {
	fmt.Println ("starting things")

	err := rpio.Open()
	if err != nil {
		log.Fatal (err)
	}
	defer rpio.Close()

	for p := 0; p < 20; p++ {
		fmt.Println ("pin ", p, "high")
		pin := rpio.Pin(p)

		pin.Output()

		pin.High() // make it high
		time.Sleep(time.Second * 4)
	}

	for p := 0; p < 20; p++ {
		fmt.Println ("pin ", p, "low")
		pin := rpio.Pin(p)

		pin.Low()
		time.Sleep(time.Second * 4)
		pin.PullOff()
	}

		
	fmt.Println ("exiting")
}