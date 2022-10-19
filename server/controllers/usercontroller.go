package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvpmatch/server/auth"
	"github.com/mvpmatch/server/database"
	"github.com/mvpmatch/server/models"
)

// DeleteRequest is the input format for deleting a user
type DeleteRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// DepositRequest is the input format for depositing money
type DepositRequest struct {
	DepositAmount int `json:"depositAmount"`
}

// LoginRequest is the input format for logging in
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUserRequest is the input format for retrieving user information
type GetUserRequest struct {
	Token string `json:"token"`
}

// RegisterUser accepts the username and password to create a new user in the database
func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user.Deposit = 0
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username, "role": user.Role, "token": tokenString})
}

// LoginUser checks user credentials and returns a new access token
func LoginUser(context *gin.Context) {
	var user models.User
	var request LoginRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if username exists and password is correct
	record := database.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"userId": user.ID, "username": user.Username, "token": tokenString, "role": user.Role})
}

// DeleteUser deletes a user in the database
func DeleteUser(context *gin.Context) {
	var request DeleteRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if username exists and password is correct
	record := database.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	database.Instance.Delete(&user)
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username, "password": user.Password})
}

// DepositFunds checks the denomination for new deposits and updates the database
func DepositFunds(context *gin.Context) {
	var user models.User
	var request DepositRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if username exists and password is correct
	record := database.Instance.Where("username = ?", context.Params.ByName("username")).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	if user.Role != "buyer" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	err := user.DepositFunds(request.DepositAmount)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record.Save(&user)
	context.JSON(http.StatusOK, gin.H{"username": user.Username, "deposit": user.Deposit})
}

// ResetDeposit sets the users current deposit to zero
func ResetDeposit(context *gin.Context) {
	var user models.User

	// check if username exists and password is correct
	record := database.Instance.Where("username = ?", context.Params.ByName("username")).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	err := user.ResetDeposit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record.Save(&user)
	context.JSON(http.StatusOK, gin.H{"username": user.Username, "deposit": user.Deposit})
}

// GetUserInfo returns the necessary information for the frontend to do access controls and update info
func GetUserInfo(context *gin.Context) {
	var user models.User

	record := database.Instance.Where("username = ?", context.Params.ByName("username")).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"username": user.Username, "deposit": user.Deposit, "role": user.Role})
}
