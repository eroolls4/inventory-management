package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"inventory-management/controllers"
	"inventory-management/initializers"
	"inventory-management/middleware"
	"inventory-management/repository"
	"inventory-management/services"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.ConnectToRedis()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("Hello World")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           30 * time.Hour,
	}))
	db := initializers.DB

	// Repositories
	inventoryRepo := repository.NewInventoryRepository(db)
	itemRepo := repository.NewItemRepository(db)
	restockRepo := repository.NewRestockRepository(db)

	// Services
	inventoryService := services.NewInventoryService(inventoryRepo)
	itemService := services.NewItemService(itemRepo)
	restockService := services.NewRestockService(itemRepo, restockRepo)

	// Controllers
	inventoryController := controllers.NewInventoryController(inventoryService)
	itemController := controllers.NewItemController(itemService)
	restockController := controllers.NewRestockController(restockService)

	// Auth routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.POST("/signout", middleware.RequireAuth, controllers.Signout)

	// Protected API routes
	api := router.Group("/api").Use(middleware.RequireAuth)
	{
		api.GET("/inventory/:id", inventoryController.GetInventory)
		api.POST("/inventory", inventoryController.CreateInventory)
		api.PUT("/inventory/:id", inventoryController.UpdateInventory)
		api.DELETE("/inventory/:id", inventoryController.DeleteInventory)

		api.GET("/items/:id", itemController.GetItem)
		api.POST("/items", itemController.CreateItem)
		api.PUT("/items/:id", itemController.UpdateItem)
		api.DELETE("/items/:id", itemController.DeleteItem)

		api.POST("/items/:id/restock", restockController.RestockItem)

		api.GET("/items", itemController.GetAllItems)
		api.GET("/inventory", inventoryController.GetAllInventories)
		api.GET("/items/:id/restock-history", restockController.GetRestockHistory)
	}

	router.Run()
}
