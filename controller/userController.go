package controller

import (
	"github.com/gin-gonic/gin"
	"golang_crud/model"
	"golang_crud/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// SetAllUserController is a function to set all handler of user entity
func SetAllUserController(r *gin.Engine) {
	userGroup := r.Group("/user")
	// Insert of user
	userGroup.POST("", InsertUser)
	// Delete users by user ids
	userGroup.DELETE("/:userIds", DeleteUserByUserIds)
	// Search user by user id
	userGroup.GET("/:userId", SearchUserByUserId)
	// List of user
	userGroup.GET("", SearchUserList)
	// Update user
	userGroup.PUT("/:userId", UpdateUser)
}

// InsertUser insert a user to database
func InsertUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("Bind insert data happend an error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    service.InsertUser(&user),
			"message": "success",
		})
	}
}

// DeleteUserByUserIds delete users by user ids
func DeleteUserByUserIds(c *gin.Context) {
	userIdsStr := c.Param("userIds")
	values := strings.Split(userIdsStr, ",")
	userIds := make([]uint64, len(values))
	for index, value := range values {
		userIds[index], _ = strconv.ParseUint(value, 10, 64)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"rows":    service.DeleteUserByUserIds(&userIds),
		"message": "success",
	})
}

// SearchUserByUserId search user by userId
func SearchUserByUserId(c *gin.Context) {
	var user model.User
	err := c.ShouldBindUri(&user)
	if err != nil {
		log.Printf("Bind query uri happend an error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    service.SearchUserByUserId(user.UserId),
			"message": "success",
		})
	}
}

// SearchUserList is a function to search all user on the database
func SearchUserList(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Printf("Bind query data happend an error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    service.SearchUserList(&user),
			"message": "success",
		})
	}
}

// UpdateUser update user
func UpdateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("Bind update data happend an error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
	} else {
		err = c.ShouldBindUri(&user)
		if err != nil {
			log.Printf("Bind update data happend an error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"data":    service.UpdateUser(&user),
				"message": "success",
			})
		}
	}
}
