package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowAccount godoc
// @Summary Create a new pet
// @Description Create a new pet
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
func (c *Router) CreatePet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Create Pet")
}

func (c *Router) GetPet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Get Pet")
}

func (c *Router) DeletePet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Delete Pet")
}

func (c *Router)UpdatePet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Update Pet")
}

