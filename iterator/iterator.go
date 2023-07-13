package iterator

// 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// 具体迭代器
type CollectionIterator struct {
	collection *Collection
	index      int
}

func (it *CollectionIterator) HasNext() bool {
	return it.index < len(*(it.collection))
}

func (it *CollectionIterator) Next() interface{} {
	item := (*it.collection)[it.index]
	it.index++
	return item
}

// 具体集合
type Collection []interface{}

func (c *Collection) Iterator() Iterator {
	return &CollectionIterator{collection: c, index: 0}
}
