package schema

type SeverityKind string

const (
	SeverityInfo    SeverityKind = "->"
	SeverityWarning SeverityKind = "~>"
	SeverityCaution SeverityKind = "!>" // is caution the right word to use here?
)

type Note struct {
	Severity SeverityKind
	Content  string
}

// Arguably this isn't necessary and downstream tools consuming the `WebsiteNotes` field can apply their own formatting
func (n Note) String() string {
	if n.Severity == "" {
		n.Severity = SeverityInfo
	}

	return string(n.Severity) + n.Content
}
