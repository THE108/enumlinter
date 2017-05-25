package checker

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/loader"
)

type IPositionHelper interface {
	GetContentByPosition(node ast.Node) (string, error)
	GetStringPosition(pos token.Pos) string
}

func getExprType(pkg *loader.PackageInfo, expr ast.Expr) string {
	if typ := pkg.TypeOf(expr); typ != nil {
		return trimPackageName(typ.String())
	}
	return "<unknown type>"
}

func trimPackageName(typeName string) string {
	index := strings.Index(typeName, ".")
	if index < 0 {
		return typeName
	}

	return typeName[index+1:]
}
