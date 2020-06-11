package websocket

import (
	"encoding/json"
	"errors"
	"go-retro/database"
	"go-retro/logger"
)

// Event represents an event received through websocket
type Event struct {
	Type      string                 `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
	Timestamp int64                  `json:"timestamp,omitempty"`
}

var (
	createCardEvent   = "CREATE_CARD"
	createColumnEvent = "CREATE_COLUMN"
	errInvalidEvent   = errors.New("Invalid Event")
)

func unmarshallMessage(message []byte) (event Event, err error) {
	err = json.Unmarshal(message, &event)
	return
}

func processEvent(db database.Service, event Event) error {
	logger.Info("Processing action")
	switch event.Type {
	case createColumnEvent:
		boardID := event.Payload["board_id"].(string)
		columnName := event.Payload["column_name"].(string)
		if boardID != "" && columnName != "" {
			colID, _ := db.CreateColumn(boardID, columnName)
			logger.Info("Created Column: " + colID)
		}
	default:
		return errInvalidEvent
	}
	return nil
}
