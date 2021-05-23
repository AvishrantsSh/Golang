package main

import "C"
import (
	"github.com/avishrantssh/Package-Test/convertor"
)

//export CtoF
func CtoF(num int) int {
	return int(convertor.CtoF(convertor.Celsius(num)))
}
func main() {
}
