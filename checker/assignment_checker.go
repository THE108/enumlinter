package checker

import (
	"fmt"
	"go/ast"

	lg "github.com/THE108/enumlinter/logger"

	"golang.org/x/tools/go/loader"
)

type AssignmentChecker struct {
	Pkg         *loader.PackageInfo
	PosHelper   IPositionHelper
	Logger      lg.ILogger
	TypeToCheck string
}

func (ac *AssignmentChecker) Run(node ast.Node) (proceed bool) {
	//fmt.Printf("Run: %+v: %T\n", node, node)

	switch typ := node.(type) {
	case *ast.AssignStmt:
		proceed = ac.checkAssignStmt(typ)
	case *ast.SendStmt:
		proceed = ac.checkSendStmt(typ)
	case ast.Expr:
		if !ac.isTypeToCheck(typ) {
			return true
		}
		proceed = ac.checkExpr(typ)
	default:
		proceed = true
	}
	return
}

func (ac *AssignmentChecker) isTypeToCheck(expr ast.Expr) bool {
	return getExprType(ac.Pkg, expr) == ac.TypeToCheck
}

func (ac *AssignmentChecker) checkSendStmt(sendStmt *ast.SendStmt) bool {
	//fmt.Printf("\nsendStmt: %+v: %T\n", sendStmt, sendStmt)
	//fmt.Printf("type => Chan: %s Value: %s\n",
	//	getExprType(ac.Pkg, sendStmt.Chan), getExprType(ac.Pkg, sendStmt.Value))
	if !ac.isTypeToCheck(sendStmt.Chan) {
		return true
	}

	return ac.checkExpr(sendStmt.Value)
}

func (ac *AssignmentChecker) checkAssignStmt(asst *ast.AssignStmt) bool {
	if len(asst.Lhs) == 0 || len(asst.Rhs) == 0 {
		return true
	}

	if !ac.isTypeToCheck(asst.Lhs[0]) {
		return true
	}

	expr := asst.Rhs[0]
	if _, ok := expr.(*ast.Ident); ok {
		return true
	}

	return ac.checkExpr(expr)
}

func (ac *AssignmentChecker) reportError(expr ast.Expr, format string) {
	exprType := getExprType(ac.Pkg, expr)
	exprString, err := ac.PosHelper.GetContentByPosition(expr)
	if err != nil {
		exprString = fmt.Sprintf("can't get file content: %s", err.Error())
	}

	ac.Logger.Errorf(ac.PosHelper.GetStringPosition(expr.Pos())+": "+format, exprType, exprString)
}

func (ac *AssignmentChecker) checkExpr(expr ast.Expr) bool {
	if expr == nil {
		return false
	}

	switch typ := expr.(type) {
	case *ast.BasicLit:
		ac.reportError(expr, "%s type constant must be used instead of a basic literal %s")
		return false
	case *ast.CallExpr:
		//fmt.Printf("\nCallExpr: %+v: %T\n", typ, typ)
		if ac.isTypeToCheck(typ.Fun) {
			ac.reportError(expr, "To%s() func must be used instead of a type casting %s")
			return false
		}
	case *ast.CompositeLit:
		for _, elmt := range typ.Elts {
			if kve, ok := elmt.(*ast.KeyValueExpr); ok {
				ac.checkExpr(kve.Value)
			}
		}
		return false
	}

	return true
}
