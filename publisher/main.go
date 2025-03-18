package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dvl-numeez/nats"
	"github.com/dvl-numeez/nats/config"
	"github.com/dvl-numeez/nats/models"
)

func main() {
	js, err := nats.JetStreamInit()
	if err != nil {
		log.Fatal("JS Stream Init error : ", err)
	}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> Enter data to be published: ")
		reader.Scan()
		data := reader.Text()
		if len(data) == 0 {
			continue
		}
		info, err := parseData(data)
		if err != nil {
			fmt.Println("> error: ", err)
			continue
		}

		 publishMessage(js,config.StreamName+"."+info.TrackingId, info)
		

	}

}

func parseData(data string) (models.MessageInfo, error) {
	var messageInfo models.MessageInfo
	err := json.Unmarshal([]byte(data), &messageInfo)
	if err != nil {
		return messageInfo, err
	}
	return messageInfo, nil

}
