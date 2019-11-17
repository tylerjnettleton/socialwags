package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tylerjnettleton/socialwags/pkg/database"
	"net/http"
)


// Binding for Owner requests
type CreateOwnerRequest struct {
	First_Name string `json:"first_name" binding:"required"`
	Last_Name string `json:"last_name" binding:"required"`
	Email_Address string `json:"email_address" binding:"required"`
	Password string `json:"password" binding:"required"`
	Zip_Code string `json:"zip_code" binding:"required"`
	Picture_Data string `json:"picture_data" binding:"-"`
}

// Create a new Owner account
// @Summary Create a new owner account
// @Description Create a new owner
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param Body body api.CreateOwnerRequest false "Owners First Name"
// @Success 200
// @Router /owner [put]
func (r *Router) CreateOwner(ctx *gin.Context) {

	var json CreateOwnerRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o := &database.Owner{
		First_Name: "Tyler",
		Last_Name: "Nettleton",
		Email_Address: "Tyler@tylernettleton.com",
	}

	r.DB.Create(o)

	ctx.String(http.StatusOK, "Create Owner")
}

func (r *Router) GetOwner(ctx *gin.Context) {

	o := &database.Owner{
		First_Name: "Tyler",
		Last_Name: "Nettleton",
		Email_Address: "Tyler@tylernettleton.com",
	}

	r.DB.Create(o)
	ctx.String(http.StatusOK, "GetOwner")
}

func (r *Router) DeleteOwner(ctx *gin.Context) {
	ctx.String(http.StatusMethodNotAllowed, "Not implemented")
}

func (r *Router)UpdateOwner(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UpdateOwner")
}

