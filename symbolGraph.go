package graph

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// SymbolGraph represents a graph with
// symbolic vertex names
type SymbolGraph struct {
	Graph
	st   map[string]int
	keys []string
}

// NewSymbolGraphFromReader create a SymbolGraph from a Reader
// Ex :
// airport.txt
// JFK MCO
// ORD DEN
// ORD HOU
// DFW PHX
// JFK ATL
// ORD DFW
// ORD PHX
// ATL HOU
func NewSymbolGraphFromReader(r io.Reader, sep string) *SymbolGraph {
	sg := &SymbolGraph{}
	sg.st = make(map[string]int)
	scanner := bufio.NewScanner(r)
	edges := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		symbols := strings.Split(line, sep)
		if !sg.Contains(symbols[0]) {
			sg.st[symbols[0]] = len(sg.st)
			sg.keys = append(sg.keys, symbols[0])
		}
		for i := 1; i < len(symbols); i++ {
			if !sg.Contains(symbols[i]) {
				sg.st[symbols[i]] = len(sg.st)
				sg.keys = append(sg.keys, symbols[i])
				edges[symbols[0]] = append(edges[symbols[0]], symbols[i])
			}
		}
	}
	sg.nbVertices = len(sg.st)
	sg.adj = make([][]int, sg.nbVertices)
	for source, targets := range edges {
		for _, target := range targets {
			sg.AddEdge(sg.st[source], sg.st[target])
		}
	}
	return sg
}

// Contains returns whether sg contains vertex with key symbolic name
func (sg *SymbolGraph) Contains(key string) bool {
	_, ok := sg.st[key]
	return ok
}

// Index returns index used internally for key vertex
func (sg *SymbolGraph) Index(key string) int {
	return sg.st[key]
}

// Name returns v symbolic name
func (sg *SymbolGraph) Name(v int) string {
	return sg.keys[v]
}

// String returns a string representation of the graph's adjacency lists
func (sg SymbolGraph) String() string {
	s := fmt.Sprintln(sg.nbVertices, "vertices,", sg.nbEdges, "edges")
	for v := 0; v < sg.nbVertices; v++ {
		s += sg.Name(v) + ": "
		for _, w := range sg.adj[v] {
			s += sg.Name(w) + ","
		}
		s += fmt.Sprintln()
	}
	return s
}
