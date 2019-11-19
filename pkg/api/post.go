package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tylerjnettleton/socialwags/pkg/database"
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

// Gin Bindings
type CreatePostRequest struct {
	Owner_ID  uint   `json:"owner_id" binding:"required"`
	Post_Type uint   `json:"post_type" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Body      string `json:"body" binding:"required"`
}

type GetPostRequest struct {
	Post_ID uint `form:"post_id" binding:"required"`
}

type DeletePostRequest struct {
	GetPostRequest
}

type UpdatePostRequest struct {
	Post_ID uint   `form:"post_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

// Create a new post
// @Summary Create a new post
// @Description Create a new post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param Body body api.CreatePostRequest true "Create post request body"
// @Success 200
// @Router /post [put]
func (c *Router) CreatePost(ctx *gin.Context) {

	var request CreatePostRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := Post{
		OwnerID:     request.Owner_ID,
		Post_Type:   request.Post_Type,
		Picture_URL: "",
		Title:       request.Title,
		Body:        request.Body,
		Time:        time.Time{},
		TimeZone:    "",
	}

	if result := database.DB().Create(&p); result.Error != nil {
		errors := result.GetErrors()
		for _, err := range errors {
			fmt.Println(err)
		}
		// Todo: Do better error handling with json response
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
		Object:  p,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, res)
}

// Get a specific post
// @Summary Get a specific post
// @Description Get a specific post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param post_id query uint false "Post ID"
// @Success 200
// @Router /post [get]
func (c *Router) GetPost(ctx *gin.Context) {
	var form GetPostRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := Post{}
	if result := database.DB().Preload("Comments").First(&p, form.Post_ID); result.Error != nil {
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
		Object:  p,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, res)
}

// Delete a specific `Post`
// @Summary Delete a specific post
// @Description Delete a specific post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param post_id query uint false "Post ID"
// @Success 200
// @Router /post [delete]
func (c *Router) DeletePost(ctx *gin.Context) {
	var form DeletePostRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := Post{}
	if result := database.DB().Delete(&p, form.Post_ID); result.Error != nil {
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
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, res)
}

// Update a specific `Post`
// @Summary Update a specific post
// @Description Update a specific post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param Body body api.UpdatePostRequest true "Update post body"
// @Success 200
// @Router /post [patch]
func (c *Router) UpdatePost(ctx *gin.Context) {
	var req UpdatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := Post{}
	if result := database.DB().First(&p, req.Post_ID); result.Error != nil {
		res := Response{
			Success: false,
			Object:  nil,
			Error:   nil,
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	p.Title = req.Title
	p.Body = req.Body

	if result := database.DB().Update(&p); result.Error != nil {

		errors := result.GetErrors()
		for _, err := range errors {
			fmt.Println(err)
		}

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
		Object:  p,
		Error:   nil,
	}
	ctx.JSON(http.StatusOK, res)

}
