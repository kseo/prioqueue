package main

import (
	"fmt"

	"github.com/kseo/prioqueue"
)

func intCmpFunc(a, b interface{}) int {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) > b.(int) {
		return 1
	} // else
	return 0
}

func sort(xs []int) []int {
	var result []int
	elements := make([]interface{}, len(xs))
	for i, d := range xs {
		elements[i] = d
	}
	pq := prioqueue.NewPriorityQueue(elements, intCmpFunc)
	for !pq.IsEmpty() {
		max, _ := pq.RemoveMax()
		result = append(result, max.(int))
	}
	return result
}

func main() {
	xs := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 9}
	fmt.Printf("%v\n", sort(xs))
}
