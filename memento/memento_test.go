package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshot(t *testing.T) {
	in := &InputText{"Hello"}
	snap1 := in.SnapShot()
	in.Append(" World")
	assert.Equal(t, "Hello World", in.GetText())

	in.Restore(snap1)
	assert.Equal(t, "Hello", in.GetText())
}
