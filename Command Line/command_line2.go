// This program illustrates the use of command line arguments using iterator notation
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg // The args are accessible using os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
