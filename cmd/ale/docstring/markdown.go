package docstring

import (
	"regexp"
	"strings"

	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/read"
)

var keyValue = regexp.MustCompile(`^([^:]+):\s*(.*)$`)

// ParseMarkdown parses the kind of markdown document that might be processed
// by a static site generator. It will parse any prologue parameters into the
// resulting object and return the remaining content as individual lines
func ParseMarkdown(doc string) (data.Object, []string) {
	obj, rest := parseKeyValues(doc)
	return obj, skipEmptyLines(rest)
}

func skipEmptyLines(lines []string) []string {
	var first = 0
	for first < len(lines) && strings.TrimSpace(lines[first]) == "" {
		first++
	}
	return lines[first:]
}

func parseKeyValues(doc string) (data.Object, []string) {
	obj := data.Object{}
	lines := strings.Split(doc, "\n")
	if strings.TrimSpace(lines[0]) != "---" {
		return obj, lines
	}
	lines = lines[1:]
	var rest = 1
	for i, l := range lines {
		if strings.TrimSpace(l) == "---" {
			rest = i + 1
			break
		}
		if k, v, ok := parseKeyValue(l); ok {
			obj[k] = v
		}
	}
	return obj, lines[rest:]
}

func parseKeyValue(l string) (n data.Name, v data.Value, ok bool) {
	defer func() {
		if rec := recover(); rec != nil {
			ok = false
		}
	}()

	if sm := keyValue.FindStringSubmatch(l); sm != nil {
		name := data.Name(sm[1])
		value := read.FromString(data.String(sm[2])).First()
		return name, value, true
	}
	return "", nil, false
}
