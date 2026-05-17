package infrastructure

import (
	"time"
)

type Pet struct {
	Id        int
	Name      string
	BirthDate time.Time
	Breed     string
	Color     string
	OwnerId   int
	FriendIds []int
}
