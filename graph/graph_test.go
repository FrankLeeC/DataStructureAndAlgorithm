package test

import (
	"bufio"
	"os"
	"testing"
)

func TestGraph(t *testing.T) {
	file, err := os.Open("./g.txt")
	if err != nil {
		t.Log(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g := NewDirectGraph(scanner, 4, 6)
	t.Log(g.String())
}
