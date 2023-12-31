package main

import (
	"github.com/Ihpaz/golang-restapi-userfamily/config"
	"github.com/Ihpaz/golang-restapi-userfamily/controller"
	"github.com/Ihpaz/golang-restapi-userfamily/middlewares"
	"github.com/Ihpaz/golang-restapi-userfamily/repository"
	"github.com/Ihpaz/golang-restapi-userfamily/routes"
	"github.com/Ihpaz/golang-restapi-userfamily/service"
)

var (
	DB = config.Init()

	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(DB)
	customerService    service.CustomerService       = service.NewCustomerService(customerRepository)
	customerController controller.CustomerController = controller.NewCustomerController(customerService)

	nationalityRepository repository.NationalityRepository = repository.NewNationalityRepository(DB)
	nationalityService    service.NationalityService       = service.NewNationalityService(nationalityRepository)
	nationalityController controller.NationalityController = controller.NewNationalityController(nationalityService)

	httpRouter routes.Router = routes.NewMuxRouter()
)

func main() {

	const port string = ":8080"
	httpRouter.GET("/customer", middlewares.SetMiddlewareJSON(customerController.GetCustomers))
	httpRouter.POST("/customer", middlewares.SetMiddlewareJSON(customerController.AddCustomer))
	httpRouter.PUT("/customer/{id}", middlewares.SetMiddlewareJSON(customerController.UpdateCustomer))
	httpRouter.GET("/customer/{id}", middlewares.SetMiddlewareJSON(customerController.GetCustomer))
	httpRouter.DELETE("/customer/{id}", middlewares.SetMiddlewareJSON(customerController.DeleteCustomer))

	httpRouter.GET("/nationality", middlewares.SetMiddlewareJSON(nationalityController.GetNationalities))
	httpRouter.POST("/nationality", middlewares.SetMiddlewareJSON(nationalityController.AddNationality))

	httpRouter.SERVE(port)

}
