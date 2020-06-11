package websocket

import (
	"encoding/json"
	"fmt"
	"go-retro/database"
	"go-retro/logger"
)

// Event represents an event received through websocket
type Event struct {
	Type      string                 `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
	Timestamp int64                  `json:"timestamp,omitempty"`
}

func unmarshallMessage(message []byte) (event Event, err error) {
	err = json.Unmarshal(message, &event)
	return
}

func processEvent(db database.Service, event Event) error {
	logger.Info("Processing action")
	fmt.Println(event)
	return nil
}
