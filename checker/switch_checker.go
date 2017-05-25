package checker

import (
	"fmt"
	"go/ast"

	lg "github.com/THE108/enumlinter/logger"

	"golang.org/x/tools/go/loader"
)

type SwitchChecker struct {
	Pkg       *loader.PackageInfo
	PosHelper IPositionHelper
	Logger    lg.ILogger
}

func (sc *SwitchChecker) Run(node ast.Node) {
	switch typ := node.(type) {
	case *ast.SwitchStmt:
		sc.checkBlockStmt(typ.Body)
	}
}

func (sc *SwitchChecker) checkBlockStmt(body *ast.BlockStmt) {
	for _, stmt := range body.List {
		sc.checkStatements(stmt)
	}
}

func (sc *SwitchChecker) checkStatements(stmt ast.Stmt) {
	caseClause, ok := stmt.(*ast.CaseClause)
	if !ok {
		return
	}

	for _, expr := range caseClause.List {
		if sc.errorIfNotIndent(expr) {
			continue
		}
	}
}

func (sc *SwitchChecker) errorIfNotIndent(expr ast.Expr) bool {
	switch expr.(type) {
	case *ast.BasicLit:
		return false
	case *ast.Ident:
		return false
	}

	exprType := getExprType(sc.Pkg, expr)
	exprString, err := sc.PosHelper.GetContentByPosition(expr)
	if err != nil {
		exprString = fmt.Sprintf("can't get file content: %s", err.Error())
	}
	sc.Logger.Errorf(sc.PosHelper.GetStringPosition(expr.Pos())+": %s type constant must be used instead %s",
		exprType, exprString)
	return true
}
