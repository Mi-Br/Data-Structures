package main

import "fmt"

type Graph struct {
	numberOfNodes int
	adjacentList  map[int][]int
}

func (g *Graph) construct() {
	g.adjacentList = make(map[int][]int)
}

func (g *Graph) addVertex(v int) {
	g.adjacentList[v] = []int{}
	g.numberOfNodes++
}
func (g *Graph) addEdge(a, b int) {
	g.adjacentList[a] = append(g.adjacentList[a], b)
	g.adjacentList[b] = append(g.adjacentList[b], a)
}
func (g *Graph) showConnections() {
	for k, v := range g.adjacentList {
		fmt.Print(k, "---->")
		for _, n := range v {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}
}

func main() {
	var G Graph
	G.construct()
	G.addVertex(0)
	G.addVertex(1)
	G.addVertex(2)
	G.addVertex(3)
	G.addVertex(4)
	G.addVertex(5)
	G.addVertex(6)
	G.showConnections()
	G.addEdge(3, 1)
	G.addEdge(3, 4)
	G.addEdge(4, 2)
	G.addEdge(4, 5)
	G.addEdge(1, 2)
	G.addEdge(1, 0)
	G.addEdge(0, 2)
	G.addEdge(6, 5)
	G.showConnections()
}
