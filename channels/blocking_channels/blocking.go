//the following example won't work bcz send/receive in channels block,
//whenever a value is send to a channel,there should be someone to receive
// from the channel at that very instance

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
