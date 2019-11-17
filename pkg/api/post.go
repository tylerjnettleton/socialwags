package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Router) CreatePost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Create Post")
}

func (c *Router) GetPost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetPost")
}

func (c *Router) DeletePost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DeletePost")
}

func (c *Router)UpdatePost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UpdatePost")
}

