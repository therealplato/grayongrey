package main

import (
	"flag"
	"fmt"
)

func main() {
	var n uint
	flag.UintVar(&n, "n", 2, "number of aliens")
	flag.Parse()
	fmt.Println(n)
}
