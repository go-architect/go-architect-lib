package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/project"
)

func project_information_example() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Project Information")
	fmt.Printf("* Project Name: %s\n", projectInfo.Name)
	fmt.Printf("* Project Path: %s\n", projectInfo.Path)
	fmt.Printf("* Project Main Package: %s\n", projectInfo.Package)
}
