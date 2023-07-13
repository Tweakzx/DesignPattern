package command

import (
	"fmt"
)

type ICommand interface {
	Execute() error
}

type StartCommand struct{}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (s *StartCommand) Execute() error {
	fmt.Println("StartCommand: Execute")
	return nil
}

type ArchiveCommand struct{}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}

func (a *ArchiveCommand) Execute() error {
	fmt.Println("ArchiveCommand: Execute")
	return nil
}

type StopCommand struct{}

func NewStopCommand() *StopCommand {
	return &StopCommand{}
}

func (s *StopCommand) Execute() error {
	fmt.Println("StopCommand: Execute")
	return nil
}
