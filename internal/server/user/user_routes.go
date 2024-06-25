package user

import "github.com/gin-gonic/gin"

func LoadUserRoutes(routerGrp *gin.RouterGroup, usrh *UserSrvPkg) {
	// GET routes
	routerGrp.GET("/user/id/:id", usrh.GetUserByID)
	
	// POST routes
	routerGrp.POST("/create/user", usrh.PostUser)
}
