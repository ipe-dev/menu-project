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
	authMiddleware := middleware.NewAuthMddleware(userUseCase)
	jwtMiddleware, err := authMiddleware.NewGinJWTMiddleware()
	if err != nil {
		log.Println(err)
	}

	r.POST("/login", jwtMiddleware.LoginHandler)
	r.POST("/logout", jwtMiddleware.LogoutHandler)
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

	// sub_food_stuff
	subFoodStuffPersistence := persistence.NewSubFoodStuffPersistence()
	subFoodStuffUseCase := usecase.NewSubFoodStuffUseCase(subFoodStuffPersistence)
	subFoodStuffHandler := handler.NewSubFoodStuffHandler(subFoodStuffUseCase)

	subfoodstuff := r.Group("/api/sub-foodstuff")
	subfoodstuff.Use(jwtMiddleware.MiddlewareFunc())
	{
		subfoodstuff.POST("/get/list", subFoodStuffHandler.HandleGetList())
		subfoodstuff.POST("/get", subFoodStuffHandler.HandleGet())
		subfoodstuff.POST("/create", subFoodStuffHandler.HandleBulkCreate())
		subfoodstuff.POST("/update", subFoodStuffHandler.HandleBulkUpdate())
		subfoodstuff.POST("/status", subFoodStuffHandler.HandleChangeStatus())
	}
	r.Run()
}
