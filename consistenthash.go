package main

import (
	_ "crypto/md5"
	_ "fmt"
	_ "sort"
	_ "strconv"
	_ "sync"
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

	nodes []*Node

	// Все известные системе хеши ключей хранящиеся в порядке возрастания значений.
	// Значения находятся в пределе допустимых значений хэш функции hashFunc
	keys []uint64

	// Кольцо хеширования
	// key - хеш ключа
	// value - узлы, содержащие данные сектора
	ring map[uint64]*Node
}

var _ IStorage = (*ConsistentHashStorage)(nil)

// NewConsistentHashStorage - конструктор ConsistentHashStorage
func NewConsistentHashStorage(hashFunc HashFunc, nodes []*Node) *ConsistentHashStorage {
	return &ConsistentHashStorage{
		hashFunc: hashFunc,
		ring:     make(map[uint64]*Node),
		nodes:    nodes,
	}
}

func (s *ConsistentHashStorage) Get(key string) any {
	// keyHash := s.hashFunc([]byte(key))

}

func (s *ConsistentHashStorage) Put(key string, value any) {
	//TODO implement me
	panic("implement me")
}
