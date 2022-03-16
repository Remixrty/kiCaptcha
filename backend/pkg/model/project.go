package model

import "gorm.io/gorm"

//stuct for projects
//-------------------
type Project struct {
	gorm.Model
	Id        uint   `json:"id"`
	UserId    int    `json:"userid"`
	URL       string `json:"url"`
	SiteKey   string `json:"sitekey"`
	SecretKey string `json:"secretkey"`
}

func (pr *Project) CreateProject() {

	db := new(DB)
	db.InitDb()

	pr = &Project{Id: 2, UserId: 2, URL: "wow", SiteKey: "www", SecretKey: "alex"}
	db.Db.Create(&pr)

}

func (pr *Project) UpdateProject() {
	db := new(DB)
	db.InitDb()

	db.Db.First(&pr)

	pr.URL = "lol"
	db.Db.Save(&pr)

}
func (pr *Project) DeleteProject() {
	db := new(DB)
	db.InitDb()

	db.Db.First(&pr).Delete(&pr)

}
