package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

// LogData log data
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

// Init init
func Init(address string, chanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Printf("failed to connect es, err: %v\n", err)
		return
	}
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go SendToES()
	}
	return
}

// SendToESChan send to es chan
func SendToESChan(msg *LogData) {
	ch <- msg
}

// SendToES send to es chan
func SendToES() {
	for {
		select {
		case msg := <-ch:
			put, err := client.Index().Index(msg.Topic).Type("XXX").BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Printf("failed to send log es, err: %v\n", err)
			}
			fmt.Printf("Indexed student %v to index %s, type %s \n", put.Id, put.Index, put.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
