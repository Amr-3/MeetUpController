package _const

type Span int

const (
	Weekly    Span = 0 + iota
	Monthly
	OneTime
)