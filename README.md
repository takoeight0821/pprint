# pprint

pprint is a port of Haskell's `wl-print` library to Go.
It provides a text rendering engine with automatic layout and formatting.

example:
```go
package main

import (
    "github.com/takoeight0821/pprint"
)

func main() {
    doc := pprint.Vsep(
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
		)

    pprint.PutDoc(doc)
}

```

output:
```
for i := range a {
  fmt.Println(i)
}
```

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/takoeight0821/pprint)