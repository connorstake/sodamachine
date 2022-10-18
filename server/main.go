package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mvpmatch/server/controllers"
	"github.com/mvpmatch/server/database"
	"github.com/mvpmatch/server/middlewares"
)

func main() {
	database.Connect("root:password@tcp(localhost:3306)/mvpmatch?parseTime=true")
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/user/login", controllers.LoginUser)
		api.GET("/products", controllers.GetAllProducts)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.DELETE("/user", controllers.DeleteUser)
			secured.GET("/user", controllers.GetUserInfo)
			secured.POST("/user/reset", controllers.ResetDeposit)
			secured.POST("/deposit", controllers.DepositFunds)
			// PRODUCTS
			secured.GET("/products", controllers.GetAllProductsBySeller)
			secured.POST("/product", controllers.AddProduct)
			secured.POST("/product/buy", controllers.BuyProduct)
			secured.POST("/product/delete", controllers.DeleteProduct)
		}
	}
	return router
}
