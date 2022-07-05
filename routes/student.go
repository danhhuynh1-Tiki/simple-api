package routes

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func StudentRoute(route *gin.Engine){
	route.GET("/user",AllUser)
}

func AllUser(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message" : "pong",
	})
}