package prioqueue

import (
	"math/rand"
	"testing"

	"github.com/golangplus/testing/assert"
)

func intCmpFunc(a, b interface{}) int {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) > b.(int) {
		return 1
	} // else
	return 0
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue([]interface{}{}, intCmpFunc)
	for i := 0; i < 1000; i++ {
		pq.Add(rand.Int())
	}
	assert.Equal(t, "pq.Len()", pq.Len(), 1000)

	peek, err := pq.Peek()
	assert.Equal(t, "no error", err, nil)
	peekVal := peek.(int)

	max, err := pq.RemoveMax()
	assert.Equal(t, "no error", err, nil)
	maxVal := max.(int)

	assert.Equal(t, "pq.Peek()", peekVal, maxVal)

	for i := 1; i < 1000; i++ {
		cur, err := pq.RemoveMax()
		assert.Equal(t, "no error", err, nil)
		curVal := cur.(int)

		if curVal > maxVal {
			t.Errorf("%d should be larger than %d", maxVal, curVal)
		}
		maxVal = curVal
	}
}
