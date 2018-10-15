package test

import (
	"bufio"
	"math"
	"os"
	"testing"

	"github.com/FrankLeeC/DataStructureAndAlgorithm/graph/dg"
)

func TestGraph(t *testing.T) {
	file, err := os.Open("./g.txt")
	if err != nil {
		t.Log(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g := dg.NewDG(scanner, 4, 6)
	t.Log(g.String())
	t.Log(math.MaxFloat64 < math.Inf(1))
}

func TestDijkstra(t *testing.T) {
	file, err := os.Open("./g2.txt")
	if err != nil {
		t.Log(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g := dg.NewDG(scanner, 5, 5)
	path, w := g.ShortestPath("1", "5")
	for _, v := range path {
		t.Log(*v.Value())
	}
	t.Log("==========")
	for _, v := range w {
		t.Log(v)
	}
}
