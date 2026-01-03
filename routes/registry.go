package routes

import (
	"customer-service/controllers"
	routes "customer-service/routes/customer"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IRouteRegister interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) IRouteRegister {
	return &Registry{controller: controller, group: group}
}

func (r *Registry) customerRoute() routes.ICustomerRoute {
	return routes.NewCustomerRoute(r.controller, r.group)
}

func (r *Registry) Serve() {
	r.customerRoute().Run()
}
