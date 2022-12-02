package main

import (
	"fmt"
	"strings"
	"testing"
)

var (
	array []string
)

func init() {
	array = []string{}
	for i := 0; i < 200; i++ {
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
