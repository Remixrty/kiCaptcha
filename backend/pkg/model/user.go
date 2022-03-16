package model

import "gorm.io/gorm"

//struct for users
//----------------
type User struct {
	gorm.Model
	Id       uint      `json:"-"`
	Name     string    `json:"name" binding:"required"`
	LastName string    `json:"lastname" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Projects []Project `gorm:"foreignkey:UserId;references:Id" json:"projects,omitempty"`
}

func (us *User) CreateUser() {
	db := new(DB)
	db.InitDb()

	us = &User{Id: 2, Name: "Alex", LastName: "Litovchenko", Password: "ajk354rmlet", Phone: "89041006234", Email: "e22l16h@mail.ru"}
	db.Db.Create(&us)

}

func (us *User) UpdateUser() {
	db := new(DB)
	db.InitDb()

	db.Db.First(&us)

	us.Password = "lol"
	db.Db.Save(&us)

}
func (us *User) DeleteUser() {
	db := new(DB)
	db.InitDb()

	db.Db.First(&us).Delete(&us)

}
