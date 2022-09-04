package main

import (
	"fmt"
	"hex/internal/adapters/core/arithmetic"
)

func main() {
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1,2)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(result)


}