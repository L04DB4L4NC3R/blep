package subscriber

import (
	"log"

	"github.com/angadsharma1016/c2c/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

type callback func(interface{}) (string, error)

func Subscribe(sub string, cb callback) {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.Subscribe(sub, func(msg *nats.Msg) {
		logsStore := pb.LogStore{}
		err := proto.Unmarshal(msg.Data, &logsStore)
		if err == nil {
			// Handle the message
			log.Printf("Received message in LogStore service: %+v\n", logsStore)
		}
		message, err := cb(logsStore)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(message)
	})

	// Keep the connection alive
	select {}
	//runtime.Goexit()
}
