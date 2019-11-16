package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tylerjnettleton/socialwags/pkg/api"
	"github.com/tylerjnettleton/socialwags/pkg/database"
)

func main ()  {

	// Gorm
	db, err := gorm.Open("sqlite3", "socialwags.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&database.Owner{})
	db.AutoMigrate(&database.Pet{})
	db.AutoMigrate(&database.Post{})
	db.AutoMigrate(&database.Comment{})
	db.AutoMigrate(&database.Pack{})
	db.AutoMigrate(&database.Pack_Message{})



	// Setup our router
	r := gin.Default()

	router := api.NewRouter()

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

	r.Run()

}
