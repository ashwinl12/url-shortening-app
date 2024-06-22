package internal

import (
	"go-url-shortner/internal/server/url"
	"go-url-shortner/internal/server/user"

	"github.com/gin-gonic/gin"
)

var (
	Usrh *user.UserSrvPkg
	Urlh *url.UrlSrvPkg
)

func LoadRoutes(server *gin.Engine) {
	routerGrp := server.Group("api/v1")

	user.LoadUserRoutes(routerGrp, Usrh)
}
