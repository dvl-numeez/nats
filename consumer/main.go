package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	var trackingIds []string
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> Enter trackingId to be subscribed: ")
		reader.Scan()
		data := reader.Text()
		if len(data) == 0 {
			continue
		}
		ids, err := parseData(data)
		if err != nil {
			fmt.Println("> error: ", err)
			continue
		}
		if len(ids) != 0 {
			trackingIds = ids
			break
		}

	}
	for _, id := range trackingIds {
		consumeMessages(js, id)
	}

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Println("\nShutting consumer")

}

func parseData(data string) ([]string, error) {
	var trackingIds []string
	err := json.Unmarshal([]byte(data), &trackingIds)
	if err != nil {
		return trackingIds, err
	}
	return trackingIds, nil

}
