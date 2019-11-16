package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowAccount godoc
// @Summary Create a new owner
// @Description Create a new owner
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
func (c *Router) CreateOwner(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Create Owner")
}

func (c *Router) GetOwner(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetOwner")
}

func (c *Router) DeleteOwner(ctx *gin.Context) {
	ctx.String(http.StatusMethodNotAllowed, "Not implemented")
}

func (c *Router)UpdateOwner(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UpdateOwner")
}

