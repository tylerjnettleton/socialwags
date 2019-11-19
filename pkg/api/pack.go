package api

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Pack struct {
	gorm.Model
	OwnerID      uint
	Time         time.Time
	TimeZone     string
	RecipientIDs []Owner
	Messages     []Pack_Message
}

type Pack_Message struct {
	gorm.Model
	PackID        uint
	Message_Owner Owner
	Body          string
	Time          time.Time
	TimeZone      string
}
