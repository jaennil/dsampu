package main

import (
	"fmt"
	"math"
	"strconv"
)

type customArray struct {
	value []int
	size  int
}

func newCustomArray(size int) customArray {
	array := customArray{make([]int, size), size}

	for i := range size {
		array.value[i] = randInt(math.MinInt32, math.MaxInt32)
	}

	return array
}

func (ca customArray) find(elem int) (int, error) {
	for i := range ca.size {
		if ca.value[i] == elem {
			return i, nil
		}
	}

	return -1, fmt.Errorf("elem %v not found", elem)
}

func (ca *customArray) insert(elem int) {
	ca.value = append(ca.value, elem)
	ca.size++
}

func (ca *customArray) delete(elem int) error {
	array := ca.value
	for i := range ca.size {
		if ca.value[i] == elem {
			for k := i; k < ca.size-1; k++ {
				array[k] = array[k+1]
			}
			ca.size--
			return nil
		}
	}

	return fmt.Errorf("element %v not found", elem)
}

func (ca customArray) String() string {
	var result string

	for i := range ca.size {
		result += strconv.Itoa(ca.value[i]) + " "
	}

	result += fmt.Sprintf("\nsize=%d", ca.size)

	return result
}
