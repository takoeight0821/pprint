package pprint_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takoeight0821/pprint"
)

func TestPprint(t *testing.T) {
	t.Parallel()

	tests := []test{
		helloworld(),
		forStatement(),
	}

	for _, test := range tests {
		test := test // capture range variable
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var got strings.Builder
			pprint.FputDoc(&got, test.doc)
			pprint.FputDoc(&got, pprint.Line())

			var want strings.Builder
			for _, s := range test.wantLines {
				want.WriteString(s)
				want.WriteString("\n")
			}

			if diff := cmp.Diff(want.String(), got.String()); diff != "" {
				t.Errorf("output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

type test struct {
	name      string
	doc       pprint.Doc
	wantLines []string
}

func helloworld() test {
	return test{
		name: "Hello World",
		doc:  pprint.Text("Hello World"),
		wantLines: []string{
			"Hello World",
		},
	}
}

func forStatement() test {
	return test{
		name: "For Statement",
		doc: pprint.Vsep(
			pprint.Hsep(
				pprint.Text("for"),
				pprint.Text("i"),
				pprint.Text(":="),
				pprint.Text("range"),
				pprint.Text("a"),
				pprint.Text("{"),
			),
			pprint.Indent(2, pprint.Hcat(
				pprint.Text("fmt"),
				pprint.Text("."),
				pprint.Text("Println"),
				pprint.Text("("),
				pprint.Text("i"),
				pprint.Text(")"),
			)),
			pprint.Text("}"),
		),
		wantLines: []string{
			"for i := range a {",
			"  fmt.Println(i)",
			"}",
		},
	}
}
