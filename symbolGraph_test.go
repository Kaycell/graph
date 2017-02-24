package graph

import (
	"strings"
	"testing"
)

func TestNewSymbolGraphFromReader(t *testing.T) {
	sep := " "
	input := "JFK MCO\nORD DEN\nORD HOU\nDFW PHX\nJFK ATL\n"
	r := strings.NewReader(input)

	sg := NewSymbolGraphFromReader(r, sep)
	index := sg.Index("JFK")
	name := sg.Name(index)
	if name != "JFK" {
		t.Error("Expected JFK got", name)
	}
}

func TestSgString(t *testing.T) {
	sep := " "
	input := "JFK MCO\nJFK ATL\n"
	r := strings.NewReader(input)

	sg := NewSymbolGraphFromReader(r, sep)

	expectedResult := "3 vertices, 2 edges\nJFK: MCO,ATL,\nMCO: JFK,\nATL: JFK,\n"
	res := sg.String()
	if res != expectedResult {
		t.Error("Expected", expectedResult, "got", res)
	}
}
