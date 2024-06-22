package user

import "github.com/gin-gonic/gin"

func LoadUserRoutes(routerGrp *gin.RouterGroup, usrh *UserSrvPkg) {
	routerGrp.GET("/user/id/:id", usrh.GetUserByID)
}
