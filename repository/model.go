package repository

// A RepositoryInfo represents data about an initialized Git repository
type RepositoryInfo struct {
	Path            string `json:"path"`          // local path where the repository is located
	Url             string `json:"url"`           // Remote URL
	CurrentBranch   string `json:"branch"`        // current branch
	CurrentRevision string `json:"revision"`      // current revision
	IsUpToDate      bool   `json:"is_up_to_date"` // True if there is no uncommitted changes in the local repository
}
