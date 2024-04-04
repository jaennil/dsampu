package main

import "math/rand"

const (
	find = iota + 1
	insert
	delete
	print
	bfind
)

type IArray interface {
	find(elem int) (int, error)
	insert(elem int)
	delete(elem int) error
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}
