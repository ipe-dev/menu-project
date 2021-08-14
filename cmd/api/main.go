package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu_project/domain/service"
	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/ipe-dev/menu_project/infrastructure/factory"
	"github.com/ipe-dev/menu_project/infrastructure/persistence"
	"github.com/ipe-dev/menu_project/interface/handler"
	"github.com/ipe-dev/menu_project/interface/middleware"
	"github.com/ipe-dev/menu_project/usecase"
)

func init() {
	database.Connect()
}
func main() {
	r := gin.Default()

	// user
	userPersistence := persistence.NewUserPersistence()
	userService := service.NewUserService(userPersistence)
	userUseCase := usecase.NewUserUseCase(userPersistence, userService)
	userHandler := handler.NewUserHandler(userUseCase)
	authMiddleware := middleware.NewAuthMiddleware(userUseCase)
	jwtMiddleware, err := authMiddleware.NewGinJWTMiddleware()
	headerMiddleWare := middleware.NewHeaderMiddleware()

	r.Use(headerMiddleWare.SetHeader())
	if err != nil {
		log.Println(err)
	}
	errorMiddleware := middleware.NewErrorMiddleware()
	if err != nil {
		log.Println(err)
	}
	r.Use(errorMiddleware.ErrorHandle())

	r.POST("/api/login", jwtMiddleware.LoginHandler)
	r.POST("/api/logout", jwtMiddleware.LogoutHandler)
	user := r.Group("/api/user")
	{
		user.POST("/get", userHandler.Get())
		user.POST("/create", userHandler.Create())
		user.POST("/update", userHandler.Update())
	}

	// menu
	menuPersistence := persistence.NewMenuPersistence()
	weekIDFactory := factory.NewWeekIDFactory()
	menuUsecase := usecase.NewMenuUseCase(menuPersistence, weekIDFactory)
	menuHandler := handler.NewMenuHandler(menuUsecase)
	menu := r.Group("/api/menu")
	menu.Use(jwtMiddleware.MiddlewareFunc())
	{
		menu.POST("/get/list", menuHandler.HandleGetList())
		menu.POST("/get", menuHandler.HandleGet())
		menu.POST("/create", menuHandler.HandleBulkCreate())
		menu.POST("/update", menuHandler.HandleBulkUpdate())
	}

	// sub_menu
	subMenuPersistence := persistence.NewSubMenuPersistence()
	subMenuUsecase := usecase.NewSubMenuUseCase(subMenuPersistence)
	subMenuHandler := handler.NewSubMenuHandler(subMenuUsecase)

	submenu := r.Group("/api/sub-menu")
	submenu.Use(jwtMiddleware.MiddlewareFunc())
	{
		submenu.POST("/get/list", subMenuHandler.HandleGetList())
		submenu.POST("/get", subMenuHandler.HandleGet())
		submenu.POST("/create", subMenuHandler.HandleBulkCreate())
		submenu.POST("/update", subMenuHandler.HandleBulkUpdate())
	}

	// food_stuff
	foodStuffPersistence := persistence.NewFoodStuffPersistence()
	foodStuffUseCase := usecase.NewFoodStuffUseCase(foodStuffPersistence)
	foodStuffHandler := handler.NewFoodStuffHandler(foodStuffUseCase)

	foodstuff := r.Group("/api/foodstuff")
	foodstuff.Use(jwtMiddleware.MiddlewareFunc())
	{
		foodstuff.POST("/get/list", foodStuffHandler.HandleGetList())
		foodstuff.POST("/get", foodStuffHandler.HandleGet())
		foodstuff.POST("/create", foodStuffHandler.HandleBulkCreate())
		foodstuff.POST("/update", foodStuffHandler.HandleBulkUpdate())
		foodstuff.POST("/status", foodStuffHandler.HandleChangeBuyStatus())
	}

	r.Run()
}
