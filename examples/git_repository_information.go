package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/repository"
)

func git_repository_information() {
	projectPath := "full_path_to_golang_project"

	repo, err := repository.LoadRepositoryInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("GIT Repository Information")
	fmt.Printf("%+v\n", repo)
}
