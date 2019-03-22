package main

import (
	"log"

	"github.com/angadsharma1016/c2c/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

// publishOrderCreated publish an event via NATS server
func publishLogs(logs *pb.LogStore) {
	// Connect to NATS server
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()
	subject := "logs." + logs.GetHost()
	data, _ := proto.Marshal(logs)

	// Publish message on subject
	natsConnection.Publish(subject, data)
	log.Println("Published message on subject " + subject)
}

func main() {
	logs := pb.LogStore{
		LogId:     "1",
		Timestamp: "20/01/01",
		Log:       "1",
		Host:      "1",
	}
	publishLogs(&logs)
}
