package model

import "gorm.io/gorm"

//struct for pictures
//-------------------
type Picture struct {
	gorm.Model
	Id     uint   `json:"id"`
	Way    string `json:"way"`
	Number int    `json:"number"`
}

func (pic *Picture) CreatePicture() {
	db := new(DB)
	db.InitDb()

	pic = &Picture{Id: 3, Way: "way", Number: 1}
	db.Db.Create(&pic)
}
func (pic *Picture) UpdatePicture() {
	db := new(DB)
	db.InitDb()

	db.Db.First(&pic)

	pic.Way = "lol"
	db.Db.Save(&pic)

}
func (pic *Picture) DeletePicture() {
	db := new(DB)
	db.InitDb()

	db.Db.First(&pic).Delete(&pic)

}
