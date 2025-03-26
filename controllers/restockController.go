package controllers

import (
	"github.com/gin-gonic/gin"
	"inventory-management/services"
	"net/http"
	"strconv"
)

type RestockController struct {
	Service *services.RestockService
}

func NewRestockController(service *services.RestockService) *RestockController {
	return &RestockController{Service: service}
}

func (c *RestockController) RestockItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var requestDTO struct {
		Amount int `json:"amount"`
	}
	if err := ctx.ShouldBindJSON(&requestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := c.Service.RestockItem(id, requestDTO.Amount)
	if err != nil {
		switch err.Error() {
		case "restock amount must be between 10 and 1000":
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case "rate limit exceeded: max 3 restocks per 24 hours":
			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restock item"})
		}
		return
	}
	ctx.JSON(http.StatusOK, item)
}
