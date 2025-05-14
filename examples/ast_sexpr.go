// Example: Print S-expression of a simple AST
package examples

import (
	"fmt"

	"github.com/takoeight0821/pprint"
)

type Expr any

type Number struct {
	Value int
}

type Add struct {
	Left, Right Expr
}

func SExpr(e Expr) pprint.Doc {
	switch v := e.(type) {
	case Number:
		return pprint.Text(fmt.Sprintf("%d", v.Value))
	case Add:
		return pprint.Sep(
			pprint.Beside(pprint.Text("("), pprint.Text("+")),
			SExpr(v.Left),
			pprint.Beside(SExpr(v.Right), pprint.Text(")")),
		)
	default:
		panic(fmt.Sprintf("unsupported expression type: %T", v))
	}
}
