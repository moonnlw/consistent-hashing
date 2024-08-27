package main

import (
	"github.com/spaolacci/murmur3"
)

func MurMurHash(data []byte) uint64 {
	return murmur3.Sum64(data)
}
