package command

import (
	"testing"
)

const (
	cmdText = "give"
	desc    = "Gives x y to player z"
	usage   = "Usage text"
	message = "/give 1 64 admin"
)

func TestCommandFramework(t *testing.T) {
	cf := NewCommandFramework("/")
	cmdHandler := func(msg string) {
		if msg != message {
			t.Errorf("Input message %s is not equal to received message %s .", message, msg)
		}
	}
	cmd := NewCommand(cmdText, desc, usage, cmdHandler)
	cf.AddCommand(cmd)
	cf.Message <- message
}
