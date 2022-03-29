package usecase

type Cursor string

func CursorFromRef(c *string) *Cursor {
	if c == nil {
		return nil
	}
	d := Cursor(*c)
	return &d
}

func (c Cursor) Ref() *Cursor {
	return &c
}

func (c *Cursor) CopyRef() *Cursor {
	if c == nil {
		return nil
	}
	d := *c
	return &d
}

func (c *Cursor) StringRef() *string {
	if c == nil {
		return nil
	}
	s := string(*c)
	return &s
}

type PageInfo struct {
	TotalCount      int
	StartCursor     *Cursor
	EndCursor       *Cursor
	HasNextPage     bool
	HasPreviousPage bool
}

type Pagination struct {
	Before *Cursor
	After  *Cursor
	First  *int
	Last   *int
}

func NewPagination(first *int, last *int, before *Cursor, after *Cursor) *Pagination {
	// Relay-Style Cursor Pagination
	// ref: https://www.apollographql.com/docs/react/features/pagination/#relay-style-cursor-pagination
	return &Pagination{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}
}
