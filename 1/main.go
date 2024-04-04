package main

import (
	"fmt"
	"strconv"
)

func main() {

	size := scanSize()

	customArray := newCustomArray(size)

	mainloop(customArray)
}

func mainloop(customArray customArray) {
	for {
		fmt.Println(customArray)

		fmt.Println("1. find")
		fmt.Println("2. insert")
		fmt.Println("3. delete")

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

			index, err := customArray.find(value)
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

			customArray.insert(value)

			fmt.Println(customArray)
		case delete:
			fmt.Print("entere value to delete: ")

			value, err := scanInt()
			if err != nil {
				fmt.Println("you should enter integer")
				continue
			}

			err = customArray.delete(value)
			if err != nil {
				fmt.Println(err)
			}
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
