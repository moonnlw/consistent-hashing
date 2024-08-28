package main

import (
	"sort"
)

// IStorage - интерфейс хранилища данных вида ключ - значение
type IStorage interface {
	// Get возвращает хранимое значение по ключу
	Get(key string) any
	// Put сохраняет запись ключ-значение
	Put(key string, value any)
}

// HashFunc представляет хэш-функцию
type HashFunc func(data []byte) uint64

// ConsistentHashStorage реализует секционирование методом согласованного хеширования (хеш-секционирование)
// в рамках распределенной системы хранения данных вида ключ-значение
type ConsistentHashStorage struct {
	// Хеш-функция используемая для кодирования ключей
	hashFunc HashFunc

	// Ноды кластера хранящие данные
	nodes []*Node

	// Хеш значения - границы секторов
	// Отсортированы в порядке возрастания. Принимают значения допустимые хэширующей функцией hashFunc
	sectorBorders []uint64

	// Кольцо хеширования
	// key - хеш ключа
	// value - узлы, содержащие данные сектора
	ring map[uint64]*Node
}

var _ IStorage = (*ConsistentHashStorage)(nil)

// NewConsistentHashStorage - конструктор ConsistentHashStorage
func NewConsistentHashStorage(hashFunc HashFunc, nodes []*Node) *ConsistentHashStorage {
	sb := calculateSectorsHash(len(nodes))
	return &ConsistentHashStorage{
		hashFunc:      hashFunc,
		nodes:         nodes,
		sectorBorders: sb,
		ring:          matchNodesToSectors(nodes, sb),
	}
}

func (s *ConsistentHashStorage) Get(key string) any {
	keyHash := s.hashFunc([]byte(key))
	return s.findKeysNode(keyHash).Get(keyHash)
}

func (s *ConsistentHashStorage) Put(key string, value any) {
	keyHash := s.hashFunc([]byte(key))
	s.findKeysNode(keyHash).Put(keyHash, value)
}

// Просчитывает хэши для секторов данных.
// Возвращает список хэшей - границ секторов в порядке возрастания.
func calculateSectorsHash(sectorsAmount int) []uint64 {
	hashRingStep := MaxUint64 / uint64(sectorsAmount) // шаг между хэш значениями границ секторов
	sb := make([]uint64, sectorsAmount)               // хэши границ секторов

	// вычисление границ секторов
	lastHash := uint64(0)
	for i := 0; i < sectorsAmount; i++ {
		lastHash += hashRingStep
		sb[i] = lastHash
	}

	return sb
}

// Соотносит сектора с узлами их обслуживающими
func matchNodesToSectors(nodes []*Node, sectorsBorders []uint64) map[uint64]*Node {
	if len(nodes) != len(sectorsBorders) {
		panic("Virtual nodes are not supported. Amount of nodes should match sectors amount")
	}

	ring := make(map[uint64]*Node)

	for idx := range nodes {
		ring[sectorsBorders[idx]] = nodes[idx]
	}
	return ring
}

// Производит поиск ноды хранящей значение по ключу
func (s *ConsistentHashStorage) findKeysNode(keyHash uint64) *Node {
	idx := sort.Search(len(s.sectorBorders), func(i int) bool {
		return s.sectorBorders[i] >= keyHash
	})
	sector := s.sectorBorders[idx]
	return s.ring[sector]
}
