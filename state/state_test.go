package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_Handle(t *testing.T) {
	isu := NewIssue()
	assert.Equal(t, "Backlog", isu.state.GetName())

	isu.NextStatus()
	assert.Equal(t, "Todo", isu.state.GetName())

	isu.NextStatus()
	assert.Equal(t, "Doing", isu.state.GetName())

	isu.NextStatus()
	assert.Equal(t, "Done", isu.state.GetName())
}
