package eventbus

import (
	"fmt"
	"testing"
	"time"
)

func TestEventbus(t *testing.T) {
	sub1 := func(msg1, msg2 string) {
		fmt.Println("sub1:", msg1, msg2)
	}

	sub2 := func(msg1, msg2 string) {
		fmt.Println("sub2:", msg1, msg2)
	}

	bus := NewAsyncEventBus()
	bus.Subscribe("topic1", sub1)
	bus.Subscribe("topic1", sub2)

	bus.Publish("topic1", "hello, ", "world")
	time.Sleep(1 * time.Second)

	bus.Publish("topic1", "你好, ", "世界")
	time.Sleep(1 * time.Second)
}
