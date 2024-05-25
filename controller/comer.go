package controller

import (
	"net/http"
	"rest-api/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreatemerchantInput struct {
	Name      string `json:"name" binding:"required"`
	Comission string `json:"commission" binding:"required"`
}

type UpdatemerchantInput struct {
	Merchant_name string `json:"merchant_name"`
	Commission    string `json:"commission,"`
}

// GET /merchant
// Find all merchants
func Findmerchants(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var merchants []model.Comercio
	db.Find(&merchants)
	c.JSON(http.StatusOK, gin.H{"data": merchants})
}

// GET /merchant/:id
// Find a merchant
func Findmerchant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model
	var merchant model.Comercio
	if err := db.Where("merchant_id = ?", c.Param("id")).First(&merchant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

// POST /merchant
// Create new merchant
func Createmerchant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreatemerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commi, err := strconv.ParseInt(input.Comission, 10, 64)
	if commi < 1 || commi > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La comision debe ser ente 1 y 100 %" + err.Error()})
		return
	}
	// Create merchant

	merchant := model.Comercio{
		Merchant_name: input.Name,
		Commission:    input.Comission,
		Created_at:    int(time.Now().Local().Unix()),
	}
	db.Create(&merchant)

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

// PATCH /merchant/:id
// Update a merchant
func Updatemerchant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var merchant model.Comercio
	idm, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := db.Where("merchant_id = ?", idm).First(&merchant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatemerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commi, err := strconv.ParseInt(input.Commission, 10, 64)
	if commi < 1 || commi > 100 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La comision debe ser ente 1 y 100 %"})
		return
	}
	merchant.Commission = input.Commission
	merchant.Merchant_name = input.Merchant_name
	db.Model(&merchant).UpdateColumns(merchant)

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

// DELETE /merchant/:id
// Delete a merchant
func Deletemerchant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var merchant model.Comercio
	if err := db.Where("id = ?", c.Param("id")).First(&merchant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&merchant)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
