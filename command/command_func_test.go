package command

import (
	"testing"
	"time"
)

func TestCommandFunc(t *testing.T) {
	eventChan := make(chan string, 10)
	go func() {
		events := []string{"start", "stop", "archive", "start", "stop"}
		for _, event := range events {
			eventChan <- event
		}
	}()
	defer close(eventChan)

	commands := make(chan Command, 10)
	go func() {
		for {
			event, ok := <-eventChan
			if ok {
				switch event {
				case "start":
					commands <- StartCommandFunc()
				case "stop":
					commands <- StopCommandFunc()
				case "archive":
					commands <- ArchiveCommandFunc()
				default:
					t.Log("unknown event")
				}
			}
		}
	}()
	defer close(commands)

	for {
		select {
		case c := <-commands:
			c()
		case <-time.After(time.Second * 1):
			t.Log("timeout 1s")
			return
		}
	}
}
