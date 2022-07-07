package routes

import (
	repository "api/repository/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine) {

	// Simple Api : v1
	v1 := route.Group("/v1")
	{
		v1.GET("/user", AllUser)
		v1.GET("/add_user", AddUser)
		v1.GET("/update_user", UpdateUser)
		v1.GET("/delete_user", DeleteUser)
		v1.GET("/find_user", FindUser)
	}
}

func AllUser(c *gin.Context) {
	alluser := repository.GetUser()
	if alluser == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail",
		})
	}
	c.JSON(http.StatusOK, alluser)
}

func AddUser(c *gin.Context) {

	name := c.Query("name")
	check := repository.AddUser(name)

	if check == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Successfull",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Fail",
		})
	}

}

func UpdateUser(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	check := repository.UpdateUser(id, name)

	if check == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Successful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Fail",
		})
	}

}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")

	check := repository.DeleteUser(id)

	if check == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Successful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Fail",
		})
	}
}

func FindUser(c *gin.Context) {
	id := c.Query("id")

	us := repository.FindUser(id)

	if us == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "",
		})
	} else {
		c.JSON(http.StatusOK, us)
	}
}
