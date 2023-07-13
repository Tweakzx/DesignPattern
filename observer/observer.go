package observer

import (
	"fmt"
)

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify()
}

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer IObserver) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type Observer1 struct{}

func (o *Observer1) Update(msg string) {
	fmt.Printf("Observer1 has received the message: %s!\n", msg)
}

type Observer2 struct{}

func (o *Observer2) Update(msg string) {
	fmt.Printf("Observer2 has received the message:%s!", msg)
}
