package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvpmatch/server/database"
	"github.com/mvpmatch/server/models"
)

// DeleteProductRequest struct
type DeleteProductRequest struct {
	ProductID uint
}

// BuyProductRequest struct
type BuyProductRequest struct {
	Amount    int  `json:"amount" binding:"required"`
	ProductID uint `json:"productID" binding:"required"`
}

// Ping is a handler to check the validity of a user token
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

//DeleteProduct deletes a product item from the database with a specific productID
func DeleteProduct(context *gin.Context) {
	var request DeleteProductRequest
	var product models.Product
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if username exists and password is correct
	record := database.Instance.Where("id = ?", request.ProductID).First(&product)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	userRecord := database.Instance.Where("username = ?", context.Params.ByName("username")).First(&user)
	if userRecord.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": userRecord.Error.Error()})
		context.Abort()
		return
	}

	if user.ID != product.SellerID {
		context.JSON(http.StatusBadRequest, gin.H{"error": "must be seller to delete"})
		context.Abort()
		return
	}

	database.Instance.Delete(&product)
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username, "productName": product.ProductName, "productID": product.ID})

}

// BuyProduct handles logic for adjusting user and product tables when an item is purchase
func BuyProduct(context *gin.Context) {
	var request BuyProductRequest
	var product models.Product
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if product exists and password is correct
	record := database.Instance.Where("id = ?", request.ProductID).First(&product)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	userRecord := database.Instance.Where("username = ?", context.Params.ByName("username")).First(&user)
	if userRecord.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": userRecord.Error.Error()})
		context.Abort()
		return
	}

	totalCost := product.Cost * request.Amount
	err := user.DecreaseFunds(totalCost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err = product.DecreaseStock(request.Amount)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	change, err := user.GetChange()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record.Save(&product)
	userRecord.Save(&user)
	context.JSON(http.StatusOK, gin.H{"change": change, "deposit": user.Deposit})
}
