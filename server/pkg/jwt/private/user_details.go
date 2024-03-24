package private

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/projects/watch-list/server/constants"
	commonError "github.com/projects/watch-list/server/pkg/error"
	"log"
)

type JWTClaim struct {
	jwt.RegisteredClaims
	UserDetails UserDetails `json:"user_details"`
}

type UserDetails struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserTimeZone string `json:"user_time_zone"`
}

func GetUserID(ctx context.Context) (string, commonError.CustomError) {
	userDetails, ok := ctx.Value(constants.PrivateUserDetails).(*UserDetails)
	if !ok {
		return "", commonError.NewCustomError(commonError.UserNotFoundInCtx, "user details not found in ctx")
	}
	return userDetails.UserID, commonError.CustomError{}
}

func GetUserTimeZone(ctx context.Context) (string, commonError.CustomError) {
	userDetails, ok := ctx.Value(constants.PrivateUserDetails).(*UserDetails)
	if !ok {
		return "", commonError.NewCustomError(commonError.RulesNotFoundInCtx, "user timezone not found in ctx")
	}
	return userDetails.UserTimeZone, commonError.CustomError{}
}

func GetUserName(ctx context.Context) (string, commonError.CustomError) {
	userDetails, ok := ctx.Value(constants.PrivateUserDetails).(*UserDetails)
	if !ok {
		return "", commonError.NewCustomError(commonError.RulesNotFoundInCtx, "user name not found in ctx")
	}
	return userDetails.UserName, commonError.CustomError{}
}

func GetUserEmail(ctx context.Context) (string, commonError.CustomError) {
	userDetails, ok := ctx.Value(constants.PrivateUserDetails).(*UserDetails)
	if !ok {
		return "", commonError.NewCustomError(commonError.RulesNotFoundInCtx, "user email not found in ctx")
	}
	return userDetails.UserEmail, commonError.CustomError{}
}

func GetUserDetails(ctx context.Context) (UserDetails, commonError.CustomError) {
	userId, cusErr := GetUserID(ctx)
	if cusErr.Exists() {
		log.Printf("not able to fetch user id, error is %+v", cusErr.Error())
		return UserDetails{}, cusErr
	}

	userName, cusErr := GetUserName(ctx)
	if cusErr.Exists() {
		log.Printf("not able to fetch user name, error is %+v", cusErr.Error())
		return UserDetails{}, cusErr
	}

	userEmail, cusErr := GetUserEmail(ctx)
	if cusErr.Exists() {
		log.Printf("not able to fetch user email, error is %+v", cusErr.Error())
		return UserDetails{}, cusErr
	}

	timeZone, cusErr := GetUserTimeZone(ctx)
	if cusErr.Exists() {
		log.Printf("not able to fetch user timezone, error is %+v", cusErr.Error())
		return UserDetails{}, cusErr
	}

	return UserDetails{
		UserID:       userId,
		UserName:     userName,
		UserEmail:    userEmail,
		UserTimeZone: timeZone,
	}, commonError.CustomError{}
}
