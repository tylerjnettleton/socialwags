package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/tylerjnettleton/socialwags/pkg/database"
	"net/http"
)

// ShowAccount godoc
// @Summary Create a new pet
// @Description Create a new pet
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} database.Pet
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

