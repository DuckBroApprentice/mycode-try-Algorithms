package classfour

import "fmt"

//先嘗試實現graph

//node { val []int, adjMatrix [][]bool}
//每個節點都維護自己指出去的邊

type graph struct {
	nodeNum   int
	adjMatrix [][]bool
}

func NewGraph(n int) *graph {
	matrix := make([][]bool, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}
	return &graph{
		nodeNum:   n,
		adjMatrix: matrix,
	}
}

func (g *graph) AddEdge(start, to int) {
	if start >= 0 && to >= 0 && start < g.nodeNum && to < g.nodeNum && start != to {
		g.adjMatrix[start][to] = true
	}
	fmt.Printf("wrong input start >= 0 : %t ; to >= 0 : %t ; start < nodeNum : %t ; to < nodeNum : %t ", start >= 0, to >= 0, start < g.nodeNum, to < g.nodeNum)
}

//start != to 不能指著自己
//  0  1 2 3 4 5
//0    t             //0指向1
//1
//2
//3
//4
//5 t                //5指向0
func (g *graph) HasEdge(start, to int) bool {
	if start > 0 && to > 0 && start < g.nodeNum && to < g.nodeNum && start != to {
		return g.adjMatrix[start][to]
	}
	return false
}
