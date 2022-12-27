package examples

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/dsm"
	"github.com/fdaines/go-architect-lib/project"
)

func load_dsm() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	matrix, err := dsm.GetDependencyStructureMatrix(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Dependency Structure Matrix")
	fmt.Printf("%+v\n", matrix)
}
