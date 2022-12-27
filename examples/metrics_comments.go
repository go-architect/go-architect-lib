package examples

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/metrics/comments"
	"github.com/fdaines/go-architect-lib/project"
)

func comments_metrics() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	commentMetrics, err := comments.ResolveCommentsMetrics(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Comments Metrics")
	fmt.Printf("Total Lines of Comments: %v\n", commentMetrics.TotalLines)
	fmt.Printf("Total Files with Comments: %v\n", commentMetrics.FilesWithComments)

	fmt.Printf("Ratio Lines of Comments: %v\n", commentMetrics.CommentsRatio)
	fmt.Printf("Ratio Files with Comments: %v\n", commentMetrics.FilesWithCommentsRatio)
}
