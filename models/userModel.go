//Models/UserModel.go
package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"id"  gorm:"primary_key:auto_increment" `
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	UserName  string    `json:"userName"  gorm:"type:varchar(20)"`
	Password  string    `json:"password"  gorm:"type:varchar(20)"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP `
	CartId    int
	Cart      *Cart `gorm:"foreignKey:CartId"`
}

type Cart struct {
	Id          uint      `json:"id"  gorm:"primary_key:auto_increment" `
	IsPurchased bool      `json:"isPurchased"  gorm:"type:bool"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP `
	Items       []*Item   `gorm:"many2many:cart_item;"`
	UserId      int
	User        *User `gorm:"foreignKey:UserId"`
}

type Item struct {
	Id        uint      `json:"id"  gorm:"primary_key:auto_increment" `
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Carts     []*Cart   ` gorm:"many2many:cart_item;"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP `
}

type Order struct {
	CartId    uint      `json:"cartId"  gorm:"primary_key:auto_increment" `
	UserId    uint      `json:"userId" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP `
	Cart      Cart      `gorm:"foreignKey:CartId"`
	User      User      `gorm:"foreignKey:UserId"`
}
