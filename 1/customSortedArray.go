package main

import "sort"

type customSortedArray struct {
	*customArray
}

func NewCustomSortedArray(size int) *customSortedArray {
	customArray := newCustomArray(size)

	sort.Ints(customArray.value)

	customSortedArray := customSortedArray{customArray}

	return &customSortedArray
}
