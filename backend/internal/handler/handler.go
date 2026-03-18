package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iflow/monitor/pkg/jwt"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "login"})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register"})
}

func GetUserInfo(c *gin.Context) {
	userID, _ := jwt.GetUserID(c)
	c.JSON(http.StatusOK, gin.H{"id": userID, "username": "admin"})
}

func UpdateUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": []interface{}{}})
}

func GetClusters(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"clusters": []interface{}{}})
}

func CreateCluster(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "created"})
}

func GetDataSources(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"datasources": []interface{}{}})
}

func GetDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "dashboard"})
}
