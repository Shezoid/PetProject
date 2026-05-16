package internal

import (
	"time"
)

type Pet struct {
	id        int
	name      string
	birthDate time.Time
	breed     string
	color     string
	ownerId   int
	friendIds []int
}
