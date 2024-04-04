package main

import (
	"fmt"
	"strconv"
	"time"
)

const print_array = 4

func main() {

	size := scanSize()

	customArray := newCustomArray(size)

	mainloop(customArray)
}

func mainloop(customArray customArray) {
	for {
		fmt.Println("1. find")
		fmt.Println("2. insert")
		fmt.Println("3. delete")
		fmt.Println("4. print array")

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

			index, err := customArray.find(value)

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

			customArray.insert(value)

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

			err = customArray.delete(value)

			took_time := time.Since(start)

			fmt.Printf("operation took %v\n", took_time)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("element %v deleted successfully\n", value)

		case print_array:
			fmt.Println(customArray)
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
