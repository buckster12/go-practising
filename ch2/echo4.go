package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "skip line break") // n is *bool
var sep = flag.String("s", " ", "separator")     // sep is *string !!!

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
