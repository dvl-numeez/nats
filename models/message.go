package models

import (
	"encoding/json"
	"os"
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func GetMessages() ([]Message, error) {
	rawReviews, err := os.ReadFile("/Users/numeezbaloch17/Documents/nats/models/message.json")
	if err != nil {
		return []Message{}, err
	}
	var messageObject []Message
	err = json.Unmarshal(rawReviews, &messageObject)

	return messageObject, err
}
