package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yankie9486/TheRoofApp/backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		//Estimates route
		v1.GET("estimates", controllers.GetAllEstimates)
		v1.POST("estimates", controllers.CreateEstimate)
		v1.GET("estimates/:id", controllers.GetEstimate)
		v1.PUT("estimates/:id", controllers.UpdateEstimate)
		v1.DELETE("estimates/:id", controllers.DeleteEstimate)
	}
	return r
}
