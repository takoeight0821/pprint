// Package pprint is a port of the Haskell library `wl-pprint` to Go.
// It provides a pretty-printing library for rendering documents in a human-readable format.
package pprint

import (
	"fmt"
	"io"
	"math"
	"os"
)

// Combinators

// Punctuate concatenates a list of documents with a separator.
func Punctuate(p Doc, docs ...Doc) []Doc {
	if len(docs) == 0 {
		return nil
	}

	result := make([]Doc, 0, len(docs))
	for _, d := range docs {
		result = append(result, Beside(d, p))
	}
	result[len(docs)-1] = docs[len(docs)-1]

	return result
}

// Sep concatenates documents either horizonally, if it fits the page, or veritically, if it doesn't.
func Sep(docs ...Doc) Doc {
	return Group(Vsep(docs...))
}

// FillSep concatenates documents horizontally as long as its fits the page, than, insert a `Line` and continues doing that for all documents.
func FillSep(docs ...Doc) Doc {
	if len(docs) == 0 {
		return Empty()
	}

	result := docs[0]
	for _, d := range docs[1:] {
		result = Beside(result, Beside(SoftLine(), d))
	}

	return result
}

// Hsep concatenates documents horizontally.
func Hsep(docs ...Doc) Doc {
	if len(docs) == 0 {
		return Empty()
	}

	result := docs[0]
	for _, d := range docs[1:] {
		result = Beside(result, (Beside(Char(' '), d)))
	}

	return result
}

// Vsep concatenates documents vertically.
func Vsep(docs ...Doc) Doc {
	if len(docs) == 0 {
		return Empty()
	}

	result := docs[0]
	for _, d := range docs[1:] {
		result = Beside(result, (Beside(Line(), d)))
	}

	return result
}

// Cat concatenates documents like `Sep`, but with `LineBreak` instead of `Line`.
func Cat(docs ...Doc) Doc {
	return Group(Vcat(docs...))
}

// FillCat concatenates documents like `FillSep`, but with `LineBreak` instead of `Line`.
func FillCat(docs ...Doc) Doc {
	if len(docs) == 0 {
		return Empty()
	}

	result := docs[0]
	for _, d := range docs[1:] {
		result = Beside(result, Beside(SoftBreak(), d))
	}

	return result
}

// Hcat concatenates documents without any space between them.
func Hcat(docs ...Doc) Doc {
	if len(docs) == 0 {
		return Empty()
	}

	result := docs[0]
	for _, d := range docs[1:] {
		result = Beside(result, d)
	}

	return result
}

// Vcat concatenates documents vertically like `Vsep`, but with `LineBreak` instead of `Line`.
func Vcat(docs ...Doc) Doc {
	if len(docs) == 0 {
		return Empty()
	}

	result := docs[0]
	for _, d := range docs[1:] {
		result = Beside(result, (Beside(LineBreak(), d)))
	}

	return result
}

// SoftLine is a line break that is rendered as a space if the group fits on the page, or as a line break otherwise.
func SoftLine() Doc {
	return Group(Line())
}

// SoftBreak is a line break that is rendered as nothing if the group fits on the page, or as a line break otherwise.
func SoftBreak() Doc {
	return Group(LineBreak())
}

// FillBreak inserts a line break and indents if the document exceeds the given width, otherwise pads with spaces.
func FillBreak(f int, doc Doc) Doc {
	return width(doc, func(w int) Doc {
		if w > f {
			return Nest(f, LineBreak())
		}
		return Text(Spaces(f - w))
	})
}

// Fill pads the document with spaces if its width is less than the specified value f.
func Fill(f int, doc Doc) Doc {
	return width(doc, func(w int) Doc {
		if w >= f {
			return Empty()
		}
		return Text(Spaces(f - w))
	})
}

// width takes a document and a function that uses the width of that document to create another document.
// It returns a new document where the second document is placed beside the first document.
// The function f is called with the width (in columns) of the first document.
// This is useful for aligning text or creating document layouts where the second part depends on the width of the first part.
//
// Parameters:
//   - doc: The first document
//   - f: A function that takes the width of the first document and returns a second document
//
// Returns:
//   - A new document composed of the first document beside the result of f
func width(doc Doc, f func(int) Doc) Doc {
	return Column(func(k1 int) Doc {
		return Beside(
			doc,
			Column(func(k2 int) Doc {
				return f(k2 - k1)
			}),
		)
	})
}

// Indent indents the document by the specified amount.
func Indent(i int, doc Doc) Doc {
	return Hang(i, Beside(Text(Spaces(i)), doc))
}

// Hang implements hanging indentation.
//
//	Hang(4, FillSep([]Doc{Text("the"), Text("hang"), Text("combinator"), Text("indents"), Text("these"), Text("words"), Text("!")}))
//
// will render as:
//
//	the hang combinator
//	    indents these
//	    words!
func Hang(i int, doc Doc) Doc {
	return Align(Nest(i, doc))
}

