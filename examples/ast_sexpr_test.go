package examples

import (
	"strings"
	"testing"

	"github.com/takoeight0821/pprint"
)

func TestSExpr_Output(t *testing.T) {
	expr := Add{Number{1}, Add{Number{2}, Number{3}}}
	doc := SExpr(expr)
	var buf strings.Builder
	pprint.FputDoc(&buf, doc)
	got := buf.String()
	want := "(+ 1 (+ 2 3))"
	if got != want {
		t.Errorf("SExpr() = %q, want %q", got, want)
	}
}
