package models

type Article struct {
	ID      int
	Title   string
	Content string
	Tags    []*Tag
}
