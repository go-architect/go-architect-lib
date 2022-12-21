// Package project provides functions to load information about Go projects
package project

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"os"
	"strings"
)

// LoadProjectInfo loads information for a project in the provided `path`.
// It returns an error if the path doesn't exist or the path doesn't contain a valid Golang project.
func LoadProjectInfo(path string) (*ProjectInfo, error) {
	sanitizedPath := strings.TrimPrefix(path, "file://")
	mainPackage, err := getMainPackage(sanitizedPath)
	if err != nil {
		return nil, err
	}
	name := resolveProjectName(mainPackage)

	return NewProjectInfo(name, sanitizedPath, mainPackage), nil
}

// IsValidGoProject checks if a ProjectInfo instance contains a valid Golang project.
func IsValidGoProject(project ProjectInfo) bool {
	_, err := getMainPackage(project.Path)
	if err != nil {
		return false
	}
	return true
}

func getMainPackage(path string) (string, error) {
	goModFile := path + "/go.mod"

	if _, err := os.Stat(goModFile); err == nil {
		content, _ := os.ReadFile(goModFile)
		modulePath := modfile.ModulePath(content)
		return modulePath, nil
	} else {
		return "", fmt.Errorf("Could not load %s file. %s\n", goModFile, err.Error())
	}
}

func resolveProjectName(packageName string) string {
	split := strings.Split(packageName, "/")

	return split[len(split)-1]
}
