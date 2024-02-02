package main

import (
	"fmt"
	"math/rand"
)

const (
	sorted = iota + 1
	unsorted
)

type sortedArray []int

func (arr sortedArray) find(elem int) int {
	return 0
}

type unsortedArray []int

func (arr unsortedArray) find(elem int) int {
	return 0
}

type IArray interface {
	find(elem int) int
}

func newSortedArray(size int) sortedArray {
	return nil
}

func newUnsortedArray(size int) unsortedArray {
	var array = make([]int, size)
	for i := range array {
		array[i] = rand.Intn(100)
	}
	return array
}

func main() {
	scanArrayType:
	fmt.Println("1. sorted array")
	fmt.Println("2. unsorted array")

	var arrayType int
	_, err := fmt.Scanln(&arrayType)
	if err != nil {
		panic(err)
	}

	var size int
	fmt.Print("array size?: ")
	_, err = fmt.Scanln(&size)
	if err != nil {
		panic(err)
	}

	var array IArray
	switch arrayType {
	case sorted:
		array = newSortedArray(size)
	case unsorted:
		array = newUnsortedArray(size)
	default:
		fmt.Println("unknown array type try again")
		goto scanArrayType
	}

	fmt.Println(array)
}
