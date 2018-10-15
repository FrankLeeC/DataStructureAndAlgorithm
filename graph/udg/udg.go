package udg

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/FrankLeeC/DataStructureAndAlgorithm/graph"
)

type node struct {
	v      graph.Vertex
	next   *node
	weight *interface{}
}

type udgVertex struct {
	value *interface{}
	node  *node
}

func (v *udgVertex) Value() *interface{} {
	return v.value
}

func (v *udgVertex) UniqueID() string {
	return ""
}

type udgEdge struct {
	start  graph.Vertex
	end    graph.Vertex
	weight *interface{}
}

func (e *udgEdge) Start() graph.Vertex {
	return e.start
}

func (e *udgEdge) End() graph.Vertex {
	return e.end
}

func (e *udgEdge) Weight() *interface{} {
	return e.weight
}

// UDG undirected graph
type UDG struct {
	vertices []*udgVertex
	edges    []*udgEdge
}

// V get count of vertices
func (g *UDG) V() int {
	return len(g.vertices)
}

// E get count of edges
func (g *UDG) E() int {
	return len(g.edges)
}

func (g *UDG) containsVertex(v interface{}) (bool, graph.Vertex) {
	for _, tmp := range g.vertices {
		if tmp == nil {
			return false, nil
		}
		if *(tmp.Value()) == v {
			return true, tmp
		}
	}
	return false, nil
}

// AddEdge add edge
func (g *UDG) AddEdge(e graph.Edge) {
	start := e.Start()
	n := &node{e.End(), nil, e.Weight()}
	if c, v := g.containsVertex(*(start.Value())); c {
		start = v
	}
	addTail(start, n)

	end := e.End()
	n = &node{e.Start(), nil, e.Weight()}
	if c, v := g.containsVertex(end.Value()); c {
		end = v
	}
	addTail(end, n)
}

// Adj adjancies of vertex
func (g *UDG) Adj(v graph.Vertex) []graph.Vertex {
	return nil
}

// String string
func (g *UDG) String() string {
	var b strings.Builder
	for _, v := range g.vertices {
		b.WriteString(format(v) + "\n")
	}
	return b.String()
}

func newVertex(value interface{}) *udgVertex {
	return &udgVertex{&value, nil}
}

func newEdge(v1, v2 graph.Vertex, weight *interface{}) *udgEdge {
	return &udgEdge{v1, v2, weight}
}

// NewUDG new undirect graph
func NewUDG(scanner *bufio.Scanner, vcount, ecount int) *UDG {
	vertices := make([]*udgVertex, vcount, vcount)
	edges := make([]*udgEdge, 0, ecount)
	count := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			a, b, w := readline(scanner.Text())
			var v1, v2 *udgVertex
			if c, v := containsVertex(a, vertices); c {
				v1 = v
			} else {
				v1 = newVertex(a)
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
				v2 = newVertex(b)
				if count >= vcount {
					vertices = append(vertices, v2)
				} else {
					vertices[count] = v2
					count++
				}
			}
			node := &node{v2, nil, &w}
			addTail(v1, node)
			edges = append(edges, newEdge(v1, v2, &w))
		}
	}
	return &UDG{vertices, edges}
}

func containsVertex(v interface{}, array []*udgVertex) (bool, *udgVertex) {
	for _, tmp := range array {
		if tmp == nil {
			return false, nil
		}
		if *(tmp.value) == v {
			return true, tmp
		}
	}
	return false, nil
}

func equalsVertex(v1, v2 graph.Vertex) bool {
	return *(v1.Value()) == *(v2.Value())
}

func addTail(head graph.Vertex, node *node) {
	if head.Value() == node.v.Value() {
		return
	}
	dgHead := head.(*udgVertex)
	if dgHead.node == nil {
		dgHead.node = node
	} else {
		n := dgHead.node
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

func readline(text string) (string, string, interface{}) {
	line := strings.Split(text, ",")
	weight, _ := strconv.ParseFloat(line[2], 64)
	return line[0], line[1], weight
}

func format(head graph.Vertex) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("vertex: %s   ", *head.Value()))
	p := head.(*udgVertex).node
	for p != nil {
		b.WriteString(fmt.Sprintf("%s[%f] ", *p.v.Value(), *p.weight))
		p = p.next
	}
	return b.String()
}
