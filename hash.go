package main

import (
	"github.com/spaolacci/murmur3"
)

const MaxUint64 = ^uint64(0)

func MurMurHash(data []byte) uint64 {
	return murmur3.Sum64(data)
}
