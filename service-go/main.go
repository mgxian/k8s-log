package main

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

type log struct {
	Service   string    `json:"service,omitempty"`
	Timestamp time.Time `json:"@timestamp,omitempty"`
	Message   string    `json:"message"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch-logging.logging:9200"))
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	ctx := context.Background()
	count := 1
	for {
		message := fmt.Sprintf("log from service-go %d", count)
		log := log{Service: "service-go", Message: message, Timestamp: time.Now()}
		index := fmt.Sprintf("k8s-app-%.4d.%.2d.%.2d", log.Timestamp.Year(), log.Timestamp.Month(), log.Timestamp.Day())
		resp, err := client.Index().
			Index(index).
			Type("log").
			BodyJson(log).
			Do(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Indexed log %s to index %s, type %s\n", resp.Id, resp.Index, resp.Type)
		time.Sleep(time.Second)
		count++
	}
}
