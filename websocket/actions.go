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
	errNoParam      = errors.New("required.param.not.found")
)

func unmarshallMessage(message []byte) (event Event, err error) {
	err = json.Unmarshal(message, &event)
	return
}

func makeActionsHandler(db database.Service, boardID string) func(e Event) error {
	actions := make(map[string]func(db database.Service, board string, e Event) error)

	actions[createColumnEvent] = createColumnAction

	return func(e Event) error {
		handler, ok := actions[e.Type]

		if !ok {
			return errInvalidEvent
		}

		defer func() {
			if r := recover(); r != nil {
				logger.Error(errNoParam)
			}
		}()

		handler(db, boardID, e)
		return nil
	}
}

func getStringOrPanic(e Event, key string) string {
	v, ok := e.Payload[key]
	vs := v.(string)
	if !ok || vs == "" {
		panic(errNoParam)
	}
	return vs
}

func createColumnAction(db database.Service, board string, e Event) error {
	columnName := getStringOrPanic(e, "column_name")

	colID, err := db.CreateColumn(board, columnName)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Created Column: " + colID)
	return nil
}

func createCardAction(db database.Service, board string, e Event) error {
	columnName := getStringOrPanic(e, "column_name")
	cardContent := getStringOrPanic(e, "content")

	colID, err := db.CreateCard(board, columnName, cardContent)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("Created Column: " + colID)

	return nil
}
