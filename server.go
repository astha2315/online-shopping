package main

import (
	"net/http"
	"online-shopping/controller"
	"online-shopping/middleware"
	"online-shopping/service"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	// Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	// if err != nil {
	// 	fmt.Println("Status:", err)
	// }
	// defer Config.DB.Close()
	// Config.DB.AutoMigrate(&models.User{})
	// r := routes.SetupRouter()
	// //running
	// r.Run()

	server := gin.New()
	// server.Use(gin.Recovery(), gin.Logger())

	// server.POST("/login",)
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	server.POST("/user/create", controller.CreateUser)

	server.POST("/item/create", controller.CreateItem)

	server.POST("/cart/add", middleware.AuthorizeJWT(controller.AddItemsToCart))
	server.GET("/cart/:cartId/complete", middleware.AuthorizeJWT(controller.ConvertCartToOrder))

	server.GET("/user/list", controller.GetAllUsers)

	server.GET("/item/list", controller.GetAllItems)

	server.GET("/carts/list", controller.GetAllCarts)

	port := "8080"
	server.Run(":" + port)

}

// func AuthorizeJWT(next gin.HandlerFunc) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		const BEARER_SCHEMA = "Bearer"
// 		authHeader := c.GetHeader("Authorization")
// 		tokenString := authHeader[len(BEARER_SCHEMA):]
// 		fmt.Println(tokenString)
// 		// encrypted_text, err := strconv.Unquote(`"` + tokenString + `"`)
// 		token, err := service.JWTAuthService().ValidateToken(strings.TrimSpace(tokenString))
// 		if token.Valid {
// 			claims := token.Claims.(jwt.MapClaims)
// 			fmt.Println(claims)

// 			next(c)
// 		} else {
// 			fmt.Println(err)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 	}
// }
