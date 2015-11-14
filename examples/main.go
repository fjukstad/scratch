package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fjukstad/scratch"
)

func main() {
	id := "86536890"
	p, err := scratch.GetProject(id)

	fmt.Println(p, err)

	tags := []string{"kodeklubbentromso", "tromso"}

	projects := []*scratch.Project{}
	for _, tag := range tags {
		ps, err := scratch.GetProjects(tag)
		if err != nil {
			fmt.Println(err)
			return
		}
		projects = append(projects, ps...)
	}

	filename := "ids"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, p := range projects {
		fmt.Println(p)
		_, err = f.WriteString(strconv.Itoa(p.ID) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(projects, err)
}
