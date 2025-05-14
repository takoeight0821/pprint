package examples

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takoeight0821/pprint"
)

func TestJSONToMarkdown_Output(t *testing.T) {
	cfg := Config{"localhost", 8080, true}
	doc, err := JSONToMarkdown(cfg)
	if err != nil {
		t.Fatalf("JSONToMarkdown() error = %v", err)
	}
	var buf strings.Builder
	pprint.FputDoc(&buf, doc)
	got := buf.String()
	wants := []string{
		"| Key   | Value     |",
		"|-------|-----------|",
		"| debug | true      |",
		"| host  | localhost |",
		"| port  | 8080      |",
	}
	want := strings.Join(wants, "\n")
	if !cmp.Equal(got, want) {
		t.Errorf("JSONToMarkdown() mismatch (-got +want):\n%s", cmp.Diff(got, want))
	}
}
