package main

type Node struct {
	value    int
	children []Node //array of nodes
}

var s int

func sum(n Node) int {
	if n.value == 0 {
		return 0
	}
	for i := 0; i < len(n.children); i++ {
		s = sum(n.children[i])
	}

	return s + n.value
}

// PFS
