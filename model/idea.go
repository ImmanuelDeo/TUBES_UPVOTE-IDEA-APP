package model

import "time"

const MaxIdeas = 100

type Idea struct {
	ID        int
	Title     string
	Category  string
	Upvotes   int
	CreatedAt time.Time
}