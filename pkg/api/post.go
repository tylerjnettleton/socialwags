package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowAccount godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
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

