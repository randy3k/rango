package prompt

type Content struct {
	Lines []Chars
	Cursor DocumentCursor
}


func NewContent(lines []Chars, cursor DocumentCursor) *Content {
	return &Content{
		Lines: lines,
		Cursor: cursor,
	}
}
