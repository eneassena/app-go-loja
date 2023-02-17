package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/eneassena/app-go-loja/internal/products/domain"
	rules "github.com/eneassena/app-go-loja/pkg/regras"
	"github.com/eneassena/app-go-loja/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	service domain.ProductsService
}

func NewProductsController(service domain.ProductsService) *ProductsController {
	return &ProductsController{
		service: service,
	}
}

func (c *ProductsController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.FindAll()

		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}

func (c *ProductsController) FindByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// param := ctx.Request.Header.Get("id")
		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			web.NewContextResponse(ctx, 500, errors.New("invalid id"))
			return
		}

		products, err := c.service.FindByID(id)
		if err != nil {
			web.NewContextResponse(ctx, 500, err)
			return
		}
		web.NewContextResponse(ctx, http.StatusOK, products)
	}
}

func (c *ProductsController) Create() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data domain.ProductRequest
		ok := rules.ValidateErrorInRequest(context, &data)
		if ok {
			return
		}
		newProdcut, err := c.service.Create(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, web.DecodeError(http.StatusBadRequest, err.Error()))
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, newProdcut))
	}
}
