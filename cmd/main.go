
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

	pin := rpio.Pin(10)

	pin.Output()

	for x := 0; x < 300; x++ {
		fmt.Println ("making high")
		pin.High() // make it high
		time.Sleep(time.Second * 2)
		fmt.Println ("making low")
		pin.Low()
		time.Sleep(time.Second * 2)
	}

	fmt.Println ("done with this crazyness")
	pin.PullOff()
	
	fmt.Println ("exiting")
}