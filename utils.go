package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func DeleteStreamData(js nats.JetStreamContext, streamName string) error {
	err := js.PurgeStream(streamName)
	if err != nil {
		log.Fatal("Failed to purge stream:", err)
	}

	log.Println("All messages deleted from the stream : ", streamName)
	return nil
}
