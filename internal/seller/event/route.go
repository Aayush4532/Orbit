package event

import "github.com/gin-gonic/gin"

func EventRoutes(r *gin.RouterGroup) {
	//---------Events Routes---------//
	r.POST("/create", CreateAnEventHandler);
	r.GET("/get", GetAllEventsHandler);
	r.GET("/get/:id", GetAnEventHandler);
	r.PUT("/update/:id", UpdateAnEventHandler);
	r.DELETE("/delete/:id", DeleteAnEventHandler);

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