// Align renders the document with the nesting level set to the current column.
//
//	Sep([]Doc{Text("hi"), Align(Vsep([]Doc{Text("nice"), Text("world")}))})
//
// will render as:
//
//	hi nice
//	   world
func Align(doc Doc) Doc {
	return Column(func(k int) Doc {
		return Nesting(func(i int) Doc {
			return Nest(k-i, doc)
		})
	})
}

// Spaces returns a string consisting of n space characters.
func Spaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

// Primitives

// Doc represents a pretty-printable document. Implementations of Doc are used to build and render formatted output.
type Doc interface {
	doc()
}

type empty struct{}

func (empty) doc() {}

var _ Doc = empty{}

type char rune

func (char) doc() {}

var _ Doc = char(0)

type text string

func (text) doc() {}

var _ Doc = text("")

type line struct {
	IsBreak bool
}

func (line) doc() {}

var _ Doc = line{}

type cat struct {
	First  Doc
	Second Doc
}

func (cat) doc() {}

var _ Doc = cat{}

type nest struct {
	Indent int
	Doc    Doc
}

func (nest) doc() {}

var _ Doc = nest{}

type union struct {
	Longer  Doc
	Shorter Doc
}

func (union) doc() {}

var _ Doc = union{}

type column func(int) Doc

func (column) doc() {}

var _ Doc = column(func(int) Doc { return empty{} })

type nesting func(int) Doc

func (nesting) doc() {}

var _ Doc = nesting(func(int) Doc { return empty{} })

// Empty has no content.
func Empty() Doc {
	return empty{}
}

// Char represents a single character.
func Char(c rune) Doc {
	if c == '\n' {
		return Line()
	}

	return char(c)
}

// Text represents a string of text.
func Text(s string) Doc {
	if s == "" {
		return Empty()
	}

	return text(s)
}

// Line advances to the next line and indents to the current indentation level. `Line()` behaves like `Char(' ')` if the line break is undone by `Group`.
func Line() Doc {
	return line{IsBreak: false}
}

// LineBreak advances to the next line and indents to the current indentation level. `LineBreak()` behaves like `Empty()` if the line break is undone by `Group`.
func LineBreak() Doc {
	return line{IsBreak: true}
}

// Beside concatenates two documents horizontally.
func Beside(first, second Doc) Doc {
	return cat{First: first, Second: second}
}

// Nest renders the document with the current indentation level increased by the specified amount.
//
//	Vsep(Nest(2, Vsep(Text("hello"), Text("world"))), Text("!"))
//
// will render as:
//
//	hello
//	  world
//	!
func Nest(indent int, doc Doc) Doc {
	return nest{Indent: indent, Doc: doc}
}

// Column creates a document that depends on the current column position.
func Column(f func(int) Doc) Doc {
	return column(f)
}

// Nesting creates a document that depends on the current nesting (indentation) level.
func Nesting(f func(int) Doc) Doc {
	return nesting(f)
}

// Group undoes all line breaks in the document.
func Group(doc Doc) Doc {
	return union{
		Longer:  flatten(doc),
		Shorter: doc,
	}
}

// flatten removes all line breaks from the document.
func flatten(doc Doc) Doc {
	switch d := doc.(type) {
	case cat:
		return cat{
			First:  flatten(d.First),
			Second: flatten(d.Second),
		}
	case char:
		return d
	case column:
		return column(func(width int) Doc {
			return flatten(d(width))
		})
	case empty:
		return d
	case line:
		if d.IsBreak {
			return Empty()
		}
		return Text(" ")
	case nest:
		return nest{
			Indent: d.Indent,
			Doc:    flatten(d.Doc),
		}
	case nesting:
		return nesting(func(width int) Doc {
			return flatten(d(width))
		})
	case text:
		return d
	case union:
		return flatten(d.Longer)
	default:
		panic(fmt.Sprintf("unexpected pprint.Doc: %#v", d))
	}
}

// SimpleDoc represents a rendered document in a simplified form for output.
type SimpleDoc interface {
	simpleDoc()
}

// SEmpty represents an empty SimpleDoc node.
type SEmpty struct{}

func (SEmpty) simpleDoc() {}

var _ SimpleDoc = SEmpty{}

// SChar represents a single character in a SimpleDoc.
type SChar struct {
	char rune
	rest SimpleDoc
}

func (SChar) simpleDoc() {}

var _ SimpleDoc = SChar{}

// SText represents a string of text in a SimpleDoc.
type SText struct {
	text string
	rest SimpleDoc
}

func (SText) simpleDoc() {}

var _ SimpleDoc = SText{}

// SLine represents a line break and indentation in a SimpleDoc.
type SLine struct {
	indent int
	rest   SimpleDoc
}

func (SLine) simpleDoc() {}

var _ SimpleDoc = SLine{}

// Renderers

// Docs is a list of indentation/document pairs.
type Docs struct {
	Rest   *Docs
	Indent int
	Doc    Doc
}

// Cons constructs a Docs list node with the given indentation and document.
func Cons(indent int, doc Doc, rest *Docs) *Docs {
	return &Docs{
		Rest:   rest,
		Indent: indent,
		Doc:    doc,
	}
}

