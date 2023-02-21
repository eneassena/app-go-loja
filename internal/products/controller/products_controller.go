package controller

import (
	"errors"
	"log"
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

/**
 *
 */
func NewProductsController(service domain.ProductsService) *ProductsController {
	return &ProductsController{
		service: service,
	}
}

/**
 *
 */
func (p *ProductsController) Router(engine *gin.Engine) {
	engine.GET("/products", p.FindAll())
	engine.POST("/products", p.Create())
	engine.GET("/products/:id", p.FindByID())
	engine.GET("/products/name", p.FindByName())
	engine.DELETE("/products/:id", p.Remove())
	engine.PATCH("/products/:id/:count", p.UpdateCount())
}

/**
 *
 */
func (c *ProductsController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.FindAll()

		if err != nil {
			web.NewContextResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		web.NewContextResponse(ctx, http.StatusOK, products)
	}
}

/**
 *
 */
func (c *ProductsController) FindByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := c.getParamId(ctx)

		products, err := c.service.FindByID(param)
		if err != nil {
			web.NewContextResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		web.NewContextResponse(ctx, http.StatusOK, products)
	}
}

/**
 *
 */
func (c *ProductsController) Create() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data domain.ProductRequest

		if ok := rules.ValidateErrorInRequest(context, &data); ok {
			return
		}

		newProdcut, err := c.service.Create(data)
		if err != nil {
			web.NewContextResponse(context, http.StatusInternalServerError, err.Error())
			return
		}
		web.NewContextResponse(context, http.StatusOK, newProdcut)
	}
}

/**
 *
 */
func (c *ProductsController) FindByName() gin.HandlerFunc {
	return func(context *gin.Context) {
		name := context.Query("name")

		if name == "" {
			erro := errors.New("invalid name param")
			web.NewContextResponse(context, http.StatusInternalServerError, erro.Error())
			return
		}

		product, erro := c.service.FindByName(name)
		if erro != nil {
			web.NewContextResponse(context, http.StatusNotFound, erro.Error())
			return
		}

		web.NewContextResponse(context, http.StatusOK, product)
	}
}

/**
 *
 */
func (c *ProductsController) Remove() gin.HandlerFunc {
	return func(context *gin.Context) {
		param := context.Param("id")
		log.Println(param)
		paramID, erro := strconv.ParseInt(param, 10, 64)
		if erro != nil {
			web.NewContextResponse(context, http.StatusInternalServerError, erro.Error())
			return
		}

		erro = c.service.Remove(domain.ProductRequest{ID: int(paramID)})
		if erro != nil {
			web.NewContextResponse(context, http.StatusNotFound, erro.Error())
			return
		}

		web.NewContextResponse(context, http.StatusNoContent, nil)
	}
}

/**
 *
 */
func (c *ProductsController) UpdateCount() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramCount := context.Param("count")
		paramID := context.Param("id")

		count, erro := strconv.ParseInt(paramCount, 10, 64)
		if erro != nil {
			web.NewContextResponse(context, http.StatusInternalServerError, erro.Error())
			return
		}

		id, erro := strconv.ParseInt(paramID, 10, 64)
		if erro != nil {
			web.NewContextResponse(context, http.StatusInternalServerError, erro.Error())
			return
		}

		product := domain.ProductRequest{ID: int(id), Count: int(count)}

		erro = c.service.UpdateCount(product)
		if erro != nil {
			web.NewContextResponse(context, http.StatusInternalServerError, erro.Error())
			return
		}

		web.NewContextResponse(context, http.StatusNoContent, nil)
	}
}

/**
 *
 */
func (productController *ProductsController) getParamId(c *gin.Context) int {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		web.NewContextResponse(c, 500, errors.New("invalid id"))
		return 0
	}
	return id
}
