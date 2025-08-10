package users

import (
	"event-management-api/models"
	"event-management-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//type User struct {
//	ID int64 `json:"id"`
//}

func SignUp(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing data", "error": err.Error()})
		return
	}

	if err := user.SAVE(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error saving user",
			"error":   err.Error(),
		})
		return
	}
	user.PASSWORD = ""
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": user})
}

func GetAllUsers(context *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error getting users", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Users retrieved successfully", "data": users})
}

func Login(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing data", "error": err.Error()})
		return
	}

	err = user.ValidateUserLogin()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials", "error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.EMAIL, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error generating token", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged in", "data": user.EMAIL, "token": token})
}
