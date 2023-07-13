package chain

import "strings"

type SensitiveWordFilter interface {
	Filter(content string) bool
}

type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func NewSensitiveWordChain(filters ...SensitiveWordFilter) *SensitiveWordFilterChain {
	sc := &SensitiveWordFilterChain{}
	sc.filters = append(sc.filters, filters...)
	return sc
}

func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

type SexyWordFilter struct{}

func (f *SexyWordFilter) Filter(content string) bool {
	return strings.Contains(content, "色情")
}

type PoliticalWordFilter struct{}

func (f *PoliticalWordFilter) Filter(content string) bool {
	return strings.Contains(content, "政治")
}
