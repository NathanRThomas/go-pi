
package main 

import (
    /*
	"github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
	*/

    "github.com/stianeikeland/go-rpio/v4"

	// "context"
	// "fmt"
	"log"
    "sync"
    "os"
    "os/signal"
	"syscall"
)

func main () {

    // talk to the pi
    err := rpio.Open()
	if err != nil {
		log.Fatal (err)
	}
	defer rpio.Close() // close this when main closes

    // monitor things
    running := true 
    c := make(chan os.Signal, 1)
	signal.Notify (c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-c // this sits until something comes into the channel, eg the notify interupts from above
		running = false
	}()

    wg := new(sync.WaitGroup)
    wg.Add(1)

    go monitorButton(wg, &running)

    wg.Wait() // wait for an exit trigger

    /*
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws")
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0x6A540EB4a7BA92BA0Cf9f6117042C3658CEaD298")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            fmt.Println(vLog) // pointer to event log
        }
    }
    */
}