package main

import (
	"log"

	"github.com/dvl-numeez/nats"
)

func main() {
	js, err := nats.JetStreamInit()
	if err != nil {
		log.Fatal("JS Stream Init error : ", err)
	}
	publishMessage(js)

}
