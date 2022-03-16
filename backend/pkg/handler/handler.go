/*поменять апи и роутер с другими ссылками,
см. в frontend/client/src/models/response/AuthResponse.ts*/

package handler

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//Our routes
//-----------------
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// router.Use(cors.Default())
	router.HandleMethodNotAllowed = true
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)      //  регистрация
		auth.POST("/sign-in", h.signIn)      //  вход
		auth.GET("/sign-out", h.signOut)     // выход
		auth.GET("/sign-check", h.signCheck) // для проверки потом удалить
		auth.POST("/deleteuser", h.deleteUser)
	}

	api := router.Group("/api")
	{
		list := api.Group("/lists")
		{
			//list.POST("/")
			list.GET("/", h.getList)
			//list.PUT("/:id")
			//list.DELETE("/:id")

			items := list.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:itemId", h.getItemById)
				items.PUT("/:itemId", h.changeItem)
				items.DELETE("/:itemId", h.deleteItem)
			}
		}
	}

	return router
}
