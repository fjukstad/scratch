package main

import (
	"fmt"

	"github.com/fjukstad/scratch"
)

func main() {
	id := "86536890"
	p, err := scratch.GetProject(id)

	fmt.Println(p, err)

	tag := "kodeklubbentromso"

	projects, err := scratch.GetProjects(tag)

	for _, p := range projects {
		fmt.Println(p)
	}

	fmt.Println(projects, err)
}
