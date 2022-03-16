package repository

import (
	"fmt"

	"github.com/romanSeleznev/kiCaptcha/backend/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBname   string
	Port     string
	SSLmode  string
}

//auto migration for user and connetciton db
//------------------------------------------
func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.Host, cfg.User, cfg.Password, cfg.DBname, cfg.Port, cfg.SSLmode)),
		&gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	//проверка на ping

	return db, nil
}
