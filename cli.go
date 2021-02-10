package main

import (
	"flag"
)

type Project struct {
	Name     string
	Location string
	URL      string
	Type     bool
}

func (p *Project) NewProject(args []string) (*Project, error) {

	projectCommand := flag.NewFlagSet("", flag.ContinueOnError)

	projectCommand.StringVar(&p.Location, "d", "", "Project location on disk")
	projectCommand.StringVar(&p.Name, "n", "", "Project name")
	projectCommand.StringVar(&p.URL, "r", "", "Project remote repo URL")
	projectCommand.BoolVar(&p.Type, "s", false, "Project will have static assets or not")

	err := projectCommand.Parse(args)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Project) validateSettings() []string {

	var errors []string

	if p.Name == "" {
		errors = append(errors, "Project name cannot be empty\n")
	}
	if p.Location == "" {
		errors = append(errors, "Project path cannot be empty\n")
	}
	if p.URL == "" {
		errors = append(errors, "Project repository URL cannot be empty\n")
	}

	return errors
}
