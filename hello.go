package main

import (
	"fmt"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("\n")
	fmt.Println(quote.Opt())
}
