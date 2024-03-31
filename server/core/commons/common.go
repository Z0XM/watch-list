package commons

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/projects/watch-list/server/constants"
	error2 "github.com/projects/watch-list/server/pkg/error"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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

func NewErrorResponse(ctx *gin.Context, error error2.CustomError, code int) {
	log.Println(error.Error(), error.Error(), GetRequestID(ctx))

	res := &ErrorResponse{
		IsSuccess:  false,
		StatusCode: code,
		Error: Error{
			Message: error.Error(),
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

type StatusMessage struct {
	Message string `json:"message,omitempty"`
}

func SuccessMessage(message string) StatusMessage {
	return StatusMessage{
		Message: message,
	}
}

func GetLogTag(ctx context.Context, funcName string) string {
	return fmt.Sprintf(" Request ID: %s Function: %s ", GetRequestID(ctx), funcName)
}

func Paginate(c context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		ctx := c.(*gin.Context)

		pageNumber, err := strconv.Atoi(ctx.Query(constants.PageNumber))
		if pageNumber <= 0 || err != nil {
			pageNumber = constants.DefaultPageNumber
		}

		pageSize, err := strconv.Atoi(ctx.Query(constants.PageSize))
		if pageSize <= 0 || err != nil {
			pageSize = constants.DefaultPageSize
		}

		sortCondition := ctx.Query(constants.Sort)
		if len(sortCondition) == 0 {
			sortCondition = constants.Ascending
		}

		return db.Offset(GetPaginationOffset(pageNumber, pageSize)).Limit(pageSize).Order("id " + sortCondition)
	}
}

func GetPaginationOffset(pageNumber int, pageSize int) int {
	return (pageNumber - 1) * pageSize
}
