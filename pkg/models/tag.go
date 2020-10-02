package models

type Tag struct {
	ID       int
	Articles []*Article
}
