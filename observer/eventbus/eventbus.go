package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.RWMutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: make(map[string][]reflect.Value),
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, handler interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(handler)
	if v.Kind() != reflect.Func {
		return fmt.Errorf("handler must be a function")
	}

	if _, ok := bus.handlers[topic]; !ok {
		bus.handlers[topic] = make([]reflect.Value, 0)
	}

	bus.handlers[topic] = append(bus.handlers[topic], v)

	return nil
}

func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	bus.lock.RLock()
	defer bus.lock.RUnlock()

	handlers, ok := bus.handlers[topic]
	if !ok {
		fmt.Println("no handlers for  topic:", topic)
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}
	for i := range handlers {
		go handlers[i].Call(params)
	}
}
