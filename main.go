package main

import (
	"fmt"

	foobar "github.com/snamber/module-test"
)

func main() {
	result := foobar.Add(1, 2)
	fmt.Println(result)
}
