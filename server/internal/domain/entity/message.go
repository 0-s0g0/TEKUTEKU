package entity

import (
	"time"

	"github.com/0-s0g0/TEKUTEKU/server/pkg/null"
)

type Message struct {
	ID        string
	Message   string
	Likes     int
	X         int
	Y         int
	FloatTime float32
	School    int
	CreatedAt time.Time
	ParentID  null.Null[string]
	Reply     []Message
}
