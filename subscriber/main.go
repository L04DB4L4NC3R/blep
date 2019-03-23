package main

import (
	"os"

	subscriber "github.com/angadsharma1016/c2c/subscriber/sub"
)

func main() {
	subscriber.Subscribe(os.Args[1])
	select {}
}
