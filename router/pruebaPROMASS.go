package router

import (
	"rest-api/controller"

	"github.com/gin-gonic/gin"
)

// InitializemerchantRoutes initializes the router for merchant api
func InitializemerchantRoutes(r *gin.Engine) {

	// version uno del api
	v1 := r.Group("/v1")

	// Routes merchant
	v1.GET("/merchant", controller.Findmerchants)
	v1.GET("/merchant/:id", controller.Findmerchant)
	v1.POST("/merchant", controller.Createmerchant)
	v1.PATCH("/merchant/:id", controller.Updatemerchant)
	v1.DELETE("/merchant/:id", controller.Deletemerchant)
	// router transacction
	v1.GET("/transaction", controller.FindTransaccions)
	v1.GET("/transaction/:id", controller.FindTransaccion)
	v1.POST("/transaction", controller.CreateTransaccion)

	v1.DELETE("/transaction/:id", controller.DeleteTransaccion)
	// router transacction
	v1.GET("/ReportTotal", controller.FindTransaccionsTotal)
	v1.GET("/ReportDatabyMerchant/:id", controller.FindTransaccionsbyIdMerchant)

}
