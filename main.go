package BlogManagementSystem

import (
	"BlogManagementSystem/controllers"
	"BlogManagementSystem/database"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "user=musti password=1234 dbname=blogdb host=localhost port=5432 sslmode=disable"
	r := gin.Default()
	database.Connect(dsn)

	// User Routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Post Routes
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// Comment Routes
	r.POST("/comments", controllers.CreateComment)
	r.GET("/comments", controllers.GetComments)
	r.GET("/comments/:id", controllers.GetCommentByID)
	r.DELETE("/comments/:id", controllers.DeleteComment)
}
