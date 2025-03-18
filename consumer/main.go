package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dvl-numeez/nats"
)

func main() {

	js, err := nats.JetStreamInit()
	if err != nil {
		log.Fatal(err)
	}
	consumeMessages(js)
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Println("\nShutting consumer")

}