// Nil returns an empty Docs list.
func Nil() *Docs {
	return nil
}

// RenderPretty renders the document with pretty printing.
// `RenderPretty(rfrac, w, x)` will render the document `x` with a page width of `w` and a ribbon width of `rfrac * w`.
// The ribbon width is the maximal amout of non-identation characters on a line.
// `rfrac` should be between 0 and 1.
func RenderPretty(rfrac float64, w int, x Doc) SimpleDoc {
	r := max(w, int(math.Round(float64(w)*rfrac)))

	pretty := pretty{
		rfrac: rfrac,
		w:     w,
		x:     x,
		r:     r,
	}

	return pretty.best(0, 0, Cons(0, x, Nil()))
}

type pretty struct {
	rfrac float64
	w     int
	x     Doc
	r     int
}

func (p pretty) best(n int, k int, docs *Docs) SimpleDoc {
	if docs == Nil() {
		return SEmpty{}
	}

	i := docs.Indent
	d := docs.Doc
	ds := docs.Rest

	switch d := d.(type) {
	case empty:
		return p.best(n, k, ds)
	case char:
		return SChar{
			char: rune(d),
			rest: p.best(n, k+1, ds),
		}
	case text:
		return SText{
			text: string(d),
			rest: p.best(n, k+len(d), ds),
		}
	case line:
		return SLine{
			indent: i,
			rest:   p.best(i, i, ds),
		}
	case cat:
		return p.best(n, k, Cons(i, d.First, Cons(i, d.Second, ds)))
	case nest:
		return p.best(n, k, Cons(i+d.Indent, d.Doc, ds))
	case union:
		width := min(p.w-k, p.r-k+n)

		if x := p.best(n, k, Cons(i, d.Longer, ds)); fits(width, x) {
			return x
		}

		return p.best(n, k, Cons(i, d.Shorter, ds))
	case column:
		return p.best(n, k, Cons(i, d(k), ds))
	case nesting:
		return p.best(n, k, Cons(i, d(i), ds))
	default:
		panic(fmt.Sprintf("unexpected pprint.Doc: %#v", d))
	}
}

func fits(w int, x SimpleDoc) bool {
	if w < 0 {
		return false
	}

	switch x := x.(type) {
	case SEmpty:
		return true
	case SChar:
		return fits(w-1, x.rest)
	case SText:
		s := x.text
		l := len(s)
		return fits(w-l, x.rest)
	case SLine:
		return true
	default:
		panic(fmt.Sprintf("unexpected pprint.SimpleDoc: %#v", x))
	}
}

// RenderCompact renders the document without pretty-printing, producing a SimpleDoc.
func RenderCompact(x Doc) SimpleDoc {
	return scan(0, []Doc{x})
}

func scan(k int, ds []Doc) SimpleDoc {
	if len(ds) == 0 {
		return SEmpty{}
	}

	d := ds[0]
	ds = ds[1:]

	switch d := d.(type) {
	case empty:
		return scan(k, ds)
	case char:
		return SChar{
			char: rune(d),
			rest: scan(k+1, ds),
		}
	case text:
		return SText{
			text: string(d),
			rest: scan(k+len(string(d)), ds),
		}
	case line:
		return SLine{
			indent: 0,
			rest:   scan(0, ds),
		}
	case cat:
		return scan(k, append([]Doc{d.First, d.Second}, ds...))
	case nest:
		return scan(k, append([]Doc{d.Doc}, ds...))
	case union:
		return scan(k, append([]Doc{d.Shorter}, ds...))
	case column:
		return scan(k, append([]Doc{d(k)}, ds...))
	case nesting:
		return scan(k, append([]Doc{d(0)}, ds...))
	default:
		panic(fmt.Sprintf("unexpected pprint.Doc: %#v", d))
	}
}

// Display writes the rendered SimpleDoc to the given writer.
func Display(w io.Writer, x SimpleDoc) error {
	switch x := x.(type) {
	case SEmpty:
		return nil
	case SChar:
		if _, err := fmt.Fprint(w, string(x.char)); err != nil {
			return err
		}

		return Display(w, x.rest)
	case SText:
		if _, err := fmt.Fprint(w, x.text); err != nil {
			return err
		}

		return Display(w, x.rest)
	case SLine:
		if _, err := fmt.Fprintf(w, "\n%s", indentation(x.indent)); err != nil {
			return err
		}

		return Display(w, x.rest)
	default:
		panic(fmt.Sprintf("unexpected pprint.SimpleDoc: %#v", x))
	}
}

func indentation(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

// PutDoc writes the pretty-printed document to standard output.
func PutDoc(doc Doc) error {
	return FputDoc(os.Stdout, doc)
}

// FputDoc writes the pretty-printed document to the given writer.
func FputDoc(w io.Writer, doc Doc) error {
	simpleDoc := RenderPretty(0.4, 80, doc)

	return Display(w, simpleDoc)
}

// Pretty is an interface for types that can be pretty-printed as a Doc.
type Pretty interface {
	Pretty() Doc
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
