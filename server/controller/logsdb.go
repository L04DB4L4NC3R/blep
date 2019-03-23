package controller

import (
	"encoding/json"
	"net/http"

	"github.com/angadsharma1016/c2c/model"
	"github.com/angadsharma1016/c2c/pb"
	publisher "github.com/angadsharma1016/c2c/publisher/pub"
	subscriber "github.com/angadsharma1016/c2c/subscriber/sub"
)

type Publisher struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Log       string `json:"log"`
	Host      string `json::"host"`
}

type Subscriber struct {
	Service  string `json:"service"`
	Callback string `json:"callback"`
}

func readLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := make(chan []model.Logs)
		go model.Read(c)
		msg := <-c
		json.NewEncoder(w).Encode(msg)
	}
}

func publish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pub Publisher
		json.NewDecoder(r.Body).Decode(&pub)
		logg := pb.LogStore{
			LogId:     pub.ID,
			Timestamp: pub.Timestamp,
			Log:       pub.Log,
			Host:      pub.Host,
		}
		publisher.PublishLogs(&logg)
		l := model.Logs{pub.ID, pub.Timestamp, pub.Log, pub.Host}
		l.Create()
	}
}

func subscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sub Subscriber
		json.NewDecoder(r.Body).Decode(&sub)

		subscriber.Subscribe(sub.Service, sub.Callback)
	}
}
