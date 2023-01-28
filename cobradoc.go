// Package cobradoc implements alternative documentation generator for cobra.
package cobradoc

import (
	"bytes"
	"io"

	"github.com/spf13/cobra"
)

// Generation options
// All fields have reasonable defaults
type Options struct {
	// Page section number (defaults to "1")
	// In man page, defines manual page section number
	// In markdown, ignored
	SectionNumber string

	// Name of the tool (defaults to cmd.Name())
	// Used in numerous places
	Name string

	// Creation date (defaults to current month and year)
	// In man page, located at the center of the bottom line
	// In markdown, ignored
	Date string

	// BCP 47 language for converting the tool name to title case (defaults to "en")
	Language string

	// Page header (defaults to "{Name} Manual")
	// In man page, located at the center of the top line
	// In markdown, defines page header
	Header string

	// Page footer (defaults to "{Name} Manual")
	// In man page, located at the left corner of the bottom line
	// In markdown, ignored
	Footer string

	// Short description of the tool (defaults to cmd.Short)
	// In man page, located in NAME section
	// In markdown, located in the first section, used if LongDescription is unset
	ShortDescription string

	// Long description of the tool (defaults to cmd.Long)
	// In man page, located in DESCRIPTION section
	// In markdown, located in the first section
	LongDescription string

	// Array of additional sections (optional)
	// Sections are added to the end of the document
	ExtraSections []ExtraSection
}

// Extra section contents
type ExtraSection struct {
	// Section title
	Title string

	// Section text
	Text string
}

// Common titles for extra sections
// Section title can be any string, however these ones are used frequently
const (
	EXAMPLES    = "Examples"
	FILES       = "Files"
	ENVIRONMENT = "Environment"
	BUGS        = "Reporting Bugs"
	AUTHORS     = "Authors"
	COPYRIGHT   = "Copyright"
	LICENSE     = "License"
	HISTORY     = "History"
	SEEALSO     = "See Also"
	NOTES       = "Notes"
)

// Output format
type Format int

const (
	// Troff format (for manual page)
	Troff Format = iota

	// Markdown format
	Markdown Format = iota
)

// Generate single documentation page for given command tree
func GetDocument(cmd *cobra.Command, fmt Format, opts Options) (string, error) {
	var b bytes.Buffer

	if err := generate(cmd, fmt, opts, &b); err != nil {
		return "", err
	}

	return b.String(), nil
}

// Generate single documentation page for given command tree
func WriteDocument(w io.Writer, cmd *cobra.Command, fmt Format, opts Options) error {
	var b bytes.Buffer

	if err := generate(cmd, fmt, opts, &b); err != nil {
		return err
	}

	// send to output writer only if there are no errors
	if _, err := io.Copy(w, &b); err != nil {
		return err
	}

	return nil
}
