package setup

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
	routes.StudentRoute(app)
	return app
}