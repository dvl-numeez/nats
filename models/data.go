package models

import "time"

type MessageInfo struct {
	TrackingId string    `json:"trackingId"`
	Content    string    `json:"content"`
	Timestamp  time.Time `json:"timestamp"`
}
