/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package app

import (
	"github.com/Totus-Floreo/goconstdoc/internal/domain"
	"github.com/Totus-Floreo/goconstdoc/pkg/util"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

const (
	goconstdocPrefix = "goconstdoc:"
	goconstdocColumn = "goconstdoc:column"
	goconstdocIgnore = "goconstdoc:ignore"
)

func ParseGoFile(fileName string) (*domain.Table, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var table domain.Table
	table.Columns = processColumns(node.Comments)

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Comment:
			if strings.Contains(x.Text, goconstdocIgnore) {
				return false
			}

		case *ast.CommentGroup:
			if strings.Contains(x.Text(), goconstdocIgnore) {
				return false
			}

		case *ast.GenDecl:
			if x.Tok == token.CONST {
				for _, spec := range x.Specs {
					vspec := spec.(*ast.ValueSpec)
					if util.IsIOTA(vspec) {
						panic("iota is not allowed")
					}
					for i, name := range vspec.Names {
						val := vspec.Values[i]
						if vspec.Comment != nil {
							row, ok := processComments(vspec.Comment.List, name, val)
							if !ok {
								continue
							}
							table.Rows = append(table.Rows, row)
						} else if vspec.Doc != nil {
							row, ok := processComments(vspec.Doc.List, name, val)
							if !ok {
								continue
							}
							table.Rows = append(table.Rows, row)
						}
					}
				}
			}
		}
		return true
	})

	return &table, nil
}

func processComments(comments []*ast.Comment, name *ast.Ident, val ast.Expr) (domain.Row, bool) {
	row := make(domain.Row)
	for _, comment := range comments {
		text := util.RemoveCommentSymbols(comment.Text)
		if strings.HasPrefix(text, goconstdocIgnore) {
			return nil, false
		}
		if strings.HasPrefix(text, goconstdocPrefix) {
			fields := strings.Split(text, ":")[1:]
			for i := 0; i < len(fields); i = i + 2 {
				row[fields[i]] = fields[i+1]
			}
			row.SetName(name.Name)
			row.SetValue(val.(*ast.BasicLit).Value)
		}
	}

	return row, true
}

func processColumns(commentGroups []*ast.CommentGroup) (columns []domain.Column) {
	for _, commentGroup := range commentGroups {
		comments := strings.Split(util.RemoveCommentSymbols(commentGroup.Text()), "\n")
		for _, comment := range comments {
			if strings.HasPrefix(comment, goconstdocColumn) {
				columns = append(columns,
					domain.Column{
						Name:  strings.Split(comment, ":")[2],
						Value: strings.Split(comment, ":")[3],
					})
			}
		}
	}

	return columns
}
