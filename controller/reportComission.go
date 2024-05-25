package controller

import (
	"fmt"
	"net/http"
	"rest-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindTransaccionsbyIdMerchant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model

	var Transaccions []model.Transaction
	if err := db.Where("Merchant_id = ?", c.Param("idm")).Find(&Transaccions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "info not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Transaccions})
}
func FindTransaccionsTotal(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	params := c.Request.URL.Query()

	var sum float32
	query := db.Table("transactions").Select("sum(fee) as sum")

	if params["idm"] != nil && params["date"] != nil {
		query.Where("merchant_id = ? AND Created_at = ?", params["idm"], params["date"]).Row().Scan(&sum)
	} else if params["idm"] != nil {
		fmt.Println(params)

		query.Where("merchant_id = ?", params["idm"]).Scan(&sum)
	} else if params["date"] != nil {
		query.Where("Created_at = ?", params["date"]).Scan(&sum)
	} else {
		query.Scan(&sum)
	}
	fmt.Println(sum)
	fmt.Println(db.Statement.SQL.String())
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprint("el total de comisiones de esta consulta  es ", sum)})
}
