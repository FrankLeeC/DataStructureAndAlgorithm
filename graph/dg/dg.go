package dg

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/FrankLeeC/DataStructureAndAlgorithm/graph"
)

type node struct {
	v      graph.Vertex
	next   *node
	weight *interface{}
}

type dgVertex struct {
	value *interface{}
	node  *node
}

func (v *dgVertex) Value() *interface{} {
	return v.value
}

func (v *dgVertex) UniqueID() string {
	return ""
}

type dgEdge struct {
	start  graph.Vertex
	end    graph.Vertex
	weight *interface{}
}

func (e *dgEdge) Start() graph.Vertex {
	return e.start
}

func (e *dgEdge) End() graph.Vertex {
	return e.end
}

func (e *dgEdge) Weight() *interface{} {
	return e.weight
}

// DG directed graph
type DG struct {
	vertices []*dgVertex
	edges    []*dgEdge
}

// V get count of vertices
func (g *DG) V() int {
	return len(g.vertices)
}

// E get count of edges
func (g *DG) E() int {
	return len(g.edges)
}

func (g *DG) containsVertex(v interface{}) (bool, graph.Vertex) {
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
func (g *DG) AddEdge(e graph.Edge) {
	start := e.Start()
	node := &node{e.End(), nil, e.Weight()}
	if c, v := g.containsVertex(*(start.Value())); c {
		start = v
	}
	addTail(start, node)
}

// Adj adjancies of vertex
func (g *DG) Adj(v graph.Vertex) []graph.Vertex {
	var idx int
	for i, v1 := range g.vertices {
		if equalsVertex(v, v1) {
			idx = i
			break
		}
	}
	v2 := g.vertices[idx]
	rs := make([]graph.Vertex, 0)
	if v2.node == nil {
		return nil
	}
	n := v2.node
	for n != nil {
		rs = append(rs, n.v)
		n = n.next
	}
	return rs
}

func (g *DG) Weight(s, t graph.Vertex) *interface{} {
	for _, edge := range g.edges {
		if equalsVertex(edge.Start(), s) && equalsVertex(edge.End(), t) {
			return edge.Weight()
		}
	}
	return nil
}

func (g *DG) index(v graph.Vertex) int {
	for i, t := range g.vertices {
		if equalsVertex(v, t) {
			return i
		}
	}
	return -1
}

// ShortestPath shortest path from s to t
func (g *DG) ShortestPath(sval, tval interface{}) ([]graph.Vertex, []interface{}) {
	s := newVertex(sval)
	t := newVertex(tval)
	q := make([]graph.Vertex, len(g.vertices), len(g.vertices))
	exclude := make(map[int]int)
	dst := make([]float64, len(g.vertices), len(g.vertices))
	path := make([][]graph.Vertex, len(g.vertices), len(g.vertices))
	var c int
	for i, v := range g.vertices {
		if equalsVertex(v, s) {
			c = i
		}
		dst[i] = math.Inf(1)
		q[i] = v
	}
	qsize := len(q)
	dst[c] = 0
	sv := make([]graph.Vertex, 1, 1)
	sv[0] = s
	path[c] = sv
	var tidx int
	for qsize > 0 {
		idx, minValue := findMinVertex(dst, exclude)
		exclude[idx] = 1
		qsize--
		u := q[idx]
		if equalsVertex(t, u) {
			tidx = idx
			break
		}
		for _, neighbor := range g.Adj(u) {
			dst2 := minValue + (*g.Weight(u, neighbor)).(float64)
			neighborIdx := g.index(neighbor)
			if dst2 < dst[neighborIdx] {
				dst[neighborIdx] = dst2
				p := append(path[idx], neighbor)
				path[neighborIdx] = p
			}
		}
	}
	w := make([]interface{}, 0, len(path[tidx])-1)
	for i := 0; i < len(path[tidx])-1; i++ {
		a := path[tidx][i]
		b := path[tidx][i+1]
		w = append(w, *g.Weight(a, b))
	}
	return path[tidx], w
}

func findMinVertex(values []float64, exclude map[int]int) (int, float64) {
	min := math.Inf(1)
	mv := math.Inf(1)
	for i, v := range values {
		if v < mv {
			if _, c := exclude[i]; !c {
				min = float64(i)
				mv = float64(v)
			}
		}
	}
	return int(min), mv
}

// String string
func (g *DG) String() string {
	var b strings.Builder
	for _, v := range g.vertices {
		b.WriteString(format(v) + "\n")
	}
	return b.String()
}

func newVertex(value interface{}) *dgVertex {
	return &dgVertex{&value, nil}
}

func newEdge(v1, v2 graph.Vertex, weight *interface{}) *dgEdge {
	return &dgEdge{v1, v2, weight}
}

// NewDG new direct graph
func NewDG(scanner *bufio.Scanner, vcount, ecount int) *DG {
	vertices := make([]*dgVertex, vcount, vcount)
	edges := make([]*dgEdge, 0, ecount)
	count := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			a, b, w := readline(scanner.Text())
			var v1, v2 *dgVertex
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
	return &DG{vertices, edges}
}

func containsVertex(v interface{}, array []*dgVertex) (bool, *dgVertex) {
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
	dgHead := head.(*dgVertex)
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
	p := head.(*dgVertex).node
	for p != nil {
		b.WriteString(fmt.Sprintf("%s[%f] ", *p.v.Value(), *p.weight))
		p = p.next
	}
	return b.String()
}
