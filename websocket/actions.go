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
	updateCardEvent   = "UPDATE_CARD"
	deleteCardEvent   = "DELETE_CARD"
	createColumnEvent = "CREATE_COLUMN"
	updateColumnEvent = "UPDATE_COLUMN"
	deleteColumnEvent = "DELETE_COLUMN"
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
	actions[updateColumnEvent] = updateColumnAction
	actions[deleteColumnEvent] = deleteColumnAction
	actions[createCardEvent] = createCardAction
	actions[updateCardEvent] = updateCardAction
	actions[deleteCardEvent] = deleteCardAction

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

		return handler(db, boardID, e)
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
		return err
	}

	logger.Info("Created Column: " + colID)
	return nil
}

func updateColumnAction(db database.Service, board string, e Event) error {
	columnName := getStringOrPanic(e, "column_name")
	columnID := getStringOrPanic(e, "column_id")

	err := db.UpdateColumn(board, columnID, columnName)
	if err != nil {
		return err
	}

	logger.Info("Updated Column " + columnID)
	return nil
}

func deleteColumnAction(db database.Service, board string, e Event) error {
	columnID := getStringOrPanic(e, "column_id")

	err := db.DeleteColumn(board, columnID)
	if err != nil {
		return err
	}

	logger.Info("Deleted Column " + columnID)
	return nil
}

func createCardAction(db database.Service, board string, e Event) error {
	columnID := getStringOrPanic(e, "column_id")
	cardContent := getStringOrPanic(e, "content")

	createdID, err := db.CreateCard(board, columnID, cardContent)
	if err != nil {
		return err
	}
	logger.Info("Created card: " + createdID)

	return nil
}

func updateCardAction(db database.Service, board string, e Event) error {
	columnID := getStringOrPanic(e, "column_id")
	cardID := getStringOrPanic(e, "card_id")
	cardContent := getStringOrPanic(e, "content")

	err := db.UpdateCard(board, columnID, cardID, cardContent)
	if err != nil {
		return err
	}
	logger.Info("Updated Card " + cardID)

	return nil
}

func deleteCardAction(db database.Service, board string, e Event) error {
	columnID := getStringOrPanic(e, "column_id")
	cardID := getStringOrPanic(e, "card_id")

	err := db.DeleteCard(board, columnID, cardID)
	if err != nil {
		return err
	}

	logger.Info("Deleted Card" + cardID)
	return nil
}
