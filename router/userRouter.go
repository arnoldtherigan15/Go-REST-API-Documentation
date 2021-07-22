package router

import (
	"github.com/arnoldtherigan15/Go-REST-API-Documentation/handler"
	"github.com/arnoldtherigan15/Go-REST-API-Documentation/user"
	"github.com/gin-gonic/gin"
)

func (r *routes) addUser(rg *gin.RouterGroup) {

	userRepository := user.NewRepository(r.db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewHandler(userService)

	user := rg.Group("/users")
	{
		user.POST("/register", userHandler.Register)
	}
}
