package event

import (
	"Orbit/internal/seller/event/product"
	"Orbit/internal/seller/event/sale"

	"github.com/gin-gonic/gin"
)

func EventRoutes(r *gin.RouterGroup) {
	//---------Events Routes---------//
	r.POST("/create", CreateAnEventHandler);
	r.GET("/get", GetAllEventsHandler);
	r.GET("/get/:id", GetAnEventHandler);
	r.PUT("/update/:id", UpdateAnEventHandler);
	r.DELETE("/delete/:id", DeleteAnEventHandler);

	// --------Event Product Routes---------//
	r.POST("/:eventId/registerProducts", product.RegisterProductsHandler);
	r.GET("/:eventId/getProducts", product.GetAllEventProductsHandler);
	r.GET("/:eventId/getProduct/:productId", product.GetAnEventProductHandler);
	r.PUT("/:eventId/updateProduct/:productId", product.UpdateAnEventProductHandler);
	r.DELETE("/:eventId/deleteProduct/:productId", product.DeleteAnEventProductHandler);

	// --------Event Booking Routes For Sellers---------//

	r.POST("/:eventId/Live", sale.LiveSaleHandler);
	// r.POST("/:eventId/Pause", sale.PauseSaleHandler);
	// r.POST("/:eventId/End", sale.StopSaleHandler);


	// --------Afer Event Routes For Sellers---------//

	r.GET("/booking/fetch/:eventId", );
	r.GET("/booking/analytics/:eventId", );
}