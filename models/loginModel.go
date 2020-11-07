//Models/UserModel.go
package models

type LoginCredentials struct {
	UserName string `form:"userName"`
	Password string `form:"password"`
}
