package examples

import (
	"strings"
	"testing"

	"github.com/takoeight0821/pprint"
)

func TestPrettyStruct_Output(t *testing.T) {
	p := Person{"Alice", 30}
	doc := PrettyStruct(p)
	var buf strings.Builder
	pprint.FputDoc(&buf, doc)
	got := buf.String()
	want := "Name: Alice\nAge: 30"
	if got != want {
		t.Errorf("PrettyStruct() = %q, want %q", got, want)
	}
}
