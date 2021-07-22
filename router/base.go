package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type routes struct {
	router *gin.Engine
	db     *gorm.DB
}

func SetupRouter(db *gorm.DB) *routes {
	r := routes{
		router: gin.Default(),
		db:     db,
	}

	v1 := r.router.Group("api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]string{
				"message": "Welcome to my server",
			})
		})
	}

	r.addUser(v1)
	return &r
}

func (r routes) Run(addr ...string) error {
	port := fmt.Sprintf(":%s", addr[0])
	return r.router.Run(port)
}
