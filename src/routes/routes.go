package routes

import (
	"github.com/gin-gonic/gin"
	handle_user "github.com/guizot/go-gin-mysql/src/handlers/user"
)

type Routes struct {
}

func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", handle_user.GetAllUser)
		api.GET("/users/:id", handle_user.GetUser)
	}
	r.Run(":8000")
}