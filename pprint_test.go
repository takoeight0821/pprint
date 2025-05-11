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
		commaSep(),
		longCommaSep(),
		longCommaFillSep(),
		commaCat(),
		longCommaFillCat(),
		fillBreak(),
		fill(),
	}

	for _, test := range tests {
		test := test // capture range variable
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var got strings.Builder
			pprint.FputDoc(&got, test.doc)
			pprint.FputDoc(&got, pprint.Text("\n"))

			var want strings.Builder
			for _, s := range test.wantLines {
				want.WriteString(s)
				want.WriteString("\n")
			}

			if diff := cmp.Diff(strings.Split(want.String(), "\n"), strings.Split(got.String(), "\n")); diff != "" {
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

func commaSep() test {
	return test{
		name: "Comma Sep",
		doc: pprint.Sep(pprint.Punctuate(pprint.Text(","),
			pprint.Text("a"),
			pprint.Text("b"),
			pprint.Text("c"),
		)...),
		wantLines: []string{
			"a, b, c",
		},
	}
}

func longCommaSep() test {
	return test{
		name: "Long Comma Sep",
		doc: pprint.Sep(pprint.Punctuate(pprint.Text(","),
			pprint.Text("item0"),
			pprint.Text("item1"),
			pprint.Text("item2"),
			pprint.Text("item3"),
			pprint.Text("item4"),
			pprint.Text("item5"),
			pprint.Text("item6"),
			pprint.Text("item7"),
			pprint.Text("item8"),
			pprint.Text("item9"),
			pprint.Text("item10"),
			pprint.Text("item11"),
			pprint.Text("item12"),
			pprint.Text("item13"),
			pprint.Text("item14"),
			pprint.Text("item15"),
			pprint.Text("item16"),
			pprint.Text("item17"),
			pprint.Text("item18"),
			pprint.Text("item19"),
		)...),
		wantLines: []string{
			"item0,",
			"item1,",
			"item2,",
			"item3,",
			"item4,",
			"item5,",
			"item6,",
			"item7,",
			"item8,",
			"item9,",
			"item10,",
			"item11,",
			"item12,",
			"item13,",
			"item14,",
			"item15,",
			"item16,",
			"item17,",
			"item18,",
			"item19",
		},
	}
}

func longCommaFillSep() test {
	return test{
		name: "Long Comma FillSep",
		doc: pprint.FillSep(pprint.Punctuate(pprint.Text(","),
			pprint.Text("item0"),
			pprint.Text("item1"),
			pprint.Text("item2"),
			pprint.Text("item3"),
			pprint.Text("item4"),
			pprint.Text("item5"),
			pprint.Text("item6"),
			pprint.Text("item7"),
			pprint.Text("item8"),
			pprint.Text("item9"),
			pprint.Text("item10"),
			pprint.Text("item11"),
			pprint.Text("item12"),
			pprint.Text("item13"),
			pprint.Text("item14"),
			pprint.Text("item15"),
			pprint.Text("item16"),
			pprint.Text("item17"),
			pprint.Text("item18"),
			pprint.Text("item19"),
		)...),
		wantLines: []string{
			"item0, item1, item2, item3, item4, item5, item6, item7, item8, item9, item10,",
			"item11, item12, item13, item14, item15, item16, item17, item18, item19",
		},
	}
}

func commaCat() test {
	return test{
		name: "Comma Cat",
		doc: pprint.Cat(pprint.Punctuate(pprint.Text(","),
			pprint.Text("item0"),
			pprint.Text("item1"),
			pprint.Text("item2"),
		)...),
		wantLines: []string{
			"item0,item1,item2",
		},
	}
}

func longCommaFillCat() test {
	return test{
		name: "Long Comma FillCat",
		doc: pprint.FillCat(pprint.Punctuate(pprint.Text(","),
			pprint.Text("item0"),
			pprint.Text("item1"),
			pprint.Text("item2"),
			pprint.Text("item3"),
			pprint.Text("item4"),
			pprint.Text("item5"),
			pprint.Text("item6"),
			pprint.Text("item7"),
			pprint.Text("item8"),
			pprint.Text("item9"),
			pprint.Text("item10"),
			pprint.Text("item11"),
			pprint.Text("item12"),
			pprint.Text("item13"),
			pprint.Text("item14"),
			pprint.Text("item15"),
			pprint.Text("item16"),
			pprint.Text("item17"),
			pprint.Text("item18"),
			pprint.Text("item19"),
		)...),
		wantLines: []string{
			"item0,item1,item2,item3,item4,item5,item6,item7,item8,item9,item10,item11,",
			"item12,item13,item14,item15,item16,item17,item18,item19",
		},
	}
}

func fillBreak() test {
	return test{
		name: "Fill Break",
		doc: pprint.Vsep(
			pprint.Hsep(pprint.FillBreak(6, pprint.Text("punctuate")), pprint.Text("="), pprint.Text("...")),
			pprint.Hsep(pprint.FillBreak(6, pprint.Text("sep")), pprint.Text("="), pprint.Text("...")),
			pprint.Hsep(pprint.FillBreak(6, pprint.Text("fillsep")), pprint.Text("="), pprint.Text("...")),
		),
		wantLines: []string{
			"punctuate",
			"       = ...",
			"sep    = ...",
			"fillsep",
			"       = ...",
		},
	}
}

func fill() test {
	return test{
		name: "Fill",
		doc: pprint.Vsep(
			pprint.Hsep(pprint.Fill(6, pprint.Text("punctuate")), pprint.Text("="), pprint.Text("...")),
			pprint.Hsep(pprint.Fill(6, pprint.Text("sep")), pprint.Text("="), pprint.Text("...")),
			pprint.Hsep(pprint.Fill(6, pprint.Text("fillsep")), pprint.Text("="), pprint.Text("...")),
		),
		wantLines: []string{
			"punctuate = ...",
			"sep    = ...",
			"fillsep = ...",
		},
	}
}
