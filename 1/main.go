package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	custom_array = iota + 1
	custom_sorted_array
)

func main() {

	size := scanSize()

	array := createArray(size)

	mainloop(array)
}

func createArray(size int) IArray {
	fmt.Println("1. unsorted array")
	fmt.Println("2. sorted array")

	fmt.Print("enter array type: ")

	arrayType, err := scanInt()
	if err != nil {
		fmt.Println("you should enter integer")
		return createArray(size)
	}

	switch arrayType {
	case custom_array:
		return newCustomArray(size)
	case custom_sorted_array:
		return newCustomSortedArray(size)
	default:
		fmt.Println("wrong array type. try again")
		return createArray(size)
	}
}

func mainloop(array IArray) {
	for {
		fmt.Println("1. find")
		fmt.Println("2. insert")
		fmt.Println("3. delete")
		fmt.Println("4. print array")

		if _, ok := array.(*customSortedArray); ok {
			fmt.Println("5. binary search")
		}

		operation, err := scanInt()
		if err != nil {
			fmt.Println("you should enter integer")
			continue
		}

		switch operation {
		case find:
			fmt.Print("enter value to find: ")

			value, err := scanInt()
			if err != nil {
				fmt.Println("you should enter integer")
				continue
			}

			start := time.Now()

			index, err := array.find(value)

			took_time := time.Since(start)

			fmt.Printf("operation took %v\n", took_time)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("found value %v at index: %v\n", value, index)

		case insert:
			fmt.Print("enter value to insert: ")

			value, err := scanInt()
			if err != nil {
				fmt.Println("you should enter integer")
				continue
			}

			start := time.Now()

			array.insert(value)

			took_time := time.Since(start)

			fmt.Printf("element %v inserted successfully\n", value)

			fmt.Printf("operation took %v\n", took_time)

		case delete:
			fmt.Print("enter value to delete: ")

			value, err := scanInt()
			if err != nil {
				fmt.Println("you should enter integer")
				continue
			}

			start := time.Now()

			err = array.delete(value)

			took_time := time.Since(start)

			fmt.Printf("operation took %v\n", took_time)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("element %v deleted successfully\n", value)

		case print:
			fmt.Println(array)
		case bfind:
			switch castedArray := array.(type) {
			case *customArray:
				fmt.Println("binary search is not supported for unsorted array")
				continue
			case *customSortedArray:
				fmt.Print("enter value to find: ")

				value, err := scanInt()
				if err != nil {
					fmt.Println("you should enter integer")
					continue
				}

				start := time.Now()

				index, err := castedArray.bfind(value)

				took_time := time.Since(start)

				fmt.Printf("operation took %v\n", took_time)

				if err != nil {
					fmt.Println(err)
					continue
				}

				fmt.Printf("using binary search found value %v at index: %v\n", value, index)
			}
		default:
			fmt.Println("wrong operation. try again")
			continue
		}
	}
}

func scanInt() (int, error) {
	var variable string
	if _, err := fmt.Scanln(&variable); err != nil {
		return -1, err
	}

	intVariable, err := strconv.Atoi(variable)
	if err != nil {
		return -1, err
	}

	return intVariable, nil
}

func scanSize() int {
	for {
		fmt.Print("array size: ")

		size, err := scanInt()
		if err != nil {
			fmt.Println("you should enter integer")
			continue
		}

		return size
	}
}
