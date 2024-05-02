/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package util

import "go/ast"

func IsIOTA(vspec *ast.ValueSpec) bool {
	for _, expr := range vspec.Values {
		bexpr, ok := expr.(*ast.BinaryExpr)
		if !ok {
			continue
		}
		switch v := bexpr.X.(type) {
		case *ast.Ident:
			if v.Name == "iota" {
				return true
			}
		}
	}
	return false
}
