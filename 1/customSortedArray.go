package main

import (
	"fmt"
	"sort"
)

type customSortedArray struct {
	*customArray
}

func newCustomSortedArray(size int) *customSortedArray {
	customArray := newCustomArray(size)

	sort.Ints(customArray.value)

	customSortedArray := customSortedArray{customArray}

	return &customSortedArray
}

func (csa *customSortedArray) insert(elem int) {
	var i int
	for i = range csa.size {
		if csa.value[i] > elem {
			break
		}
	}

	fmt.Println(i)

	csa.value = append(csa.value, csa.value[len(csa.value)-1])

	for k := csa.size-1; k > i; k-- {
		csa.value[k] = csa.value[k-1]
	}

	csa.value[i] = elem
	csa.size++
}

func (csa customSortedArray) bfind(elem int) (int, error) {
	var lowerBound int = 0
	var upperBound int = csa.size - 1
	var curln int

	for {
		curln = (lowerBound + upperBound) / 2
		if csa.value[curln] == elem {
			return curln, nil
		} else if lowerBound > upperBound {
			return -1, fmt.Errorf("element not found")
		} else {
			if csa.value[curln] < elem {
				lowerBound = curln + 1
			} else {
				upperBound = curln - 1
			}
		}
	}

}
