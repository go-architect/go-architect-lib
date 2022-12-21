package comments

// CommentsMetrics represents the comments stats in ao Golang project
type CommentsMetrics struct {
	TotalLines             int     `json:"total_lines"`               // total of comment lines
	CommentsRatio          float64 `json:"ratio"`                     // ratio of comment lines (comments_loc / project_loc)
	FilesWithComments      int     `json:"files_with_comments"`       // number of files with comments
	FilesWithCommentsRatio float64 `json:"files_with_comments_ratio"` // ratio of files with comments
}
