package model

import (
	"fmt"
	"log"

	"github.com/raindog308/gorqlite"
)

var con gorqlite.Connection

func Connect(create string) *gorqlite.Connection {
	conn, err := gorqlite.Open("http://")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// create logs DB if user specifies
	if create == "y" {
		res, err := conn.QueryOne("CREATE TABLE logs(log_id VARCHAR(10) PRIMARY KEY, timestamp VARCHAR(25), log VARCHAR(200), host VARCHAR(20))")
		if err != nil {
			log.Println(err)
		}
		fmt.Println(res)
	}

	con = conn
	return &con
}
