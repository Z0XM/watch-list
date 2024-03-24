package commons

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/projects/watch-list/server/constants"
	"log"
	"net/http"
)

type ErrorResponse struct {
	IsSuccess  bool  `json:"is_success"`
	StatusCode int   `json:"status_code"`
	Error      Error `json:"error"`
}

type SuccessResponse struct {
	IsSuccess  bool        `json:"is_success"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta"`
}

type Error struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// GetRequestID returns the request ID from context
func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(constants.RequestID).(string)
	if !ok {
		return uuid.New().String()
	}

	return requestID
}

func NewErrorResponse(ctx *gin.Context, msg string, error string, code int) {
	log.Println(msg, error, GetRequestID(ctx))

	res := &ErrorResponse{
		IsSuccess:  false,
		StatusCode: code,
		Error: Error{
			Message: error,
		},
	}

	ctx.AbortWithStatusJSON(code, res)
}

func NewSuccessResponse(ctx *gin.Context, data interface{}) {
	res := &SuccessResponse{
		IsSuccess:  true,
		StatusCode: http.StatusOK,
		Data:       data,
	}

	ctx.AbortWithStatusJSON(http.StatusOK, res)
}
