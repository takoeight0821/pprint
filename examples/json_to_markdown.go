// Example: Read a JSON config and pretty-print as Markdown
package examples

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/takoeight0821/pprint"
)

type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Debug bool   `json:"debug"`
}

func JSONToMarkdown[T any](value T) (pprint.Doc, error) {
	// Convert the struct to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}
	// Unmarshal JSON to a map
	var data map[string]any
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Get all keys
	keys := make([]string, 0, len(data))
	maxKeyLen := 0
	for k := range data {
		keys = append(keys, k)
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}
	// Sort keys
	sort.Strings(keys)

	// Get all values
	values := make([]pprint.Doc, len(data))
	maxValueLen := 0
	for i, k := range keys {
		v := data[k]
		str := fmt.Sprintf("%v", v)
		values[i] = pprint.Text(str)
		if len(str) > maxValueLen {
			maxValueLen = len(str)
		}
	}

	lines := make([]pprint.Doc, len(data)+2)
	lines[0] = pprint.Hsep(
		pprint.Text("|"),
		pprint.Fill(maxKeyLen, pprint.Text("Key")),
		pprint.Text("|"),
		pprint.Fill(maxValueLen, pprint.Text("Value")),
		pprint.Text("|"),
	)
	lines[1] = pprint.Hcat(
		pprint.Text("|"),
		pprint.Text(strings.Repeat("-", maxKeyLen+2)),
		pprint.Text("|"),
		pprint.Text(strings.Repeat("-", maxValueLen+2)),
		pprint.Text("|"),
	)

	for i, k := range keys {
		lines[i+2] = pprint.Hsep(
			pprint.Text("|"),
			pprint.Fill(maxKeyLen, pprint.Text(k)),
			pprint.Text("|"),
			pprint.Fill(maxValueLen, values[i]),
			pprint.Text("|"),
		)
	}

	return pprint.Vsep(lines...), nil
}
