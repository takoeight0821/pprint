package examples

import (
	"strings"
	"testing"

	"github.com/takoeight0821/pprint"
)

func TestJSONToMarkdown_Output(t *testing.T) {
	cfg := Config{"localhost", 8080, true}
	doc := JSONToMarkdown(cfg)
	var buf strings.Builder
	pprint.FputDoc(&buf, doc)
	got := buf.String()
	wants := []string{
		"| Key   | Value     |",
		"|-------|-----------|",
		"| host  | localhost |",
		"| port  | 8080      |",
		"| debug | true      |",
	}
	want := strings.Join(wants, "\n")
	if got != want {
		t.Errorf("JSONToMarkdown() = %q, want %q", got, want)
	}
}
