package examples

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/metrics/loc"
	"github.com/fdaines/go-architect-lib/project"
)

func count_loc() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	locInfo, err := loc.CountLoc(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Project Lines of Code")
	fmt.Printf("Total LOC: %d\n", locInfo.Total)

	fmt.Println("LOC by Package")
	for _, pkgLoc := range locInfo.Packages {
		fmt.Printf("\tPackage: %s \t\t %d\n", pkgLoc.Package, pkgLoc.LOC)
	}

	fmt.Println("LOC by File")
	for _, fileLoc := range locInfo.Files {
		fmt.Printf("\tPackage [%s] \t\tFile [%s]  \t\t %d\n", fileLoc.Package, fileLoc.File, fileLoc.LOC)
	}
}
