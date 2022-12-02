// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v", os.Args[0])

	s, sep := "", "\n"
	for d, arg := range os.Args[1:] {
		s += sep + fmt.Sprintf("%02d %v", d, arg)
	}
	fmt.Println(s)
}
