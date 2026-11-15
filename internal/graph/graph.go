package graph

import "sort"

type Graph struct {
	adj [][]int
}

func New(n int) *Graph {
	return &Graph{adj: make([][]int, n)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) Vertices() int {
	return len(g.adj)
}

func (g *Graph) Components() [][]int {
	n := g.Vertices()
	visited := make([]bool, n)
	components := [][]int{}

	for v := 0; v < n; v++ {
		if visited[v] {
			continue
		}

		component := []int{}
		g.dfs(v, visited, &component)
		sort.Ints(component)
		components = append(components, component)
	}

	return components
}

func (g *Graph) IsConnected() bool {
	if g.Vertices() == 0 {
		return true
	}
	return len(g.Components()) == 1
}

func (g *Graph) dfs(v int, visited []bool, component *[]int) {
	visited[v] = true
	*component = append(*component, v)

	for _, u := range g.adj[v] {
		if !visited[u] {
			g.dfs(u, visited, component)
		}
	}
}
