package beeep

import (
	"testing"
)

func TestAlert(t *testing.T) {
	err := Alert("Alert title", "Message body", "assets/warning.png", "AppID")
	if err != nil {
		t.Error(err)
	}
}
