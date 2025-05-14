// Example: Print S-expression of a simple AST
package examples

import (
	"fmt"

	"github.com/takoeight0821/pprint"
)

type Expr struct {
	Node interface {
		expr()
		pprint.Pretty
	}
}

func (e Expr) Pretty() pprint.Doc {
	return e.Node.Pretty()
}

//go:generate go run golang.org/x/tools/cmd/stringer@v0.32.0 -type=ExprTag
type ExprTag int

const (
	ExprTagNumber ExprTag = iota
	ExprTagAdd
)

type Number struct {
	Value int
}

func (Number) expr() {}

func (n Number) Pretty() pprint.Doc {
	return pprint.Text(fmt.Sprintf("%d", n.Value))
}

func NewNumber(value int) Expr {
	return Expr{Node: Number{Value: value}}
}

type Add struct {
	Left, Right Expr
}

func (Add) expr() {}

func (a Add) Pretty() pprint.Doc {
	return pprint.Sep(
		pprint.Beside(pprint.Text("("), pprint.Text("+")),
		a.Left.Pretty(),
		pprint.Beside(a.Right.Pretty(), pprint.Text(")")),
	)
}

func NewAdd(left, right Expr) Expr {
	return Expr{Node: Add{Left: left, Right: right}}
}

func SExpr(e Expr) pprint.Doc {
	return e.Pretty()
}
