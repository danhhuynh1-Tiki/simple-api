package routes

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"api/models"
)

func UserRoute(route *gin.Engine){
	route.GET("/user",AllUser)
	route.GET("/add_user",AddUser)
}

func AllUser(c *gin.Context){
	alluser := models.GetUser()
	if alluser == nil{
		c.JSON(http.StatusNotFound,gin.H{
			"message" : "Fail",
		})
	}
	c.JSON(http.StatusOK,alluser)
}

func AddUser(c *gin.Context){

	id := c.Query("id")
	name := c.Query("name")
	check := models.AddUser(id,name)
	
	if check == true{
		c.JSON(http.StatusOK,gin.H{
			"message" : "Successfull",
		})
	}else{
		c.JSON(http.StatusNotFound,gin.H{
			"message" : "Fail",
		})
	}

}