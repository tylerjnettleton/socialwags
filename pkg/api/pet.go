package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tylerjnettleton/socialwags/pkg/database"
	"net/http"
)

type Pet struct {
	gorm.Model
	OwnerID     uint
	Name        string
	Birth_Date  string
	Picture_URL string
}

type PetResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

type CreatePetRequest struct {
	Owner_ID    uint   `json:"owner_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Birth_Date  string `json:"birth_date" binding:"required"`
	Picture_URL string
}

type UpdatePetRequest struct {
	CreatePetRequest
	Pet_ID uint `form:"owner_id" binding:"required"`
}

type GetPetRequest struct {
	Pet_ID uint `form:"pet_id" binding:"required"`
}

type DeletePetRequest struct {
	GetPetRequest
}

// Create a new pet
// @Summary Create a new pet
// @Description Create a new pet
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param Body body api.CreatePetRequest true "Create pet request body"
// @Success 200
// @Router /pet [put]
func (c *Router) CreatePet(ctx *gin.Context) {
	var request CreatePetRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet := Pet{
		OwnerID:     request.Owner_ID,
		Name:        request.Name,
		Birth_Date:  request.Birth_Date,
		Picture_URL: request.Picture_URL,
	}

	if result := database.DB().Create(&pet); result.Error != nil {
		resJson := PetResponse{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := OwnerResponse{
		Success: true,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, resJson)
}

// Get a specific pet
// @Summary Get a specific pet
// @Description Get a specific pet
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param pet_id query uint false "Pet ID"
// @Success 200
// @Router /pet [get]
func (r *Router) GetPet(ctx *gin.Context) {
	var form GetPetRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet := Pet{}
	if result := database.DB().First(&pet, form.Pet_ID); result.Error != nil {
		resJson := PetResponse{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	ctx.JSON(http.StatusOK, pet)
}

// Delete a specific pet
// @Summary Delete a specific pet
// @Description Delete a specific pet
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param pet_id query uint false "Pet ID"
// @Success 200
// @Router /pet [delete]
func (r *Router) DeletePet(ctx *gin.Context) {
	var form DeletePetRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet := Pet{}
	if result := database.DB().Delete(&pet, form.Pet_ID); result.Error != nil {
		resJson := PetResponse{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := PetResponse{
		Success: true,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, resJson)
}

// Update pet
// @Summary Update pet
// @Description Update pet
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param Body body api.UpdatePetRequest true "update pet request body"
// @Success 200
// @Router /pet [patch]
func (r *Router) UpdatePet(ctx *gin.Context) {
	var json UpdatePetRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the owner we would like to update
	pet := Pet{}
	if result := database.DB().First(&pet, json.Owner_ID); result.Error != nil {
		resJson := PetResponse{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	pet.Name = json.Name
	pet.Birth_Date = json.Birth_Date
	pet.Picture_URL = json.Picture_URL

	if result := database.DB().Update(&pet); result.Error != nil {
		resJson := PetResponse{
			Success: false,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, resJson)
		return
	}

	resJson := PetResponse{
		Success: true,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, resJson)
}
