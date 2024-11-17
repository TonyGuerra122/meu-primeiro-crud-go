package routes

import (
	"github.com/TonyGuerra122/meu-primeiro-crud-go/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/get-user-by-id/:userId", controllers.FindUserById)
	r.GET("/get-user-by-email/:userEmail", controllers.FindUserByEmail)
	r.POST("/create-user", controllers.CreateUser)
	r.PUT("/update-user", controllers.UpdateUser)
	r.DELETE("/delete-user/:userId", controllers.DeleteUser)
}
