package test

import (
	"code.project.com/InstantMessaging/pkg/config"
	"testing"
)

func TestConfig(t *testing.T) {
	err := config.LoadConfig("../config.ini")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(config.Config)
}
