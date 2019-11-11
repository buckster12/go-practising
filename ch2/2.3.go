package main

import (
	"fmt"
)

func main2() {
	var i, j, k int
	// var b, f, s = true, 2.3, "four"
	// b=true
	// f=2.3
	// s="four"

	fmt.Printf("%d %d %d", i, j, k)
	// fmt.Printf("%s %f %s", b, f, s)
	fmt.Println("")

	// name := "test.txt"
	// var file, err = os.Open(name) // Open returns file and error (almost all the functions return error/nil as a last parameter)

	// 2.3.2
	x := 1
	p := &x // p contains something like 0xc000086010
	// fmt.Println(*p)
	*p = 555
	fmt.Println(p)  // p = 0xc000086010
	fmt.Println(*p) // p = 2
	fmt.Println(x)  // x = 2

	v := 1
	fmt.Printf("v=%d\n", v) // v = 2

	incr(&v)
	fmt.Printf("*p=%d\n", *p) // p = 555 - old value

	fmt.Printf("v=%d\n", v) // v = 2
}

func incr(p *int) int {
	*p++ // we must use * because int is different to *int
	fmt.Printf("p=%q\n", *p)
	return *p // we must
}
