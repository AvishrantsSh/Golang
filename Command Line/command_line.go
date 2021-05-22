// This program illustrates the use of command line arguments within Go
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // The args are accessible using os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
