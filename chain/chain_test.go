package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain_Filter(t *testing.T) {
	tests1 := []struct {
		name     string
		constant string
		want     bool
	}{
		{"safe", "今天天气多云转阴", false},
		{"sexy", "色情内容", true},
		{"politic", "政治敏感内容", false},
	}

	tests2 := []struct {
		name     string
		constant string
		want     bool
	}{
		{"safe", "今天天气多云转阴", false},
		{"sexy", "色情内容", true},
		{"politic", "政治敏感内容", true},
	}

	chain := NewSensitiveWordChain(&SexyWordFilter{})
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got := chain.Filter(tt.constant)
			assert.Equal(t, tt.want, got)
		})
	}

	chain.AddFilter(&PoliticalWordFilter{})
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := chain.Filter(tt.constant)
			assert.Equal(t, tt.want, got)
		})
	}
}
