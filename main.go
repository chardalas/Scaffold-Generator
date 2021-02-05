package main

import (
	"flag"
	"fmt"
	"os"
)
type Project struct {
	Name     *string
	Location *string
	URL      *string
	Type     *bool
}

func NewProject() *Project {
	return &Project{}
}

func main() {

	project := NewProject()

	projectCommand := flag.NewFlagSet("", flag.ExitOnError)

	project.Location  = projectCommand.String("d", "", "Project location on disk")
	project.Name  = projectCommand.String("n", "", "Project name")
	project.URL  = projectCommand.String("r", "", "Project remote repo URL")
	project.Type = projectCommand.Bool("s", false, "Project will have static assets or not")

	projectCommand.Parse(os.Args[1:])

	if projectCommand.Parsed() {
		if *project.Location == "" ||  *project.Name == "" || *project.URL == ""{
			projectCommand.PrintDefaults()
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stdout,"Generating scaffold for project %s in %s \n", *project.Name, *project.Location)
}

