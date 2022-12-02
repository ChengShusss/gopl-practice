package main

import (
	"fmt"
	"strings"
	"time"
)

func CustomedJoin(array []string, sep string) string {
	s := ""
	for _, ch := range array {
		s += sep + ch
	}
	return s
}

func main() {
	array := []string{}
	for i := 0; i < 200; i++ {
		array = append(array, fmt.Sprint(i))
	}

	start := time.Now()
	s := strings.Join(array, ",")

	secs := time.Since(start).Nanoseconds()
	fmt.Printf("%dns  for Join %v\n", secs, s[:10])

	start = time.Now()
	s = CustomedJoin(array, ",")

	secs = time.Since(start).Nanoseconds()
	fmt.Printf("%dns  for CustomJoin %v\n", secs, s[:10])

}
