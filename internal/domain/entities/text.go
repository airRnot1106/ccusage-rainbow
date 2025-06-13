package entities

// Text represents a text to be displayed as ASCII art
type Text struct {
	Content string
}

// NewText creates a new Text entity
func NewText(content string) *Text {
	return &Text{
		Content: content,
	}
}

// IsEmpty returns true if the text content is empty
func (t *Text) IsEmpty() bool {
	return t.Content == ""
}

// Length returns the length of the text content
func (t *Text) Length() int {
	return len(t.Content)
}
