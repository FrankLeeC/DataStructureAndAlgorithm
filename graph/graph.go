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
	weight *float64
}

type Vertex struct {
	value interface{}
	node  *node
}

func (v *Vertex) Value() interface{} {
	return v.value
}

type Edge struct {
	start  *Vertex
	end    *Vertex
	weight *float64
}

type Graph struct {
	vertices []*Vertex
	edges    []*Edge
}

// V get count of vertices
func (g *Graph) V() int {
	return len(g.vertices)
}

// E get count of edges
func (g *Graph) E() int {
	return len(g.edges)
}

func (g *Graph) ContainsVertex(v interface{}) (bool, *Vertex) {
	for _, tmp := range g.vertices {
		if tmp == nil {
			return false, nil
		}
		if tmp.value == v {
			return true, tmp
		}
	}
	return false, nil
}

// AddEdge add edge
func (g *Graph) AddEdge(e *Edge) {
	start := e.start
	node := &node{e.end, nil, e.weight}
	if c, v := g.ContainsVertex(start.value); c {
		start = v
	}
	addTail(start, node)
}

func (g *Graph) Adj(v *Vertex) []*Vertex {
	return nil
}

func (g *Graph) String() string {
	var b strings.Builder
	for _, v := range g.vertices {
		b.WriteString(format(v) + "\n")
	}
	return b.String()
}

func NewVertex(value string) *Vertex {
	return &Vertex{value, nil}
}

func NewEdge(v1, v2 *Vertex, weight *float64) *Edge {
	return &Edge{v1, v2, weight}
}

func NewDirectGraph(scanner *bufio.Scanner, vcount, ecount int) *Graph {
	vertices := make([]*Vertex, vcount, vcount)
	edges := make([]*Edge, 0, ecount)
	count := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			a, b, w := readline(scanner.Text())
			var v1, v2 *Vertex
			if c, v := containsVertex(a, vertices); c {
				v1 = v
			} else {
				v1 = NewVertex(a)
				if count >= vcount {
					vertices = append(vertices, v1)
				} else {
					vertices[count] = v1
					count++
				}
			}
			if c, v := containsVertex(b, vertices); c {
				v2 = v
			} else {
				v2 = NewVertex(b)
				if count >= vcount {
					vertices = append(vertices, v2)
				} else {
					vertices[count] = v2
					count++
				}
			}
			node := &node{v2, nil, &w}
			addTail(v1, node)
			edges = append(edges, NewEdge(v1, v2, &w))
		}
	}
	return &Graph{vertices, edges}
}

func containsVertex(v interface{}, array []*Vertex) (bool, *Vertex) {
	for _, tmp := range array {
		if tmp == nil {
			return false, nil
		}
		if tmp.value == v {
			return true, tmp
		}
	}
	return false, nil
}

func equalsVertex(v1, v2 *Vertex) bool {
	return v1.value == v2.value
}

func addTail(head *Vertex, node *node) {
	if head.value == node.v.value {
		return
	}
	if head.node == nil {
		head.node = node
	} else {
		n := head.node
		for n.next != nil {
			if equalsVertex(n.v, node.v) {
				return
			}
			n = n.next
		}
		if equalsVertex(n.v, node.v) {
			return
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
	b.WriteString(fmt.Sprintf("vertex: %s   ", head.value))
	p := head.node
	for p != nil {
		b.WriteString(fmt.Sprintf("%s[%f] ", p.v.value, *p.weight))
		p = p.next
	}
	return b.String()
}

func getID(name string) int64 {
	i, _ := strconv.ParseInt(name, 10, 64)
	return i
}
