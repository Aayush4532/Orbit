package event

import "github.com/gin-gonic/gin"

func EventRoutes(r *gin.RouterGroup) {
	//---------Events Routes---------//
	r.POST("/create", );
	r.GET("/get", );
	r.GET("/get/:id", );
	r.PUT("/update/:id", );
	r.DELETE("/delete/:id", );

	// --------Event Product Routes---------//
	r.POST("/:eventId/registerProducts", );
	r.GET("/:eventId/getProducts", );
	r.GET("/:eventId/getProduct/:productId", );
	r.PUT("/:eventId/updateProduct/:productId", );
	r.DELETE("/:eventId/deleteProduct/:productId", );

	// --------Event Booking Routes For Sellers---------//

	r.POST("/:eventId/Live/:id", );
	r.POST("/:eventId/Pause/:id", );
	r.POST("/:eventId/End/:id", );


	// --------Afer Event Routes For Sellers---------//

	r.GET("/booking/fetch/:eventId", );
	r.GET("/booking/analytics/:eventId", );
}