package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvpmatch/server/database"
	"github.com/mvpmatch/server/models"
)


func AddProduct(context *gin.Context) {
	var user models.User
	var product models.Product
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if username exists and password is correct
	record := database.Instance.Where("id = ?", product.SellerID).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	if user.Role != "seller" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	createRecord := database.Instance.Create(&product)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": createRecord.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sellerId": user.ID, "productName": product.ProductName, "cost": product.Cost, "amountAvailable": product.AmountAvailable})
}




