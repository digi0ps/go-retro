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

const (
	createCardEvent   = "CREATE_CARD"
	createColumnEvent = "CREATE_COLUMN"
)

var (
	errInvalidEvent = errors.New("invalid.event")
)

func unmarshallMessage(message []byte) (event Event, err error) {
	err = json.Unmarshal(message, &event)
	return
}

func makeActionsHandler(db database.Service, boardID string) func(e Event) error {
	actions := make(map[string]func(db database.Service, e Event) error)

	actions[createColumnEvent] = createColumnAction

	return func(e Event) error {
		handler, ok := actions[e.Type]

		if !ok {
			return errInvalidEvent
		}

		handler(db, e)
		return nil
	}
}

func createColumnAction(db database.Service, e Event) error {
	boardID := e.Payload["board_id"].(string)
	columnName := e.Payload["column_name"].(string)

	if boardID != "" && columnName != "" {
		colID, _ := db.CreateColumn(boardID, columnName)
		logger.Info("Created Column: " + colID)
	}

	return nil
}
