package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	eventChan := make(chan string)
	go func() {
		events := []string{"start", "archive", "stop", "start", "stop", "archive"}
		for _, e := range events {
			eventChan <- e
		}
	}()
	defer close(eventChan)

	commands := make(chan ICommand, 1000)
	defer close(commands)

	go func() {
		for {
			if event, ok := <-eventChan; ok {
				switch event {
				case "start":
					commands <- NewStartCommand()
				case "stop":
					commands <- NewStopCommand()
				case "archive":
					commands <- NewArchiveCommand()
				}
			} else {
				return
			}
		}
	}()

	for {
		select {
		case c := <-commands:
			c.Execute()
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1s")
			return
		}
	}
}
