package graph

import "fmt"

type GraphList struct {
	Directed bool
	Weight   bool
	Nodes    map[int]bool
	Edges    map[int][]Edge
}

func (g *GraphList) AddNode(node int) error {
	if _, exists := g.Nodes[node]; exists {
		return fmt.Errorf("%v is exists", node)
	}
	g.Nodes[node] = true
	return nil
}

func (g *GraphList) AddEdge(from, to int, weight int) error {
	if !g.Weight && weight != 0 {
		return fmt.Errorf("this graph is a unweight graph")
	}
	if _, exists := g.Nodes[from]; !exists {
		return fmt.Errorf("edge is from empty node")
	}
	if _, exists := g.Nodes[to]; !exists {
		return fmt.Errorf("edge is to empty node")
	}
	g.Edges[from] = append(g.Edges[from], Edge{To: to, Weight: weight})
	//無向圖要自動建立返回的邊
	//這樣就能精簡的表達有向及無向 牛逼
	//之前的想法太執著於pointer來表現 這會很難設計結構體
	if !g.Directed {
		g.Edges[to] = append(g.Edges[to], Edge{To: from, Weight: weight})
	}
	return nil
}

//判斷節點是否存在
func (g *GraphList) HasNode(node int) bool {
	_, exists := g.Nodes[node]
	return exists
}

//判斷存不存在這條邊
func (g *GraphList) HasEdge(from, to int) bool {
	//Edge[from] for range
	//map: [5] : {{To:1},{To:2}}
	for _, end := range g.Edges[from] {
		if end.To == to {
			return true
		}
	}
	return false
}

func (g *GraphList) IsDirected() bool {
	return g.Directed
}

func (g *GraphList) IsWeight() bool {
	return g.Weight
}

//計算節點總數
func (g *GraphList) NodeCount() int {
	return len(g.Nodes)
}

//計算邊總數
func (g *GraphList) EdgeCount() int {
	count := 0
	for _, node_edges := range g.Edges {
		count += len(node_edges)
	}
	if !g.Directed {
		count /= 2
	}
	return count
}

//返回所有節點
func (g *GraphList) GetNodes() []int {
	res := make([]int, 0)
	for key, _ := range g.Edges {
		res = append(res, key)
	}
	return res
}

//返回指定節點所有邊
func (g *GraphList) GetEdges(node int) ([]Edge, error) {
	if _, exists := g.Nodes[node]; !exists {
		return nil, fmt.Errorf("node is not exists")
	}
	return g.Edges[node], nil
}
