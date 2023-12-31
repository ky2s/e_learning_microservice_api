package handler

import (
	"net/http"
	"transaction-service/helper"
	"transaction-service/orders"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService orders.Service
}

func NewOrderHandler(orderService orders.Service) *orderHandler {
	return &orderHandler{orderService}
}

func (h *orderHandler) CreateNewOrder(c *gin.Context) {
	var input orders.CreateOrderInput
	err := c.ShouldBind(&input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, err := h.orderService.Create(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Chapter", http.StatusOK, "Success", orders.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) Index(c *gin.Context) {

	ordersResult, err := h.orderService.Index()
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("success get data", http.StatusOK, "success", ordersResult)
	c.JSON(http.StatusOK, response)

}

func (h *orderHandler) Show(c *gin.Context) {

	var uri orders.InputUriID
	err := c.ShouldBindUri(&uri)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid Uri", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	println("uri.ID :: ", uri.ID)

	ordersResult, err := h.orderService.Show(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("success get data", http.StatusOK, "success", ordersResult)
	c.JSON(http.StatusOK, response)
}
