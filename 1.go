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

const (
	find = iota + 1
	insert
	delete_
	bfind
)

type IArray interface {
	find(elem int) (int, error)
	insert(elem, index int) error
	delete(elem int) error
	binary_search(elem int) (int, error)
}

type sortedArray []int

func newSortedArray(size int) *sortedArray {
	// cant use array because var array[size]int (size should be known at compile time)
	// https://stackoverflow.com/questions/8539551/dynamically-initialize-array-size-in-go
	var array = make(sortedArray, size)
	for i := range array {
		if i == 0 {
			array[0] = randInt(-10*size, 10*size)
			continue
		}

		array[i] = randInt(array[i-1], 10*size)
	}
	return &array
}

func (arr *sortedArray) find(elem int) (int, error) {
	for i, v := range *arr {
		if v == elem {
			return i, nil
		}
	}
	return -1, fmt.Errorf("element %v not found", elem)
}

func (arr *sortedArray) binary_search(elem int) (int, error) {
	var left int = 0
	var right int = len(*arr) - 1
	for left <= right {
		var middle int = int(math.Floor((float64(left) + float64(right)) / 2.0))
		if (*arr)[middle] < elem {
			left = middle + 1
		} else if (*arr)[middle] > elem {
			right = middle - 1
		} else {
			return middle, nil
		}
	}
	return -1, fmt.Errorf("element %v not found")
}

func (arr *sortedArray) insert(elem, index int) error {
	if index < 0 || index > len(*arr) {
		return fmt.Errorf("index out of bounds")
	}
	var newarray = make([]int, len(*arr)+1)
	for i := range index {
		newarray[i] = (*arr)[i]
	}
	newarray[index] = elem
	for i := index + 1; i < len(newarray); i++ {
		newarray[i] = (*arr)[i-1]
	}
	*arr = make([]int, len(newarray))
	for i, v := range newarray {
		(*arr)[i] = v
	}
	return nil
}

func (arr *sortedArray) delete(elem int) error {
	newarray := make([]int, len(*arr)-1)
	var j int
	var flag bool
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == elem && !flag {
			flag = true
			continue
		}
		newarray[j] = (*arr)[i]
		j++
	}

	if !flag {
		return fmt.Errorf("element not fould %v, elem")
	}

	*arr = make([]int, len(newarray))
	for i, v := range newarray {
		(*arr)[i] = v
	}
	return nil
}

type unsortedArray []int

func newUnsortedArray(size int) *unsortedArray {
	var array = make(unsortedArray, size)
	for i := range array {
		array[i] = randInt(-10*size, 10*size)
	}
	return &array
}

func (arr *unsortedArray) find(elem int) (int, error) {
	for i, v := range *arr {
		if v == elem {
			return i, nil
		}
	}
	return -1, fmt.Errorf("element %v not found", elem)
}

func (arr *unsortedArray) binary_search(elem int) (int, error) {
	var left int = 0
	var right int = len(*arr) - 1
	for left <= right {
		var middle int = int(math.Floor((float64(left) + float64(right)) / 2.0))
		if (*arr)[middle] < elem {
			left = middle + 1
		} else if (*arr)[middle] > elem {
			right = middle - 1
		} else {
			return middle, nil
		}
	}
	return -1, fmt.Errorf("element %v not found")
}

func (arr *unsortedArray) insert(elem, index int) error {
	if index < 0 || index > len(*arr) {
		return fmt.Errorf("index out of bounds")
	}
	var newarray = make([]int, len(*arr)+1)
	for i := range index {
		newarray[i] = (*arr)[i]
	}
	newarray[index] = elem
	for i := index + 1; i < len(newarray); i++ {
		newarray[i] = (*arr)[i-1]
	}
	*arr = make([]int, len(newarray))
	for i, v := range newarray {
		(*arr)[i] = v
	}
	return nil
}

func (arr *unsortedArray) delete(elem int) error {
	newarray := make([]int, len(*arr)-1)
	var j int
	var flag bool
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == elem && !flag {
			flag = true
			continue
		}
		newarray[j] = (*arr)[i]
		j++
	}

	if !flag {
		return fmt.Errorf("element not fould %v, elem")
	}

	*arr = make([]int, len(newarray))
	for i, v := range newarray {
		(*arr)[i] = v
	}
	return nil
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
scanArrayType:

	fmt.Println("select array type:")
	fmt.Println("1. sorted array")
	fmt.Println("2. unsorted array")
	fmt.Print(">> ")
	var arrayType int
	fmt.Scanln(&arrayType)

	fmt.Print("array size?: ")
	var size int
	fmt.Scanln(&size)

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

	fmt.Printf("generated array: %v\n", array)

scanOperation:
	fmt.Println("select operation:")
	fmt.Println("1. find")
	fmt.Println("2. insert")
	fmt.Println("3. delete")
	fmt.Println("4. binary search")
	fmt.Print(">> ")
	var operation int
	fmt.Scanln(&operation)
	switch operation {
	case find:
		fmt.Print("value to find: ")
		var value int
		fmt.Scanln(&value)
		index, err := array.find(value)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("founded value %v at index %v\n", value, index)
	case bfind:
		fmt.Print("value to find: ")
		var value int
		fmt.Scanln(&value)
		index, err := array.binary_search(value)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("founded value %v at index %v\n", value, index)
	case insert:
		fmt.Print("value to insert: ")
		var value int
		fmt.Scanln(&value)
		fmt.Print("index: ")
		var index int
		fmt.Scanln(&index)
		err := array.insert(value, index)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("inserted value %v at index %v\n", value, index)
		fmt.Println("updated array:", array)
	case delete_:
		fmt.Print("value to delete: ")
		var value int
		fmt.Scanln(&value)
		err := array.delete(value)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("deleted value %v\n", value)
		fmt.Println("updated array:", array)
	default:
		fmt.Println("unknown operation try again")
		goto scanOperation
	}
	goto scanOperation
}
