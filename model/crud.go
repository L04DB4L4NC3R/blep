package model

import (
	"fmt"
	"log"
)

type Logs struct {
	Log_id    string
	Timestamp string
	Log       string
	Host      string
}

func (l *Logs) Create() {
	log.Println(l)
	res, err := con.WriteOne(fmt.Sprintf("INSERT INTO logs(log_id, timestamp, log, host) VALUES (%s,%s,%s,%s)", l.Log_id, l.Timestamp, l.Log, l.Host))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Created")
	log.Println(res)

}

func (l Logs) Read(c chan []Logs) {
	var logs []Logs

	res, err := con.QueryOne("SELECT * FROM logs")
	if err != nil {
		log.Println(err)
		c <- logs
		return
	}

	for res.Next() {
		var result Logs
		err = res.Scan(&result)
		log.Println(result)
		if err != nil {
			log.Println(err)
			c <- logs
			return
		}
		logs = append(logs, result)
	}
	c <- logs
	return
}
