package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Post struct {
	gorm.Model
	OwnerID     uint
	Post_Type   uint
	Picture_URL string
	Title       string
	Body        string
	Time        time.Time
	TimeZone    string
	Comments    []Comment
}

func (c *Router) CreatePost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Create Post")
}

func (c *Router) GetPost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetPost")
}

func (c *Router) DeletePost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DeletePost")
}

func (c *Router) UpdatePost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UpdatePost")
}
