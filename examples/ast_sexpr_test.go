package examples

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takoeight0821/pprint"
)

func TestSExpr_Output(t *testing.T) {
	expr := NewAdd(NewNumber(1), NewAdd(NewNumber(2), NewNumber(3)))
	doc := SExpr(expr)
	var buf strings.Builder
	pprint.FputDoc(&buf, doc)
	got := buf.String()
	want := "(+ 1 (+ 2 3))"
	if !cmp.Equal(got, want) {
		t.Errorf("SExpr() mismatch (-got +want):\n%s", cmp.Diff(got, want))
	}
}
