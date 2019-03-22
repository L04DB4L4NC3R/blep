package main

import (
	"log"
	"runtime"
	"time"

	"github.com/angadsharma1016/c2c/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

const subject = "logs.>"

func main() {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		logsStore := pb.LogStore{}
		err := proto.Unmarshal(msg.Data, &logsStore)
		if err == nil {
			// Handle the message
			log.Printf("Received message in LogStore service: %+v\n", logsStore)
		}
	})

	// Keep the connection alive
	time.Sleep(time.Second * 3)
	runtime.Goexit()
}
