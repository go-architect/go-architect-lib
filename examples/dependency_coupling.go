package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/coupling"
	"github.com/go-architect/go-architect-lib/project"
)

func dependency_coupling() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	cl, err := coupling.CalculateCoupling(projectInfo, "fmt")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("Dependency: %s\n", cl.Dependency)
	fmt.Printf("CouplingLevel: %v\n", cl.CouplingLevel)
	for _, p := range cl.PackageDetails {
		fmt.Printf("\tPackage: %s\n", p.Package)
		fmt.Printf("\tCouplingLevel: %v\n", p.CouplingLevel)
		for _, f := range p.FileDetails {
			fmt.Printf("\t\tFile: %s\n", f.File)
			fmt.Printf("\t\tCouplingLevel: %v\n", f.CouplingLevel)
			for _, d := range f.Details {
				fmt.Printf("\t\t\tLine: %d \tDetails: %s\n", d.Line, d.Comments)
			}
		}
	}
}
