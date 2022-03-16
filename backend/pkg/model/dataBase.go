package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (d *DB) AutoMigrates() {
	ds := new(DB)
	ds.InitDb()
	ds.Db.AutoMigrate(&User{}, &Picture{}, &Project{})
}

type DB struct {
	Db  *gorm.DB
	err error
}

func (d *DB) InitDb() error {
	dsn := "host=localhost user=postgres password=ajk354rmlet dbname=KiCaptcha port=5432 sslmode=disable "
	d.Db, d.err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if d.err != nil {
		log.Fatal(d.err)
	}
	return d.err
}
