package api

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	PostID   uint
	OwnerID  uint
	Body     string
	Likes    uint
	Time     time.Time
	TimeZone string
}
