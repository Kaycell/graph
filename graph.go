package graph

// Graph represents an undirected graph.
// The zero value in an empty graph ready to use
type Graph struct {
	nbVertices int
	nbEdges    int
	adj        [][]int
}

// New returns an initialized graph with n vertices
func New(n int) *Graph {
	g := &Graph{}
	g.nbVertices = n
	g.nbEdges = 0
	g.adj = make([][]int, n)
	return g
}

// AddEdge links vertices v and w in both direction
func (g *Graph) AddEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.nbEdges++
}
