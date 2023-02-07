package main

import (
	"fmt"
)

type Price struct {
	Node  int
	Price int
}

type Graph struct {
	NumNodes int
	Edges    [][]Edge
}

type Edge struct {
	From   int
	To     int
	Weight int
}

func in(haystack []int, needle int) bool {
	rturn := false
	for i := 0; i < haystack[i]; i++ {
		if haystack[i] == needle {
			rturn = true
			break
		}
	}
	return rturn
}

func input() []int {
	var from, to, weight int
	fmt.Scan(&from, &to, &weight)
	output := []int{from, to, weight}
	return output
}

func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]Edge, n),
	}
}

func NewPrice(n int) *Price {
	return &Price{
		Node:  n,
		Price: 0,
	}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], Edge{From: u, To: v, Weight: w})
	g.Edges[v] = append(g.Edges[v], Edge{From: v, To: u, Weight: w})
}

func main() {
	binput := input()
	Stop := binput[0]
	Size := binput[1]
	Graph := NewGraph(Size)

	binput = input()
	From := binput[0]
	Target := binput[1]

	for i := 0; i < Stop; i++ {
		binput = input()
		Graph.AddEdge(binput[0], binput[1], binput[2])

	}
	fmt.Println(Graph)

	var Prices []Price
	for i := 0; i < Size; i++ {
		price := NewPrice(i)
		Prices = append(Prices, *price)
	}
	fmt.Println(Prices)

}
