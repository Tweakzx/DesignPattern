package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection_Iterator(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}

	collection := func(array []int) Collection {
		collection := make([]interface{}, len(array))
		for i, v := range array {
			collection[i] = v
		}
		return collection
	}(array)

	for it, i := collection.Iterator(), 0; it.HasNext(); i++ {
		assert.Equal(t, array[i], it.Next())
	}
}
