package main

import "hash"

type TreeOptions struct {
	PreFilled         bool
	ZeroValue         Content
	MakeVariableDepth bool
}

type Content interface {
	CalculateHash() (hash.Hash, error)
	Equals(other Content) (bool, error)
}

type Node struct {
	leaf        bool
	right       Node
	left        Node
	parent      Node
	Data        hash.Hash
	NodeContent Content
}
type Tree struct {
	Nodes    []Node
	Depth    int
	Root     []byte
	HashFunc func() hash.Hash
}

func NewTree(c []Content, options TreeOptions) (*Tree, error) {
	return &Tree{}, nil
}
