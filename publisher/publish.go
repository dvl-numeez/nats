package main

import (
	"encoding/json"
	"log"

	"github.com/dvl-numeez/nats/models"
	"github.com/nats-io/nats.go"
)

func publishMessage(js nats.JetStreamContext, subjectName string, data models.MessageInfo) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	_, err = js.Publish(subjectName, dataBytes)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Publisher  =>  Message: %s\n", data.TrackingId)
	}
}
