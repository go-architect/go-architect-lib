package coupling

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

func calculateCouplingDetails(fileset *token.FileSet, astFile *ast.File, dep string) []Detail {
	var dependencyPrefix string
	var details []Detail
	ast.Inspect(astFile, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch t := n.(type) {
		case *ast.ImportSpec:
			if sameDependency(t.Path.Value, dep) {
				if t.Name != nil && t.Name.Name != "." {
					dependencyPrefix = t.Name.Name
				} else {
					dependencyPrefix = retrievePrefix(dep)
				}
				details = append(details, Detail{
					Line:     fileset.Position(n.Pos()).Line,
					Comments: "Imports package",
				})
			}
		case *ast.SelectorExpr:
			var buf bytes.Buffer
			printer.Fprint(&buf, fileset, t)
			if strings.HasPrefix(buf.String(), dependencyPrefix) {
				details = append(details, Detail{
					Line:     fileset.Position(t.Pos()).Line,
					Comments: fmt.Sprintf("Call to '%s'", buf.String()),
				})
			}
		}
		return true
	})
	return details
}

// This function is not used, but I kept it to traverse the ast.File structure
func astAnalysis(fileset *token.FileSet, astFile *ast.File) {
	var debug bool
	var dependencyPrefix string
	var variablesFromDependency []string
	ast.Inspect(astFile, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch t := n.(type) {
		case *ast.CallExpr:
			if debug {
				fmt.Printf("\tCallExpr: %+v\n", t)
				fmt.Printf("\t\tCallExpr1: %+v\n", t.Fun)
				fmt.Printf("\t\tCallExpr2: %+v\n", t.Args)
				/*
					if fn, ok := t.Fun.(*ast.SelectorExpr); ok {
						var buf bytes.Buffer
						printer.Fprint(&buf, fileset, fn)
						if strings.HasPrefix(buf.String(), dependencyPrefix) {
							//					fmt.Printf("\t\tFunctionX: %+v\n", buf.String())
							details = append(details, Detail{
								Line:     fileset.Position(fn.Pos()).Line,
								Comments: fmt.Sprintf("Calls function '%s'", buf.String()),
							})
						}
					}
				*/
			}
		case *ast.AssignStmt:
			if debug {
				fmt.Printf("\tAssignStmt: %+v\n", t)
				var left []string
				for _, l := range t.Lhs {
					var buf bytes.Buffer
					printer.Fprint(&buf, fileset, l)
					left = append(left, buf.String())
				}
				for i, r := range t.Rhs {
					var buf bytes.Buffer
					printer.Fprint(&buf, fileset, r)
					fmt.Printf("Right: %+v\n", buf.String())
					if strings.HasPrefix(buf.String(), dependencyPrefix) {
						variablesFromDependency = append(variablesFromDependency, left[i])
					}
				}
				fmt.Printf("\t\tVariablesFromDependency: %+v\n", variablesFromDependency)
				fmt.Printf("\t\tAssignStmtLine: %+v\n", fileset.Position(t.Pos()).Line)
				fmt.Printf("\t\tAssignStmtLeft: %+v\n", t.Lhs)
				fmt.Printf("\t\tAssignStmtRight: %+v\n", t.Rhs)
				fmt.Printf("\t\tAssignStmtOp: %+v\n", t.Tok)
			}
		case *ast.BasicLit:
			if debug {
				fmt.Printf("\tBasicLit: %+v\n", t)
			}
		case *ast.FuncLit:
			if debug {
				fmt.Printf("\tFuncLit: %+v\n", t)
			}
		case *ast.Ident:
			if debug {
				fmt.Printf("\tIdent: %+v\n", t)
				if t.Obj != nil {
					fmt.Printf("\t\tIdent2: %+v\n", t.Obj)
					fmt.Printf("\t\tIdent3: %+v\n", t.Obj.Decl)
				}
				fmt.Printf("\t\tIdentX: %+v\n", t.Name)
			}
		case *ast.BadExpr:
			if debug {
				fmt.Printf("\tBadExpr: %+v\n", t)
			}
		case *ast.Ellipsis:
			if debug {
				fmt.Printf("\tEllipsis: %+v\n", t)
			}
		case *ast.CompositeLit:
			if debug || true {
				fmt.Printf("\tCompositeLit: %+v\n", t)
			}
		case *ast.ParenExpr:
			if debug {
				fmt.Printf("\tParenExpr: %+v\n", t)
			}
		case *ast.IndexExpr:
			if debug {
				fmt.Printf("\tIndexExpr: %+v\n", t)
			}
		case *ast.IndexListExpr:
			if debug {
				fmt.Printf("\tIndexListExpr: %+v\n", t)
			}
		case *ast.SliceExpr:
			if debug {
				fmt.Printf("\tSliceExpr: %+v\n", t)
			}
		case *ast.TypeAssertExpr:
			if debug {
				fmt.Printf("\tTypeAssertExpr: %+v\n", t)
			}
		case *ast.StarExpr:
			if debug {
				fmt.Printf("\tStarExpr: %+v\n", t)
			}
		case *ast.UnaryExpr:
			if debug {
				fmt.Printf("\tUnaryExpr: %+v\n", t)
			}
		case *ast.BinaryExpr:
			if debug {
				fmt.Printf("\tBinaryExpr: %+v\n", t)
			}
		case *ast.ArrayType:
			if debug {
				fmt.Printf("\tArrayType: %+v\n", t)
			}
		case *ast.StructType:
			if debug {
				fmt.Printf("\tStructType: %+v\n", t)
			}
		case *ast.FuncType:
			if debug {
				fmt.Printf("\tFuncType: %+v\n", t)
			}
		case *ast.InterfaceType:
			if debug {
				fmt.Printf("\tInterfaceType: %+v\n", t)
			}
		case *ast.MapType:
			if debug {
				fmt.Printf("\tMapType: %+v\n", t)
			}
		case *ast.ChanType:
			if debug {
				fmt.Printf("\tChanType: %+v\n", t)
			}
		case *ast.ValueSpec:
			if debug {
				fmt.Printf("\tValueSpec: %+v\n", t)
			}
		case *ast.TypeSpec:
			if debug {
				fmt.Printf("\tTypeSpec: %+v\n", t)
			}
		case *ast.Comment:
			if debug {
				fmt.Printf("\tComment: %+v\n", t)
			}
		case *ast.CommentGroup:
			if debug {
				fmt.Printf("\tCommentGroup: %+v\n", t)
			}
		case *ast.Field:
			if debug {
				fmt.Printf("\tField: %+v\n", t)
			}
		case *ast.FieldList:
			if debug {
				fmt.Printf("\tFieldList: %+v\n", t)
			}
		case *ast.BadStmt:
			if debug {
				fmt.Printf("\tBadStmt: %+v\n", t)
			}
		case *ast.DeclStmt:
			if debug {
				fmt.Printf("\tDeclStmt: %+v\n", t)
			}
		case *ast.EmptyStmt:
			if debug {
				fmt.Printf("\tEmptyStmt: %+v\n", t)
			}
		case *ast.LabeledStmt:
			if debug {
				fmt.Printf("\tLabeledStmt: %+v\n", t)
			}
		case *ast.ExprStmt:
			if debug {
				fmt.Printf("\tExprStmt: %+v\n", t)
				fmt.Printf("\t\tExprStmt1: %+v\n", t.X)
			}
		case *ast.SendStmt:
			if debug {
				fmt.Printf("\tSendStmt: %+v\n", t)
			}
		case *ast.IncDecStmt:
			if debug {
				fmt.Printf("\tIncDecStmt: %+v\n", t)
			}
		case *ast.DeferStmt:
			if debug {
				fmt.Printf("\tDeferStmt: %+v\n", t)
			}
		case *ast.GoStmt:
			if debug {
				fmt.Printf("\tGoStmt: %+v\n", t)
			}
		case *ast.ReturnStmt:
			if debug {
				fmt.Printf("\tReturnStmt: %+v\n", t)
			}
		case *ast.BranchStmt:
			if debug {
				fmt.Printf("\tBranchStmt: %+v\n", t)
			}
		case *ast.BlockStmt:
			if debug {
				fmt.Printf("\tBlockStmt: %+v\n", t)
				fmt.Printf("\t\tBlockStmt: %+v\n", t.List)
				for _, b := range t.List {
					fmt.Printf("\t\t\tBlockStmtX: %+v\n", b)
				}
			}
		case *ast.IfStmt:
			if debug {

				fmt.Printf("\tIfStmt: %+v\n", t)
			}
		case *ast.CaseClause:
			if debug {

				fmt.Printf("\tCaseClause: %+v\n", t)
			}
		case *ast.SwitchStmt:
			if debug {

				fmt.Printf("\tSwitchStmt: %+v\n", t)
			}
		case *ast.TypeSwitchStmt:
			if debug {

				fmt.Printf("\tTypeSwitchStmt: %+v\n", t)
			}
		case *ast.CommClause:
			if debug {

				fmt.Printf("\tCommClause: %+v\n", t)
			}
		case *ast.SelectStmt:
			if debug {

				fmt.Printf("\tSelectStmt: %+v\n", t)
			}
		case *ast.ForStmt:
			if debug {
				fmt.Printf("\tForStmt: %+v\n", t)
			}
		case *ast.RangeStmt:
			if debug {
				fmt.Printf("\tRangeStmt: %+v\n", t)
			}
		case *ast.BadDecl:
			if debug {
				fmt.Printf("\tBadDecl: %+v\n", t)
			}
		case *ast.GenDecl:
			if debug {
				fmt.Printf("\tGenDecl: %+v\n", t)
			}
		case *ast.FuncDecl:
			if debug {
				fmt.Printf("\tFuncDecl: %+v\n", t)
			}
		case *ast.File:
			if debug {
				fmt.Printf("\tFile: %+v\n", t)
			}
		case *ast.Package:
			if debug {
				fmt.Printf("\tPackage: %+v\n", t)
			}
		default:
			if debug && t != nil {
				fmt.Printf("\tUnknown: %+v\n", t)
			}
		}
		return true
	})
}
