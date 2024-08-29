package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	hostsAmount = 20
	datasetSize = 10_000_000
)

func TestConsistentHashStorage(t *testing.T) {
	nodes := make([]*Node, hostsAmount)

	// создание нод
	for i := 0; i < hostsAmount; i++ {
		nodes[i] = NewNode(strconv.Itoa(i))
	}

	ch := NewConsistentHashStorage(MurMurHash, nodes)

	fmt.Printf("Saving %d records to %d nodes\n", datasetSize, hostsAmount)
	// заполнение кластера данными
	for i := 0; i < datasetSize; i++ {
		itemId := strconv.Itoa(i)
		ch.Put(itemId, GetRandomString(10, charset))
	}

	assert.Equal(t, ch.countTotalSaved(), datasetSize)
}

func GetRandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (s *ConsistentHashStorage) countTotalSaved() int {
	var totalSavedRecordsCount int

	// считаем сохраненные записи на узлах
	for _, node := range s.nodes {
		recordsCount := len(node.storage)
		totalSavedRecordsCount += recordsCount
		fmt.Printf("Node: %s. Records: %d\n", node.name, recordsCount)
	}
	fmt.Printf("Total saved records count: %d\n", totalSavedRecordsCount)
	return totalSavedRecordsCount
}
