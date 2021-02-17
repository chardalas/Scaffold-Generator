package main

import (
	"errors"
	"flag"
)

type Project struct {
	Name     string
	Location string
	URL      string
	Type     bool
}

func (p *Project) newProject(args []string) (Project, error) {

	projectCommand := flag.NewFlagSet("", flag.ContinueOnError)

	projectCommand.StringVar(&p.Location, "d", "", "Project location on disk")
	projectCommand.StringVar(&p.Name, "n", "", "Project name")
	projectCommand.StringVar(&p.URL, "r", "", "Project remote repo URL")
	projectCommand.BoolVar(&p.Type, "s", false, "Project will have static assets or not")

	err := projectCommand.Parse(args)

	if err != nil {
		return Project{}, err
	}

	return *p, nil
}

func (p *Project) validateSettings() []error {

	var errs []error

	if p.Name == "" {
		errs = append(errs, errors.New("Project name cannot be empty"))
	}
	if p.Location == "" {
		errs = append(errs, errors.New("Project path cannot be empty"))
	}
	if p.URL == "" {
		errs = append(errs, errors.New("Project repository URL cannot be empty"))
	}

	return errs
}
