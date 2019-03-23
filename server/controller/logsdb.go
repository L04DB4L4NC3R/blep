package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/angadsharma1016/c2c/model"
	"github.com/angadsharma1016/c2c/pb"
	publisher "github.com/angadsharma1016/c2c/publisher/pub"
	subscriber "github.com/angadsharma1016/c2c/subscriber/sub"
	"github.com/gogo/protobuf/proto"
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

type Num struct {
	Calls int `json:"calls"`
}

type Benchmarks struct {
	PublishingTime   time.Duration `json:"publishing-time"`
	PublishedCalls   int           `json:"published-calls"`
	ProtoMarshalTime time.Duration `json:"proto-marshal-time"`
	JSONDecodingTime time.Duration `json:"JSON-decoding-time"`
	SubscribingTime  time.Duration `json:"subscribing-time"`
	NATSConnTime     time.Duration `json:"nats-conn-time"`
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

		json.NewEncoder(w).Encode(struct {
			Status bool `json:"status"`
		}{true})
	}
}

// func subscribe() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var sub Subscriber
// 		json.NewDecoder(r.Body).Decode(&sub)

// 		subscriber.Subscribe("logs."+sub.Service, sub.Callback)

// 		json.NewEncoder(w).Encode(struct {
// 			Status bool `json:"status"`
// 		}{true})
// 	}
// }

func benchmark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			wg        sync.WaitGroup
			obj       Num
			benchmark Benchmarks
		)

		// decode JSON
		decodeJSON := time.Now()
		json.NewDecoder(r.Body).Decode(&obj)
		benchmark.JSONDecodingTime = time.Since(decodeJSON)

		benchmark.PublishedCalls = obj.Calls

		// connecting to NATS
		natsTime := time.Now()
		natsconn := publisher.ConnReturn()
		defer natsconn.Close()
		benchmark.NATSConnTime = time.Since(natsTime)

		// subscribe to testnet
		subsTime := time.Now()
		subscriber.Subscribe("logs.testnet", "callback")
		benchmark.SubscribingTime = time.Since(subsTime)

		logs := pb.LogStore{
			LogId:     "test",
			Host:      "test",
			Timestamp: fmt.Sprintf("%v", time.Now()),
			Log:       "test",
		}

		// marshal to proto
		protoTime := time.Now()
		data, _ := proto.Marshal(&logs)
		benchmark.ProtoMarshalTime = time.Since(protoTime)

		// add required number of services to waitgroup
		wg.Add(obj.Calls)

		// benchmarking pub-sub
		pubTime := time.Now()
		for i := 0; i < obj.Calls; i++ {
			go func() {
				defer wg.Done()
				natsconn.Publish("test.testnet", data)
			}()
		}
		wg.Wait()
		benchmark.PublishingTime = time.Since(pubTime)

		json.NewEncoder(w).Encode(benchmark)

	}
}
