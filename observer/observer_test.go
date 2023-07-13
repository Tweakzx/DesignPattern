package observer

import (
	"testing"
)

func TestOberser_Notify(t *testing.T) {
	ob1 := &Observer1{}
	ob2 := &Observer2{}
	sub := &Subject{}
	sub.Register(ob1)
	sub.Register(ob2)
	sub.Notify("hi")
}
