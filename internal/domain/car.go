package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type Car struct {
	ID        uuid.UUID
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
	OwnerID   uuid.UUID
	Mark      string
	Model     string
	RegNum    string
	Year      int
}
