package service

import (
	"errors"
	"fmt"
	"online-shopping/Config"
	"online-shopping/models"

	"online-shopping/repository"

	_ "github.com/go-sql-driver/mysql"
)

type userService struct{}

type UserServiceIF interface {
	GetAllUsers() ([]*models.User, error)
	CreateUser(user *models.User) (err error)
	GetUserByID(user *models.User, id string) (err error)
	UpdateUser(user *models.User, id string) (err error)
	DeleteUser(user *models.User, id string) (err error)
	CreateItem(item *models.Item) (err error)
	AddItemsToCart(cart *models.Cart) (err error)
	ConvertCartToOrder(cartId int, userId int) (err error)
	GetAllItems() ([]*models.Item, error)
	GetAllCarts() ([]*models.Cart, error)
}

func UserService() UserServiceIF {
	return &userService{}
}

//GetAllUsers Fetch all user data
func (self *userService) GetAllUsers() ([]*models.User, error) {

	users, err := repository.NewUserRepository().GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

//CreateUser ... Insert New data
func (self *userService) CreateUser(user *models.User) (err error) {

	cart := new(models.Cart)
	cart.IsPurchased = false
	user.Cart = cart

	err = repository.NewUserRepository().CreateUser(user)

	return err
}

//GetUserByID ... Fetch only one user by Id
func (self *userService) GetUserByID(user *models.User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func (self *userService) UpdateUser(user *models.User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func (self *userService) DeleteUser(user *models.User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

func (self *userService) CreateItem(item *models.Item) (err error) {

	err = repository.NewUserRepository().CreateItem(item)

	return err
}

func (self *userService) AddItemsToCart(cart *models.Cart) (err error) {

	userInfo, err := repository.NewUserRepository().GetUserByUserId(cart.UserId)

	if err != nil {
		return err
	}

	if userInfo == nil {
		return errors.New("No data found for userId")
	}

	cart.Id = uint(userInfo.CartId)

	err = repository.NewUserRepository().AddItemsToCart(cart)

	return err

}

func (self *userService) ConvertCartToOrder(cartId int, userId int) (err error) {

	userInfo, err := repository.NewUserRepository().GetUserByUserId(userId)

	if err != nil {
		return err
	}

	if userInfo == nil {
		return errors.New("No data found for the user")
	}

	if userInfo.CartId != cartId {
		return errors.New("The cart does not belong to the user")
	}

	err = repository.NewUserRepository().UpdateCartByIsPurchased(cartId)

	if err != nil {
		return err
	}

	return nil
}

func (self *userService) GetAllItems() ([]*models.Item, error) {

	items, err := repository.NewUserRepository().GetAllItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (self *userService) GetAllCarts() ([]*models.Cart, error) {
	carts, err := repository.NewUserRepository().GetAllCarts()
	if err != nil {
		return nil, err
	}
	return carts, nil
}
