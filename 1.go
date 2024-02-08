package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	sorted = iota + 1
	unsorted
)

type sortedArray []float64

func (arr sortedArray) find(elem float64) float64 {
	return 0
}

type unsortedArray []float64

func (arr unsortedArray) find(elem float64) float64 {
	return 0
}

type IArray interface {
	find(elem float64) float64
}

func newSortedArray(size int) sortedArray {
	return nil
}

func newUnsortedArray(size int) unsortedArray {
	var array = make([]float64, size)
	for i := 0; i < size; i++ {
		array[i] = randFloat64()
	}
	return array
}

func randFloat64() float64 {
	return rand.Float64() * math.Pow10(308) * (2*rand.Float64() - 1)
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
