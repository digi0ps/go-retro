package logger

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	t.Run("time should be RFC3339", func(t *testing.T) {
		testTime := getTime()
		_, err := time.Parse(time.RFC3339, testTime)

		if err != nil {
			t.Error("getTime didn't return time in correct format")
		}
	})
}
