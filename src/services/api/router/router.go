package router

import "github.com/gin-gonic/gin"

//NewRouter return router object
func NewRouter() *gin.Engine {
	r := gin.Default()

	//set userGrop router
	userGroup := r.Group("/users")
	userGroup.POST("/", RegisterUser())

	return r
}
