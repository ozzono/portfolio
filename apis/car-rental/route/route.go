package route

import (
	"car-rental/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(handler *handler.Handler, debug bool) *gin.Engine {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("api/v1")

	vehicles := v1.Group("vehicles")
	vehicles.GET(":vehicle", handler.GetVehicle)
	vehicles.GET("", handler.GetVehicles)
	vehicles.POST("", handler.AddVehicle)
	vehicles.PUT(":vehicle", handler.UpdateVehicle)
	vehicles.DELETE(":vehicle", handler.DeleteVehicle)
	{
		vehicles.PUT(":vehicle/:user/schedule", handler.ScheduleVehicle)
		vehicles.PUT(":vehicle/:user/update", handler.PickupOrDropOffVehicle)
	}

	users := v1.Group("users")
	users.GET(":user", handler.GetUser)
	users.GET("", handler.GetUsers)
	users.POST("", handler.AddUser)
	users.PUT(":user", handler.UpdateUser)
	users.DELETE(":user", handler.DeleteUser)

	rents := v1.Group("rents")
	rents.GET(":rent", handler.GetRent)
	rents.GET("", handler.GetRents)
	rents.POST("", handler.AddRent)
	rents.PUT(":rent", handler.UpdateRent)
	rents.DELETE(":rent", handler.DeleteRent)

	return r
}
