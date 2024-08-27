package main

// Node Узел системы
type Node struct {

	// storage хранилище данных
	// Значение хранится по unit64 хешу
	storage map[uint64]any
}

func (n *Node) Get(keyHash uint64) any {
	return n.storage[keyHash]
}

func (n *Node) Put(keyHash uint64, value any) {
	n.storage[keyHash] = value
}
