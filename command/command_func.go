package command

import "fmt"

type Command func() error

func StartCommandFunc() Command {
	return func() error {
		fmt.Println("start command")
		return nil
	}
}

func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("archive command")
		return nil
	}
}

func StopCommandFunc() Command {
	return func() error {
		fmt.Println("stop command")
		return nil
	}
}
