package main

// Node Узел системы
type Node struct {
	name string

	// storage хранилище данных
	// Значение хранится по unit64 хешу
	storage map[uint64]any
}

func NewNode(name string) *Node {
	return &Node{
		name:    name,
		storage: make(map[uint64]any),
	}
}

func (n *Node) Get(keyHash uint64) any {
	return n.storage[keyHash]
}

func (n *Node) Put(keyHash uint64, value any) {
	n.storage[keyHash] = value
}
