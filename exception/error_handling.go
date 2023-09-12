package exception

import (
	"Kelompok-2/dompet-online/model/dto/resp"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandling(c *gin.Context, err error) {
	if ErrNotFound(c, err) {
		return
	}

	if ErrValidation(c, err) {
		return
	}

	ErrInternalServer(c, err)
}

func ErrNotFound(c *gin.Context, err error) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")

		apiResponse := resp.ErrorHandlingResponse{
			Status:  http.StatusNotFound,
			Message: "not found",
			Data:    exception.Error(),
		}
		c.JSON(apiResponse.Status, apiResponse)
		return true
	}
	return false
}

func ErrValidation(c *gin.Context, err error) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")

		apiResponse := resp.ErrorHandlingResponse{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Data:    exception.Error(),
		}
		c.JSON(apiResponse.Status, apiResponse)
		return true
	}
	return false

}

func ErrInternalServer(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	apiResponse := resp.ErrorHandlingResponse{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
		Data:    err.Error(),
	}
	c.JSON(apiResponse.Status, apiResponse)
}
