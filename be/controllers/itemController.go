package controllers

import (
	"github.com/gin-gonic/gin"
	"inventory-management/models"
	"inventory-management/services"
	"net/http"
	"strconv"
)

type ItemController struct {
	Service *services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{Service: service}
}

func (c *ItemController) CreateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdItem, err := c.Service.Create(&item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}
	ctx.JSON(http.StatusCreated, createdItem)
}

func (c *ItemController) GetItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	item, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *ItemController) UpdateItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedItem, err := c.Service.Update(id, &item)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	ctx.JSON(http.StatusOK, updatedItem)
}

func (c *ItemController) DeleteItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *ItemController) GetAllItems(ctx *gin.Context) {
	items, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}
	ctx.JSON(http.StatusOK, items)
}
