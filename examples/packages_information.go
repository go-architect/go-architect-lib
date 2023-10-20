package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
)

func get_packages_information() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	pkgs, err := packages.GetBasicPackagesInfo(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("General Packages Information")
	for idx, pkg := range pkgs {
		fmt.Printf("Packages[%d] -> %+v\n", idx, pkg)
	}
}
