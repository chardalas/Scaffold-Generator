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

	project.Location  = flag.String("d", "", "Project location on disk")
	project.Name  = flag.String("n", "", "Project name")
	project.URL  = flag.String("r", "", "Project remote repo URL")
	project.Type = flag.Bool("s", false, "Project will have static assets or not")

	flag.Parse()

	fmt.Fprintf(os.Stdout,"Project location: %s \n", *project.Location)
	fmt.Fprintf(os.Stdout,"Project name: %s \n", *project.Name)
	fmt.Fprintf(os.Stdout,"Project type: %t \n", *project.Type)
	fmt.Fprintf(os.Stdout,"Remote repo: %s \n", *project.URL)
	fmt.Fprintf(os.Stdout,"Generating scaffold for project %s in %s \n", *project.Name, *project.Location)


}

