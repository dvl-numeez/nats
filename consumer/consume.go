package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dvl-numeez/nats/config"
	"github.com/dvl-numeez/nats/models"
	"github.com/nats-io/nats.go"
)

func consumeMessages(js nats.JetStreamContext, subjectName string) {
	_, err := js.Subscribe(config.StreamName+"."+subjectName, func(m *nats.Msg) {
		var message models.MessageInfo
		err := json.Unmarshal(m.Data, &message)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message)
		err = m.Ack()
		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}
	},
		nats.ManualAck(),
		nats.DeliverNew(),
		nats.Durable(subjectName))

	if err != nil {
		log.Println("Subscribe failed")
		return
	}
}
