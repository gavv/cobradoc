package cobradoc

import (
	_ "embed"
	"errors"
	"io"
	"regexp"
	"strings"
	"text/template"
)

type formatInfo struct {
	Options

	GlobalFlags      []flagInfo
	GlobalFlagsBlock string

	Groups []groupInfo
}

type groupInfo struct {
	Title    string
	Commands []commandInfo
}

type commandInfo struct {
	Path        string
	Usage       string
	Description string
	Flags       []flagInfo
	FlagsBlock  string
}

type flagInfo struct {
	Short           string
	Long            string
	DefaultValue    string
	ValueIsOptional bool
	IsBool          bool
	Type            string
	Description     string
}

var (
	//go:embed format_troff.tmpl
	formatTroff string

	//go:embed format_markdown.tmpl
	formatMarkdown string
)

var formatTemplates = map[Format]string{
	Troff:    formatTroff,
	Markdown: formatMarkdown,
}

var formatFuncs = map[Format]template.FuncMap{
	Troff: template.FuncMap{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"escape": func(s string) string {
			s = regexp.MustCompile(`\n+\n`).ReplaceAllString(s, "\n.PP\n")
			s = strings.NewReplacer(
				"-", "\\-", //
				"_", "\\_", //
				"&", "\\&", //
				"\\", "\\\\", //
				"~", "\\~", //
			).Replace(s)
			return s
		},
	},
	Markdown: template.FuncMap{
		"anchor": func(s string) string {
			return "#" + strings.ReplaceAll(s, " ", "-")
		},
	},
}

func format(fmt Format, fmtInfo formatInfo, w io.Writer) error {
	templateText, ok := formatTemplates[fmt]
	if !ok {
		return errors.New("invalid format")
	}

	templateFuncs, ok := formatFuncs[fmt]
	if !ok {
		return errors.New("invalid format")
	}

	t, err := template.New("cobradoc").Funcs(templateFuncs).Parse(templateText)
	if err != nil {
		return err
	}

	err = t.Execute(w, fmtInfo)
	if err != nil {
		return err
	}

	return nil
}
