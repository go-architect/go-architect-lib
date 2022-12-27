package examples

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/instability"
	"github.com/fdaines/go-architect-lib/project"
)

func get_instability_and_abstractness() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	spm, err := instability.GetInstability(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Instability & Abstractness Analysis")
	fmt.Printf("%+v\n", spm)
}
