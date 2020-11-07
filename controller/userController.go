//Controllers/User.go
package controller

import (
	"fmt"
	"net/http"
	"online-shopping/middleware"
	"online-shopping/models"
	"online-shopping/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := service.UserService().CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := service.UserService().GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := service.UserService().GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = service.UserService().UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := service.UserService().DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func CreateItem(c *gin.Context) {
	var item models.Item
	c.BindJSON(&item)
	err := service.UserService().CreateItem(&item)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func AddItemsToCart(c *gin.Context) {
	var cart models.Cart
	err := c.BindJSON(&cart)

	if err != nil {
		c.JSON(http.StatusOK, err.Error)
	}

	claims := middleware.GetJwtClaimsFromRequest(c)

	userId := claims["userId"]

	fmt.Println(userId)
	cart.UserId = int(userId.(float64))

	err = service.UserService().AddItemsToCart(&cart)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func ConvertCartToOrder(c *gin.Context) {
	cartId, err := strconv.Atoi(c.Param("cartId"))

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	claims := middleware.GetJwtClaimsFromRequest(c)

	userId := int(claims["userId"].(float64))

	err = service.UserService().ConvertCartToOrder(userId, cartId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, "Cart successfully converted to order")
	}
}

func GetAllUsers(c *gin.Context) {

	userData, err := service.UserService().GetAllUsers()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userData)
	}
}

func GetAllItems(c *gin.Context) {

	items, err := service.UserService().GetAllItems()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func GetAllCarts(c *gin.Context) {

	carts, err := service.UserService().GetAllCarts()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, carts)
	}
}
