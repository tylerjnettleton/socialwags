package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/tylerjnettleton/socialwags/pkg/api"
	_ "github.com/tylerjnettleton/socialwags/pkg/api/docs"
	"github.com/tylerjnettleton/socialwags/pkg/database"
)

func main() {

	// Config and connect our database
	dbConfg := database.DatabaseConfig{Name: "socialwags.db"}
	err := database.Connect(dbConfg, 3)

	if err != nil {
		panic("failed to connect database")
	}
	defer database.DB().Close()

	// Setup our router
	r := gin.Default()
	router, err := api.NewRouter()

	if err != nil {
		panic("Failed to create http router")
	}

	// Setup our routes
	v1 := r.Group("/api/v1")
	{
		// Owner endpoints
		owner := v1.Group("/owner")
		{
			owner.PUT("", router.CreateOwner)
			owner.GET("", router.GetOwner)
			owner.DELETE("", router.DeleteOwner)
			owner.PATCH("", router.UpdateOwner)
		}

		// pet endpoints
		pet := v1.Group("/pet")
		{
			pet.PUT("", router.CreatePet)
			pet.GET("", router.GetPet)
			pet.DELETE("", router.DeletePet)
			pet.PATCH("", router.UpdatePet)
		}

		// post endpoints
		post := v1.Group("/post")
		{
			post.PUT("", router.CreatePost)
			post.GET("", router.GetPost)
			post.DELETE("", router.DeletePost)
			post.PATCH("", router.UpdatePost)
		}
	}

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
