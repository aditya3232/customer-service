package controllers

import (
	errWrap "customer-service/common/error"
	"customer-service/common/response"
	"customer-service/domain/dto"
	"customer-service/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerController struct {
	service services.IServiceRegistry
}

type ICustomerController interface {
	FindByID(*gin.Context)
	Create(*gin.Context)
	FindAllWithoutPagination(*gin.Context)
}

func NewCustomerController(service services.IServiceRegistry) ICustomerController {
	return &CustomerController{service: service}
}

func (c *CustomerController) FindByID(ctx *gin.Context) {
	reqID, _ := strconv.Atoi(ctx.Param("id"))
	customer, err := c.service.GetCustomer().FindByID(ctx.Request.Context(), reqID)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: customer,
		Gin:  ctx,
	})
}

func (c *CustomerController) Create(ctx *gin.Context) {
	request := &dto.CustomerRequest{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	customer, err := c.service.GetCustomer().Create(ctx, request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: customer,
		Gin:  ctx,
	})
}

func (c *CustomerController) FindAllWithoutPagination(ctx *gin.Context) {
	customers, err := c.service.GetCustomer().FindAllWithoutPagination(ctx)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: customers,
		Gin:  ctx,
	})
}
