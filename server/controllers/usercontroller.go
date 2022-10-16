package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvpmatch/server/database"
	"github.com/mvpmatch/server/models"
)

type DeleteRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DepositRequest struct {
	Username      string `json:"username"`
	DepositAmount int    `json:"depositAmount"`
}

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
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username})
}

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
