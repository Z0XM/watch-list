package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/projects/watch-list/server/constants"
	"github.com/projects/watch-list/server/core/commons"
	"github.com/projects/watch-list/server/core/domain/interfaces"
	"github.com/projects/watch-list/server/pkg/jwt/private"
	"net/http"
	"sync"
)

type UserController struct {
	svc interfaces.UserService
}

var ctrl *UserController
var ctrlOnce sync.Once

func NewUserController(svc interfaces.UserService) *UserController {
	ctrlOnce.Do(func() {
		ctrl = &UserController{
			svc: svc,
		}
	})
	return ctrl
}

func (ctrl *UserController) GetMyUser(ctx *gin.Context) {
	userDetails, cusErr := private.GetUserDetails(ctx)
	if cusErr.Exists() {
		commons.NewErrorResponse(ctx, cusErr, http.StatusBadRequest)
		return
	}

	response, cusErr := ctrl.svc.GetUser(ctx, userDetails.UserID)
	if cusErr.Exists() {
		commons.NewErrorResponse(ctx, cusErr, http.StatusBadRequest)
		return
	}

	commons.NewSuccessResponse(ctx, response)
	return
}

func (ctrl *UserController) GetUser(ctx *gin.Context) {
	userID := ctx.Param(constants.UserID)

	response, cusErr := ctrl.svc.GetUser(ctx, userID)
	if cusErr.Exists() {
		commons.NewErrorResponse(ctx, cusErr, http.StatusBadRequest)
		return
	}

	commons.NewSuccessResponse(ctx, response)
	return
}

func (ctrl *UserController) GetAllUsers(ctx *gin.Context) {
	response, cusErr := ctrl.svc.GetAllUsers(ctx)
	if cusErr.Exists() {
		commons.NewErrorResponse(ctx, cusErr, http.StatusBadRequest)
		return
	}

	commons.NewSuccessResponse(ctx, response)
	return
}
