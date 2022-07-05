package app

import (
	_ "net/http"
	"api/routes"
	"github.com/gin-gonic/gin"
)
func Run(){
	app := SetupRoute()
	app.Run()
}

func SetupRoute() *gin.Engine{
	app := gin.Default()
	routes.UserRoute(app)
	return app
}