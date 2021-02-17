package main

import (
	"errors"
	"testing"
)

func TestValidateSettings(t *testing.T) {
	type testProject struct {
		project Project
		errs    []error
	}

	testCases := []testProject{
		testProject{
			project: Project{
				Name:     "MyProject",
				Location: "/path/to/dir",
				URL:      "github.com/username/myproject",
				Type:     false,
			},
			errs: []error{},
		},
		testProject{
			project: Project{
				Type: false,
			},
			errs: []error{
				errors.New("Project name cannot be empty"),
				errors.New("Project path cannot be empty"),
				errors.New("Project repository URL cannot be empty"),
			},
		},
	}

	for _, tc := range testCases {
		errs := tc.project.validateSettings()

		if len(tc.errs) == 0 && len(errs) != 0 {
			t.Errorf("Expected no errors, got: %v", errs)
		}

		for i, e := range tc.errs {
			if errs[i] == nil || errs[i].Error() != e.Error() {
				t.Errorf("Expected error: %v, Got: %v", e, errs[i])
			}
		}
	}
}
