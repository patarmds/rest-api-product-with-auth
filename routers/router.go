package routers

import (
	"belajar-middleware/controllers"
	"belajar-middleware/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.POST("/create", middlewares.Authentication(), controllers.CreateProduct)
		productRouter.PUT("/:productID", middlewares.Authentication(), middlewares.ProductAuthorization(), middlewares.AdminAuthorization(), controllers.UpdateProduct)
		productRouter.GET("/:productID", middlewares.Authentication(), middlewares.ProductAuthorization(), controllers.GetProduct)
		productRouter.GET("/", middlewares.Authentication(), middlewares.AdminAuthorization(), controllers.GetProducts)
		
		productRouter.DELETE("/:productID", middlewares.Authentication(), middlewares.ProductAuthorization(), middlewares.AdminAuthorization(), controllers.DeleteProduct)

	}

	return router
}
