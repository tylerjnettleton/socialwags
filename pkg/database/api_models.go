package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// This is our database models

type Owner struct {
	gorm.Model
	Email_Address string `gorm:"unique;not null"`
	Password string
	Salt string
	First_Name string `gorm:"not null"`
	Last_Name string `gorm:"not null"`
	Picture_URL string
	Zip_Code uint `gorm:"not null"`
	Time time.Time
	TimeZone string
	Pets []Pet
	Posts []Post
	Packs []Pack
}

type Pet struct {
	OwnerID uint
	Name string
	Birth_Date string
	Picture_URL string
}

type Post struct {
	OwnerID uint
	Post_Type uint
	Picture_URL string
	Title string
	Body string
	Time time.Time
	TimeZone string
	Comments []Comment

}

type Comment struct {
	PostID uint
	OwnerID uint
	Body string
	Likes uint
	Time time.Time
	TimeZone string
}

type Pack struct {
	OwnerID uint
	Time time.Time
	TimeZone string
	RecipientIDs []Owner
	Messages []Pack_Message
}

type Pack_Message struct {
	PackID uint
	Message_Owner Owner
	Body string
	Time time.Time
	TimeZone string
}