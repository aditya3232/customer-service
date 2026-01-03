package routes

import (
	"customer-service/controllers"

	"github.com/gin-gonic/gin"
)

type CustomerRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type ICustomerRoute interface {
	Run()
}

func NewCustomerRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup) ICustomerRoute {
	return &CustomerRoute{controller: controller, group: group}
}

func (r *CustomerRoute) Run() {
	group := r.group.Group("/customers")
	group.GET("", r.controller.GetCustomer().FindAllWithoutPagination)
	group.GET("/:id", r.controller.GetCustomer().FindByID)
	group.POST("", r.controller.GetCustomer().Create)
}
