
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
		fmt.Println ("pin ", p)
		pin := rpio.Pin(p)

		pin.Output()

		for x := 0; x < 5; x++ {
			fmt.Println ("making high")
			pin.High() // make it high
			time.Sleep(time.Second * 4)
			fmt.Println ("making low")
			pin.Low()
			time.Sleep(time.Second * 2)
		}

		fmt.Println ("done with pin")
		pin.PullOff()
	}
	
	fmt.Println ("exiting")
}