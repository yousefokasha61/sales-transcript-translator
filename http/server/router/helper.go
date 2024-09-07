package route

import (
	"chat/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success = "success"
)

type emptyData struct{}

func (r *Router) newSuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, model.HttpMainResponse{
		Status:  http.StatusOK,
		Message: Success,
		Data:    data,
	})
}

func (r *Router) newBadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, model.HttpMainResponse{
		Status:  http.StatusBadRequest,
		Message: message,
		Data:    emptyData{},
	})
}
