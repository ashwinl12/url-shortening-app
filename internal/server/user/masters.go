package user

import (
	"context"
	"net/http"
	"time"

	userDbPkg "go-url-shortner/internal/db/user"
	"go-url-shortner/pkg/logger"
	log "go-url-shortner/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func (usrh *UserSrvPkg) GetUserByID(c *gin.Context) {
	usrIdStr := c.Param("id")

	var userID primitive.ObjectID
	userID, err := primitive.ObjectIDFromHex(usrIdStr)
	if err != nil {
		log.LogError(err, "error inside GetUserByID: couldn't parse user id", map[string]interface{}{"userid": usrIdStr})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id passed, please provide correct user id.",
		})
		return
	}

	// Create a filter for the user ID
	filter := bson.M{"_id": userID}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	var user userDbPkg.User
	if err := usrh.Db.Collection(userDbPkg.UserTName).FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			log.LogError(err, "error inside GetUserByID: user not found", map[string]interface{}{"user_id": userID})
			c.JSON(http.StatusNotFound, map[string]interface{}{"message": "user not found"})
			return
		}

		log.LogError(err, "error inside GetUserByID: error occurred while fetching user from db", map[string]interface{}{"user_id": userID})
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "could not fetch user"})
		return
	}
	
	log.LogInfo(nil, "error inside GetUserByID: User fetched successfully", map[string]interface{}{"user": user})
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": user,
	})
	return
}

func (usrh *UserSrvPkg) PostUser(c *gin.Context) {
	var newUserObj userDbPkg.User

	if err := c.ShouldBindJSON(&newUserObj); err != nil {
		logger.LogError(err, "Error inside PostUser: failed to bind req body", nil)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid user object passed"})
		return
	}

	// validate the user struct
	if err := usrh.Validate.Struct(newUserObj); err != nil {
		logger.LogError(err, "Error inside PostUser: validation failed for new user object", nil)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid user object passed, please check the info provided"})
	}

	// assign new id for user
	newUserObj.ID = primitive.NewObjectID()

	// save the new user object inside db
	queryRes, err := usrh.Db.Collection(userDbPkg.UserTName).InsertOne(nil, newUserObj)
	if err != nil {
		log.LogError(err, "error inside PostUser: could not save new user object inside the db", map[string]interface{}{"user": newUserObj})
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "error occurred while creating new user object !"})
		return
	}

	log.LogInfo(nil, "new user scuccessfully created !", map[string]interface{}{"user_id": queryRes.InsertedID})
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Success", "data": queryRes.InsertedID})
}
