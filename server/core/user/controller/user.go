package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/projects/watch-list/server/core/domain/interfaces"
	"github.com/projects/watch-list/server/core/user/controller/requests"
	"github.com/projects/watch-list/server/pkg/validator"
	"log"
	"net/http"
	"sync"
)

type Controller struct {
	service interfaces.UserService
}

var ctrl *Controller
var ctrlOnce sync.Once

func NewController(svc interfaces.UserService) *Controller {
	ctrlOnce.Do(func() {
		ctrl = &Controller{
			service: svc,
		}
	})
	return ctrl
}

func (uc *Controller) Login(ctx *gin.Context) {
	var request requests.LoginRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		log.Println("Invalid request", err.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	err = validator.Get().Struct(&request)
	if err != nil {
		log.Println("Invalid request", err.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	response, err := uc.service.Login(ctx, request.ToSvcRequest())
	if err != nil {
		log.Println("Invalid request", err.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message":  "Login Success",
		"Response": response,
	})

	return
}
