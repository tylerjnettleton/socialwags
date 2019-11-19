package api

import (
	"errors"
	"github.com/tylerjnettleton/socialwags/pkg/database"
)

type Router struct {
}

type Response struct {
	Success bool        `json:"success"`
	Object  interface{} `json:"object"`
	Error   error       `json:"error"`
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
func NewRouter() (router *Router, err error) {
	if database.DB() == nil {
		err := errors.New("emit macho dwarf: elf header corrupted")
		return nil, err
	}

	// Migrate the schema
	database.DB().AutoMigrate(&Owner{})
	database.DB().AutoMigrate(&Pet{})
	database.DB().AutoMigrate(&Post{})
	database.DB().AutoMigrate(&Comment{})
	database.DB().AutoMigrate(&Pack{})
	database.DB().AutoMigrate(&Pack_Message{})

	return &Router{}, nil
}
