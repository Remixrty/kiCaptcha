package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/handler"
	"github.com/sirupsen/logrus"

	//"github.com/romanSeleznev/kiCaptcha/backend/pkg/model"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/repository"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/server"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/service"
	"github.com/spf13/viper"
	//"github.com/gin-gonic/gin"
)

func main() {
	// db := new(model.DB)
	// db.Db.AutoMigrate()
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing env variable: %s", err.Error())
	}

	//db := new(model.DB)
	//db.AutoMigrates()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		Port:     viper.GetString("db.port"),
		SSLmode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("erorr initializing db: %s", err.Error())
	}

	repos := repository.NewRepositori(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
