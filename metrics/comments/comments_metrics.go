// Package comments provides functions to know information about comments in a Golang project
package comments

import (
	"github.com/go-architect/go-architect-lib/internal/utils"
	"github.com/go-architect/go-architect-lib/internal/utils/loc"
	packagesUtils "github.com/go-architect/go-architect-lib/internal/utils/packages"
	"github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// ResolveCommentsMetrics retrieves metrics about comments in the provided Golang project.
//
// An error is returned when it's not possible to get the packages information
func ResolveCommentsMetrics(prj *project.ProjectInfo) (*CommentsMetrics, error) {
	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	var commentLines, lines, filesWithComments int
	var filesTotal int
	for _, pkgInfo := range pkgs {
		if pkgInfo.PackageData != nil {
			codeFiles := packagesUtils.GetCodeFiles(pkgInfo.PackageData)
			filesTotal += len(codeFiles)
			for _, f := range codeFiles {
				cl, _ := countComments(pkgInfo.PackageData.Dir, f)
				fLOC, _ := loc.CountLinesOfCode(pkgInfo.PackageData.Dir, f)
				commentLines += cl
				lines += fLOC
				if cl > 0 {
					filesWithComments++
				}
			}
		}
	}
	ratio := utils.RoundFloat(100*float64(commentLines)/float64(lines), 2)
	filesWithCommentsRatio := utils.RoundFloat(100*float64(filesWithComments)/float64(filesTotal), 2)

	return &CommentsMetrics{
		TotalLines:             commentLines,
		CommentsRatio:          ratio,
		FilesWithComments:      filesWithComments,
		FilesWithCommentsRatio: filesWithCommentsRatio,
	}, nil
}

func countComments(pkgPath string, srcFile string) (int, error) {
	data, err := os.ReadFile(filepath.Join(pkgPath, srcFile))
	if err != nil {
		return 0, err
	}
	fileset := token.NewFileSet()
	astFile, err := parser.ParseFile(fileset, srcFile, data, parser.ParseComments)

	if err != nil {
		return 0, err
	}
	var commentLines int
	for _, c := range astFile.Comments {
		for range c.List {
			commentLines++
		}
	}

	return commentLines, nil
}
