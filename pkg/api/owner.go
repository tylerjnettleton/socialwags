package api

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tylerjnettleton/socialwags/pkg/crypto"
	"github.com/tylerjnettleton/socialwags/pkg/database"
	"net/http"
	"time"
)

type Owner struct {
	gorm.Model
	Email_Address string `gorm:"unique;not null"`
	Password      string
	Salt          string
	First_Name    string `gorm:"not null"`
	Last_Name     string `gorm:"not null"`
	Picture_URL   string
	Zip_Code      uint `gorm:"not null"`
	Time          time.Time
	TimeZone      string
	Pets          []Pet
	Posts         []Post
	Packs         []Pack
}

// Gin Bindings
type CreateOwnerRequest struct {
	First_Name    string `json:"first_name" binding:"required"`
	Last_Name     string `json:"last_name" binding:"required"`
	Email_Address string `json:"email_address" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Zip_Code      uint   `json:"zip_code" binding:"required"`
	Picture_Data  string `json:"picture_data" binding:"-"`
}

type GetOwnerRequest struct {
	Owner_ID uint `form:"owner_id" binding:"required"`
}

type UpdateOwnerRequest struct {
	CreateOwnerRequest
	Owner_ID uint `form:"owner_id" binding:"required"`
}

type DeleteOwnerRequest struct {
	GetOwnerRequest
}

// Create a new Owner account
// @Summary Create a new owner account
// @Description Create a new owner
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param Body body api.CreateOwnerRequest true "Create owner request body"
// @Success 200
// @Router /owner [put]
func (r *Router) CreateOwner(ctx *gin.Context) {

	var request CreateOwnerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Let's create a salt and hash the users password before we store it in the database
	salt, _ := crypto.RandBytes()
	hashedPassword, _ := crypto.HashPassword([]byte(request.Password), salt)

	o := &Owner{
		First_Name:    request.First_Name,
		Last_Name:     request.Last_Name,
		Email_Address: request.Email_Address,
		Zip_Code:      request.Zip_Code,
		Salt:          hex.EncodeToString(salt),
		Password:      hex.EncodeToString(hashedPassword),
	}

	if result := database.DB().Create(o); result.Error != nil {
		errors := result.GetErrors()
		for _, err := range errors {
			fmt.Println(err)
		}
		// Todo: Do better error handling with json response
		resJson := Response{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := Response{
		Success: true,
		Error:   nil,
	}

	ctx.JSON(http.StatusOK, resJson)
}

// Get a specific `Owner`
// @Summary Get a specific owner
// @Description Get a specific owner
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param owner_id query uint false "Owner ID"
// @Success 200
// @Router /owner [get]
func (r *Router) GetOwner(ctx *gin.Context) {
	var form GetOwnerRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	own := Owner{}
	if result := database.DB().Preload("Pets").First(&own, form.Owner_ID); result.Error != nil {
		resJson := Response{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	ctx.JSON(http.StatusOK, own)
}

// Delete a specific `Owner`
// @Summary Delete a specific owner
// @Description Delete a specific owner
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param owner_id query uint false "Owner ID"
// @Success 200
// @Router /owner [delete]
func (r *Router) DeleteOwner(ctx *gin.Context) {
	var form DeleteOwnerRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	own := Owner{}
	if result := database.DB().Delete(&own, form.Owner_ID); result.Error != nil {
		resJson := Response{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := Response{
		Success: true,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, resJson)
}

// Update a specific `Owner`
// @Summary Update a specific owner
// @Description Update a specific owner
// @Tags Owner
// @Accept  json
// @Produce  json
// @Param Body body api.UpdateOwnerRequest true "Update user body"
// @Success 200
// @Router /owner [patch]
func (r *Router) UpdateOwner(ctx *gin.Context) {
	var json UpdateOwnerRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the owner we would like to update
	own := Owner{}
	if result := database.DB().First(&own, json.Owner_ID); result.Error != nil {
		res := Response{
			Success: false,
			Object:  nil,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	own.First_Name = json.First_Name
	own.Last_Name = json.Last_Name
	own.Email_Address = json.Email_Address
	own.Zip_Code = json.Zip_Code

	if result := database.DB().Update(&own); result.Error != nil {
		res := Response{
			Success: false,
			Object:  nil,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := Response{
		Success: true,
		Object:  nil,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, res)
}
