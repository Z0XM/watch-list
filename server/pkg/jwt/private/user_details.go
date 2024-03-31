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
	Sub   string `json:"sub"`
	Email string `json:"email"`
}

type UserDetails struct {
	UserID string `json:"sub"`
	Email  string `json:"email"`
}

func GetUserID(ctx context.Context) (string, commonError.CustomError) {
	userDetails, ok := ctx.Value(constants.PrivateUserDetails).(*UserDetails)
	if !ok {
		return "", commonError.NewCustomError(commonError.UserNotFoundInCtx, "user id not found in ctx")
	}
	return userDetails.UserID, commonError.CustomError{}
}

func GetEmail(ctx context.Context) (string, commonError.CustomError) {
	userDetails, ok := ctx.Value(constants.PrivateUserDetails).(*UserDetails)
	if !ok {
		return "", commonError.NewCustomError(commonError.EmailNotFoundInCtx, "email not found in ctx")
	}
	return userDetails.Email, commonError.CustomError{}
}

func GetUserDetails(ctx context.Context) (UserDetails, commonError.CustomError) {
	userId, cusErr := GetUserID(ctx)
	if cusErr.Exists() {
		log.Printf("not able to fetch sub, error is %+v\n", cusErr.Error())
		return UserDetails{}, cusErr
	}

	userEmail, cusErr := GetEmail(ctx)
	if cusErr.Exists() {
		log.Printf("not able to fetch email, error is %+v\n", cusErr.Error())
		return UserDetails{}, cusErr
	}

	return UserDetails{
		UserID: userId,
		Email:  userEmail,
	}, commonError.CustomError{}
}
