package main

import (
	"github.com/eneassena/app-go-loja/internal/products/controller"
	connect "github.com/eneassena/app-go-loja/internal/products/infra"
	"github.com/eneassena/app-go-loja/internal/products/repository"
	"github.com/eneassena/app-go-loja/internal/products/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := connect.Connect()

	repositoryProduct := repository.NewProductsRepository(db)
	serviceProduct := service.NewProductsService(repositoryProduct)
	controllerProduct := controller.NewProductsController(serviceProduct)

	router := gin.Default()

	controllerProduct.Router(router)

	router.Run(":8080")

}
