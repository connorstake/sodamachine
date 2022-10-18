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

	userRecord := database.Instance.Where("username = ?", context.Params.ByName("username")).First(&user)
	if userRecord.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": userRecord.Error.Error()})
		context.Abort()
		return
	}

	if user.Role != "seller" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	product.SellerID = user.ID

	createRecord := database.Instance.Create(&product)
	if createRecord.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": createRecord.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sellerId": user.ID, "productName": product.ProductName, "cost": product.Cost, "amountAvailable": product.AmountAvailable})
}

func GetAllProducts(context *gin.Context) {
	var products []models.Product
	record := database.Instance.Find(&products)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"products": products})
}

func GetAllProductsBySeller(context *gin.Context) {
	var products []models.Product
	var user models.User

	userRecord := database.Instance.Where("username = ?", context.Params.ByName("username")).First(&user)
	if userRecord.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": userRecord.Error.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Find(&products, "seller_id = ?", user.ID)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"products": products})
}
