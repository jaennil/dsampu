package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type tree struct {
	root *node
}

func NewTree() *tree {
	return new(tree)
}

func (tree *tree) insert(value int) {
	newNode := new(node)
	newNode.value = value

	if tree.root == nil {
		tree.root = newNode
		return
	}

	current := tree.root
	var parent *node

	for {
		parent = current

		if value < current.value {
			current = current.left
			if current == nil {
				parent.left = newNode
				return
			}
		} else {
			current = current.right
			if current == nil {
				parent.right = newNode
				return
			}
		}
	}
}

func (tree tree) String() string {

	result := ""

	terminalWidth, err := getTerminalWidth()
	if err != nil {
		log.Fatal("error occurred while getting terminal width:", err)
	}
	center := terminalWidth / 2

	current := tree.root
	if current == nil {
		return ""
	}

	result += strings.Repeat(" ", center-digitLength(current.value)/2)
	result += strconv.Itoa(current.value) + "\n"

	// unvisited := make([]*node, 1)
	level := make([]*node, 1)

	level = append(level, current)

	// for {
		if current.hasLeft() || current.hasRight() {
			result += strings.Repeat(" ", center-digitLength(current.value)/2)
			if current.hasLeft() {
				result += "/ "
			}

			if current.hasRight() {
				result += "\\ "
			}

			result += "\n"
		}
	// }

	return result
}

func digitLength(value int) int {
	valueString := strconv.Itoa(value)
	return len(valueString)
}

func getTerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	outputBytes, err := cmd.Output()
	if err != nil {
		return -1, err
	}

	output := string(outputBytes)

	trimmedOutput := strings.TrimSpace(output)

	splittedOutput := strings.Split(trimmedOutput, " ")

	widthString := splittedOutput[1]

	width, err := strconv.Atoi(widthString)
	if err != nil {
		return -1, err
	}

	return width, nil
}
