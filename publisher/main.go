package main

import (
	"log"
	"os"

	"github.com/angadsharma1016/c2c/model"
	"github.com/angadsharma1016/c2c/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

// publishOrderCreated publish an event via NATS server
func PublishLogs(logs *pb.LogStore) {
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
	con := model.Connect(os.Args[5])
	defer con.Close()
	logg := pb.LogStore{
		LogId:     os.Args[1],
		Timestamp: os.Args[2],
		Log:       os.Args[3],
		Host:      os.Args[4],
	}
	go PublishLogs(&logg)
	l := model.Logs{os.Args[1], os.Args[2], os.Args[3], os.Args[4]}
	go l.Create()
}
