package handler

import (
	"net/http"
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
	var input chapters.CreateOrderInput
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

	course, err := h.CoursesService.FindCourseByID(input.CourseID)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if course.ID == 0 {
		response := helper.ApiResponse("Course doesnt exist", http.StatusBadRequest, "error", gin.H{})

		c.JSON(http.StatusBadRequest, response)
		return
	}

	chapter, err := h.chapterService.Create(input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Chapter", http.StatusOK, "Success", chapters.FormatShowChapters(chapter))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) Index(c *gin.Context) {

}

func (h *orderHandler) Show(c *gin.Context) {

}
