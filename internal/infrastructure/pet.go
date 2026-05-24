package infrastructure

import (
	"time"
)

type Pet struct {
	ID        int
	Name      string
	BirthDate time.Time
	Breed     string
	Color     string
	OwnerID   int
	FriendIDs []int
}
