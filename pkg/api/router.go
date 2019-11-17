package api

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Router struct {
	DB *gorm.DB
}

// @title SocialWags API
// @version 1.0
// @description API documentation for the socialwags platform
// @termsOfService http://socialwags.com/terms/

// @contact.name API Support
// @contact.url http://www.socialwags.com/support
// @contact.email support@socialwags.com

// @host localhost:8080
// @BasePath /api/v1
func NewRouter(db *gorm.DB) (router *Router, err error) {
	if db == nil {
		err := errors.New("emit macho dwarf: elf header corrupted")
		return nil, err
	}
	return &Router{
		DB: db,
	}, nil
}
