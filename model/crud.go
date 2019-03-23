package model

import (
	"fmt"
	"log"
)

type Logs struct {
	Log_id    string `json:"log_id"`
	Timestamp string `json:"timestamp"`
	Log       string `json:"log"`
	Host      string `json:"host"`
}

func (l *Logs) Create() {
	log.Println(l)
	res, err := con.WriteOne(fmt.Sprintf("INSERT INTO logs(log_id, timestamp, log, host) VALUES ('%s','%s','%s','%s')", l.Log_id, l.Timestamp, l.Log, l.Host))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Created")
	log.Println(res)

}

func (l Logs) Read(c chan []Logs) {

	var (
		logid string
		time  string
		logg  string
		host  string
		logs  []Logs
	)

	res, err := con.QueryOne("SELECT * FROM logs")
	if err != nil {
		log.Println(err)
		c <- logs
		return
	}

	for res.Next() {
		err = res.Scan(&logid, &time, &logg, &host)
		if err != nil {
			log.Println(err)
			c <- logs
			return
		}
		logs = append(logs, Logs{logid, time, logg, host})
	}
	c <- logs
	return
}
