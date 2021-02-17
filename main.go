package main

import (
	"fmt"
	"os"
)

func main() {
	var p Project

	project, err := p.newProject(os.Args[1:])

	if err != nil {
		os.Exit(1)
	}

	errors := project.validateSettings()

	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Fprint(os.Stdout, e)
		}
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Generating scaffold for project %s in %s \n", project.Name, project.Location)
}
