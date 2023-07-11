package builder

import (
	"fmt"
)

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name should not be empty")
	}
	b.name = name
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		panic("maxTotal should be greater than 0")
	}
	b.maxTotal = maxTotal
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("maxIdle should be greater than or equal to 0")
	}
	b.maxIdle = maxIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("minIdle should be greater than or equal to 0")
	}
	b.minIdle = minIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.name == "" {
		return nil, fmt.Errorf("name should not be empty")
	}
	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
	}
	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
	}
	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
	}
	if b.maxTotal < b.maxIdle {
		return nil, fmt.Errorf("maxTotal(%d) should be greater than or equal to maxIdle(%d)", b.maxTotal, b.maxIdle)
	}
	if b.minIdle > b.maxIdle {
		return nil, fmt.Errorf("minIdle(%d) should be less than or equal to maxIdle(%d)", b.minIdle, b.maxIdle)
	}
	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}
