package controllers

import (
	"github.com/gin-gonic/gin"
	"inventory-management/models"
	"inventory-management/services"
	"net/http"
	"strconv"
)

type InventoryController struct {
	Service *services.InventoryService
}

func NewInventoryController(service *services.InventoryService) *InventoryController {
	return &InventoryController{Service: service}
}

func (c *InventoryController) CreateInventory(ctx *gin.Context) {
	var inventory models.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdInventory, err := c.Service.Create(&inventory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory"})
		return
	}
	ctx.JSON(http.StatusCreated, createdInventory)
}

func (c *InventoryController) GetInventory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	inventory, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	ctx.JSON(http.StatusOK, inventory)
}

func (c *InventoryController) UpdateInventory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var inventory models.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedInventory, err := c.Service.Update(id, &inventory)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	ctx.JSON(http.StatusOK, updatedInventory)
}

func (c *InventoryController) DeleteInventory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (ic *InventoryController) GetAllInventories(c *gin.Context) {
	inventories, err := ic.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventories"})
		return
	}
	c.JSON(http.StatusOK, inventories)
}
