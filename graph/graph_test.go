package test

import (
	"bufio"
)

type Node struct {
	id   int
	name string
}

type Graph struct {
	array []*Node
}

func NewGraph(scanner *bufio.Scanner, count int) *Graph {
	array := make([]*Node, count, count)
	return &Graph{array}
}
