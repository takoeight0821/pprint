// Example: Read a JSON config and pretty-print as Markdown
package examples

import (
	"fmt"

	"github.com/takoeight0821/pprint"
)

type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Debug bool   `json:"debug"`
}

func JSONToMarkdown(cfg Config) pprint.Doc {
	lines := []pprint.Doc{
		pprint.Text("| Key   | Value     |"),
		pprint.Text("|-------|-----------|"),
		pprint.Text(fmt.Sprintf("| %-5s | %-9s |", "host", cfg.Host)),
		pprint.Text(fmt.Sprintf("| %-5s | %-9d |", "port", cfg.Port)),
		pprint.Text(fmt.Sprintf("| %-5s | %-9v |", "debug", cfg.Debug)),
	}
	return pprint.Vsep(lines...)
}
