package test

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	v      *Vertex
	next   *node
	weight float64
}

type Vertex struct {
	id    int64
	value interface{}
	node  *node
}

type Edge struct {
	start  *Vertex
	end    *Vertex
	weight float64
}

type Graph struct {
	vertices []*Vertex
	edges    []*Edge
}

func (g *Graph) Walk() {
	for _, v := range g.vertices {
		fmt.Println(format(v))
	}
}

func NewVertex(id int64, name string) *Vertex {
	return &Vertex{id, name, nil}
}

func NewEdge(v1, v2 *Vertex, weight float64) *Edge {
	return &Edge{v1, v2, weight}
}

func NewDirectGraph(scanner *bufio.Scanner, vcount, ecount int) *Graph {
	vertices := make([]*Vertex, vcount, vcount)
	edges := make([]*Edge, 0, ecount)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			a, b, w := readline(scanner.Text())
			id1 := getID(a)
			id2 := getID(b)
			var v1, v2 *Vertex
			if vertices[id1] == nil {
				v1 = NewVertex(id1, a)
				vertices[id1] = v1
			} else {
				v1 = vertices[id1]
			}
			if vertices[id2] == nil {
				v2 = NewVertex(id2, b)
				vertices[id2] = v2
			} else {
				v2 = vertices[id2]
			}
			node := &node{v2, nil, w}
			addTail(v1, node)
			edges = append(edges, NewEdge(v1, v2, w))
		}
	}
	return &Graph{vertices, edges}
}

func addTail(head *Vertex, node *node) {
	if head.node == nil {
		head.node = node
	} else {
		n := head.node
		for n.next != nil {
			n = n.next
		}
		n.next = node
	}
}

func readline(text string) (string, string, float64) {
	line := strings.Split(text, ",")
	weight, _ := strconv.ParseFloat(line[2], 64)
	return line[0], line[1], weight
}

func format(head *Vertex) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("vertex: %d-%s   ", head.id, head.value))
	p := head.node
	for p != nil {
		b.WriteString(fmt.Sprintf("%d-%s ", p.v.id, p.v.value))
		p = p.next
	}
	return b.String()
}

func getID(name string) int64 {
	i, _ := strconv.ParseInt(name, 10, 64)
	return i
}
