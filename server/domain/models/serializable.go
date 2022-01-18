package models

import (
	"encoding/json"
	"log"

	"peeral.com/proxy-libp2p/domain/enums"
	"peeral.com/proxy-libp2p/domain/interfaces"
)

type logEvent struct {
	EventType enums.EventType
	Data      interfaces.Serializable
}

type logData struct {
	Message string
}

type messageData struct {
	Message string
}

func newLogEvent(eventType enums.EventType, data interfaces.Serializable) *logEvent {
	return &logEvent{EventType: eventType, Data: data}
}

func (l logEvent) Serialize() string {
	return jsonSerialize(l)
}

func NewLogData(message string) *logEvent {
	data := logData{Message: message}
	return newLogEvent(enums.INFO, data)
}

func (l logData) Serialize() string {
	return jsonSerialize(l)
}

func NewMessageData(message string) *logEvent {
	data := messageData{Message: message}
	return newLogEvent(enums.MESSAGE, data)
}

func (m messageData) Serialize() string {
	return jsonSerialize(m)
}

func jsonSerialize(s interfaces.Serializable) string {
	j, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	return string(j)
}
