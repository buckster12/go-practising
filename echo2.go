package main

import (
	"fmt"
	"os"
)

func main() {
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// s += sep + arg
	// sep = " "
	// }
	// fmt.Println(s)

	// fmt.Println(strings.Join(os.Args[1:], "-"))

	// fmt.Println(os.Args[:])

	for key, value := range os.Args {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println("---")
	}
}
