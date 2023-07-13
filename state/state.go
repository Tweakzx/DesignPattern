package state

import (
	"fmt"
)

type State interface {
	Handle(i *Issue)
	GetName() string
}

// 定义状态
type Backlog struct{}

func (s *Backlog) GetName() string {
	return "Backlog"
}

func (s *Backlog) Handle(i *Issue) {
	fmt.Println("BACKLOG")
	i.SetState(&Todo{})
}

type Todo struct{}

func (s *Todo) GetName() string {
	return "Todo"
}

func (s *Todo) Handle(i *Issue) {
	fmt.Println("TODO")
	i.SetState(&Doing{})
}

type Doing struct{}

func (s *Doing) GetName() string {
	return "Doing"
}

func (s *Doing) Handle(i *Issue) {
	fmt.Println("DOING")
	i.SetState(&Done{})
}

type Done struct{}

func (s *Done) GetName() string {
	return "Done"
}

func (s *Done) Handle(i *Issue) {
	fmt.Println("DONE")
}

// 定义事项
type Issue struct {
	state State
}

func NewIssue() *Issue {
	return &Issue{
		state: &Backlog{},
	}
}

func (s *Issue) SetState(state State) {
	s.state = state
}

func (s *Issue) NextStatus() {
	s.state.Handle(s)
}
