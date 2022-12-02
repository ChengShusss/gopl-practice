package main

import (
	"fmt"
	"strings"
	"testing"
)

const (
	ARRAY_LEN = 1000
)

var (
	array []string
)

func init() {
	array = []string{}
	for i := 0; i < ARRAY_LEN; i++ {
		array = append(array, fmt.Sprint(i))
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(array, ",")
	}
}

func BenchmarkCustomedJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CustomedJoin(array, ",")
	}
}
