package controller

import (
	"online-shopping/models"
	"online-shopping/service"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	credential := new(models.LoginCredentials)
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated, userId := controller.loginService.LoginUser(credential.UserName, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(userId, true)

	}
	return ""
}
