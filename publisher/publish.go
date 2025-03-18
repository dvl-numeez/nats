package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/dvl-numeez/nats/config"
	"github.com/dvl-numeez/nats/models"
	"github.com/nats-io/nats.go"
)

func publishMessage(js nats.JetStreamContext) {
	messages, err := models.GetMessages()
	if err != nil {
		log.Println("JSON error : ",err)
		return
	}
	for _, message := range messages {
		r := rand.Intn(1500)
		time.Sleep(time.Duration(r) * time.Millisecond)

		reviewString, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
			continue
		}
		_, err = js.Publish(config.SubjectNameReviewCreated, reviewString)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Publisher  =>  Message: %s\n", message.Content)
		}
	}
}
