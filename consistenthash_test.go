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

func TestConsistentHash(t *testing.T) {
	var nodes []*Node

	// создание нод
	for i := 0; i < hostsAmount; i++ {
		nodes[i] = NewNode(strconv.Itoa(i))
	}

	ch := NewConsistentHashStorage(MurMurHash, nodes)

	// заполнение кластера данными
	for i := 0; i < datasetSize; i++ {
		itemId := strconv.Itoa(i)
		ch.Put(itemId, GetRandomString(10, charset))
	}

	var totalSavedRecordsCount int

	// считаем сохраненные записи на узлах
	for _, node := range nodes {
		recordsCount := len(node.storage)
		totalSavedRecordsCount += recordsCount
		fmt.Printf("Amount of data stored on node named %s is %d\n", node.name, recordsCount)
	}

	assert.Equal(t, totalSavedRecordsCount, datasetSize)
}

func GetRandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
