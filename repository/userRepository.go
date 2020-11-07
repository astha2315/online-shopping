package repository

import (
	"errors"
	"fmt"
	"online-shopping/Config"
	"online-shopping/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	CloseDB()
	GetUserByUsernameAndPassword(userName string, password string) (*models.User, error)
	CreateItem(item *models.Item) error
	AddItemsToCart(cart *models.Cart) error

	GetUserByUserId(userId int) (*models.User, error)

	UpdateCartByIsPurchased(cartId int) error

	GetAllUsers() ([]*models.User, error)
	GetAllItems() ([]*models.Item, error)

	GetAllCarts() ([]*models.Cart, error)
}

type database struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	var err error

	db, err := Config.DbConnect()
	if err != nil {
		fmt.Println("FAILED TO CONNECT:", err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Cart{})

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()

	if err != nil {
		fmt.Println("Failed to close database")
	}
}

func (db *database) CreateUser(user *models.User) error {

	if err := db.connection.Create(user).Error; err != nil {
		return err
	}

	user.Cart.UserId = int(user.Id)

	db.connection.Model(&user.Cart).Update("userId", user.Id)

	return nil
}

func (db *database) GetUserByUsernameAndPassword(userName string, password string) (*models.User, error) {

	userData := new(models.User)

	if err := Config.DB.Where("user_name = ? AND password = ?", userName, password).Find(userData).Error; err != nil {
		return nil, err
	}
	return userData, nil
}

func (db *database) CreateItem(item *models.Item) error {

	if err := db.connection.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (db *database) AddItemsToCart(cart *models.Cart) error {

	if err := db.connection.Model(cart).Association("Items").Append(cart.Items).Error; err != nil {
		return err
	}
	return nil
}

func (db *database) GetUserByUserId(userId int) (*models.User, error) {

	userData := new(models.User)

	if err := db.connection.First(userData, userId).Error; err != nil {
		return nil, err
	}
	return userData, nil
}

func (db *database) UpdateCartByIsPurchased(cartId int) error {

	cart := new(models.Cart)
	cart.Id = uint(cartId)
	if err := db.connection.Model(cart).Where("id =?", cartId).Update("is_purchased", true).Error; err != nil {
		return err
	}

	return nil
}

func (db *database) GetAllUsers() ([]*models.User, error) {

	var users []*models.User
	result := db.connection.Find(&users)

	if result != nil && result.RecordNotFound() {
		return nil, errors.New("No record found")
	}

	if result != nil && result.Error != nil {
		return nil, result.Error
	}

	return users, nil

}

func (db *database) GetAllItems() ([]*models.Item, error) {

	db.connection.LogMode(true)
	var items []*models.Item
	result := db.connection.Find(&items)

	if result != nil && result.RecordNotFound() {
		return nil, errors.New("No record found")
	}

	if result != nil && result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (db *database) GetAllCarts() ([]*models.Cart, error) {

	db.connection.LogMode(true)
	var carts []*models.Cart

	// item := models.Item{}

	// result := db.connection.Model(carts).Related(&item, "Items")

	result := db.connection.Find(&carts)
	if result != nil && result.RecordNotFound() {
		return nil, errors.New("No record found")
	}

	if result != nil && result.Error != nil {
		return nil, result.Error
	}

	return carts, nil
}
