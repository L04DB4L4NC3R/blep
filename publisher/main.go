package main

import (
	"os"

	"github.com/angadsharma1016/c2c/model"
	"github.com/angadsharma1016/c2c/pb"
	publisher "github.com/angadsharma1016/c2c/publisher/pub"
)

func main() {
	con := model.Connect()
	defer con.Close()
	logg := pb.LogStore{
		LogId:     os.Args[1],
		Timestamp: os.Args[2],
		Log:       os.Args[3],
		Host:      os.Args[4],
	}
	publisher.PublishLogs(&logg)
	l := model.Logs{os.Args[1], os.Args[2], os.Args[3], os.Args[4]}
	l.Create()
}
