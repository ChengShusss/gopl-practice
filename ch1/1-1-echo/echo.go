// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// practice in 1.1
	//====================
	fmt.Printf("%v ", os.Args[0])
	//====================

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
