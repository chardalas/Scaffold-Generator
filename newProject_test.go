package main

import (
	"errors"
	"testing"
)

func TestNewProject(t *testing.T) {

	type testProject struct {
		project                Project
		args                   []string
		err                    error
		expectedOutputContains string
	}

	testCases := []testProject{
		testProject{
			args: []string{"-n", "MyProject", "-d", "/path/to/dir", "-r", "github.com/username/myproject"},
			err:  nil,
			project: Project{
				Name:     "MyProject",
				Location: "/path/to/dir",
				URL:      "github.com/username/myproject",
				Type:     false,
			},
		},
		testProject{
			args:    []string{"foo"},
			err:     errors.New("No positional parameters expected"),
			project: Project{},
		},
		testProject{
			args:                   []string{"-h"},
			err:                    errors.New("flag: help requested"),
			project:                Project{},
			expectedOutputContains: "Usage of scaffold-gen:",
		},
	}

	for _, tc := range testCases {

		p, err := tc.project.newProject(tc.args)

		if tc.err == nil && err != nil {
			t.Errorf("Expected non-nil error, got: %v", err)
		}

		if tc.err != nil {
			if err == nil || err.Error() != tc.err.Error() {
				//t.Errorf("Expected error: %v, Got: %v", tc.err, err)
			}
		}

		if p != tc.project {
			t.Errorf("Expected:%#v Got: %#v", p, tc.project)
		}
	}
}
