package main

import (
	"fmt"

	"github.com/fjukstad/scratch"
)

func main() {
	id := "86536890"
	p, err := scratch.GetProject(id)

	fmt.Println(p, err)
}
