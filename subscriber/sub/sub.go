package subscriber

import (
	"log"
	"os"
	"os/exec"

	"github.com/angadsharma1016/c2c/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

func Subscribe(sub string, cb string) {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.Subscribe(sub, func(msg *nats.Msg) {
		logsStore := pb.LogStore{}
		err := proto.Unmarshal(msg.Data, &logsStore)
		if err == nil {
			// Handle the message and run callback
			log.Printf("Received message in LogStore service: %+v\n", logsStore)
			cmd := exec.Command("./bin/" + cb)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalln(err)
			}
		}
	})
}
