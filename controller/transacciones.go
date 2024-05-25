package controller

import (
	"fmt"
	"net/http"
	"rest-api/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateTransaccionInput struct {
	Idmerchant uint   `json:"idmerchant"`
	Amount     string `json:"Amount"`
}

type UpdateTransaccionInput struct {
	Idmerchant uint   `json:"idmerchant"`
	Fee        string `json:"Fee"`
	Amount     string `json:"Amount"`
	Commission string `json:"commi"`
	Created_at string `json:"Created_at"`
}

// GET /Transaccion
// Find all Transaccions
func FindTransaccions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Transaccions []model.Transaction
	db.Find(&Transaccions)

	c.JSON(http.StatusOK, gin.H{"data": Transaccions})
}

// GET /Transaccion/:id
// Find a Transaccion
func FindTransaccion(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model
	var Transaccion model.Transaction
	if err := db.Where("id = ?", c.Param("id")).First(&Transaccion).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Transaccion})
}

// POST /Transaccion
// Create new Transaccion
func CreateTransaccion(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateTransaccionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Amount, err := strconv.ParseFloat(input.Amount, 64)

	if Amount < 1 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el Monto debe ser mayor a cero"})
		return
	}
	// Get model
	var merchant model.Comercio
	if err := db.Where("merchant_id = ?", input.Idmerchant).First(&merchant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El comercio no existe!"})
		return
	}

	merchantcomi, errc := strconv.ParseFloat(merchant.Commission, 64)
	if errc != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en datos "})
		return
	}

	comision := (Amount * merchantcomi) / 100

	// Create Transaccion
	monto, err := strconv.ParseFloat(input.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El monto debe ser un numero mayor a Cero "})
		return
	}

	Transaccion := model.Transaction{
		Merchant_id: int(input.Idmerchant),
		Fee:         fmt.Sprintf("%f", comision),
		Amount:      monto,
		Commission:  merchant.Commission,
		Created_at:  int(time.Now().Local().Unix()),
	}

	db.Create(&Transaccion)

	c.JSON(http.StatusOK, gin.H{"data": Transaccion})
}

// DELETE /Transaccion/:id
// Delete a Transaccion
func DeleteTransaccion(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var Transaccion model.Transaction
	if err := db.Where("id = ?", c.Param("id")).First(&Transaccion).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&Transaccion)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
