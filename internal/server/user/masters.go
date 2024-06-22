package user

import (
	"net/http"
	"strconv"

	log "go-url-shortner/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (usrh *UserSrvPkg) GetUserByID(c *gin.Context) {
	usrIdStr := c.Param("id")
	var userID int
	userID, err := strconv.Atoi(usrIdStr)
	if err != nil {
		log.LogError(err, "error inside GetUserByID: couldn't parse user id", map[string]interface{}{"userid": usrIdStr})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id passed, please provide correct user id.",
		})
		return
	}

	log.LogInfo(nil, "user info parsed successfully", map[string]interface{}{"userID": userID})

	c.JSON(http.StatusOK, gin.H{
		"Status":"success",
		"Msg": "User parsed successfully",
		"UserID": userID,
	})
}
