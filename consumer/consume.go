package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dvl-numeez/nats/config"
	"github.com/dvl-numeez/nats/models"
	"github.com/nats-io/nats.go"
)

func consumeMessages(js nats.JetStreamContext) {
	_, err := js.Subscribe(config.SubjectNameReviewCreated, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		var message models.Message
		err = json.Unmarshal(m.Data, &message)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message)
	})

	if err != nil {
		log.Println("Subscribe failed")
		return
	}
}
