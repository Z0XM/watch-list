package error

import "fmt"

type Code string

func (e Code) ToString() string {
	return fmt.Sprintf("%s", e)
}

const (
	RequestNotValid          Code = "REQUEST_NOT_VALID"
	RequestInvalid           Code = "REQUEST_INVALID"
	SqlInsertError           Code = "SQL_INSERT_ERROR"
	SqlUpdateError           Code = "SQL_UPDATE_ERROR"
	SqlFetchError            Code = "SQL_FETCH_ERROR"
	BadRequestError          Code = "BAD_REQUEST"
	ApiUrlParsingError       Code = "API_URL_PARSING_ERROR"
	ApiRequestCreationError  Code = "API_REQUEST_CREATION_ERROR"
	ApiRequestError          Code = "API_REQUEST_ERROR"
	ApiRequestStatusError    Code = "API_REQUEST_STATUS_ERROR"
	JsonSerializationError   Code = "JSON_SERIALIZATION_ERROR"
	JsonDeserializationError Code = "JSON_DESERIALIZATION_ERROR"
	FormSerializationError   Code = "FORM_SERIALIZATION_ERROR"
	UserNotFoundInCtx        Code = "USER_NOT_FOUND_IN_CTX"
	UserTypeNotFoundInCtx    Code = "USER_TYPE_NOT_FOUND_IN_CTX"
	RulesNotFoundInCtx       Code = "RULES_NOT_FOUND_IN_CTX"
	PlanNotFoundInCtx        Code = "PLAN_NOT_FOUND_IN_CTX"
	PermissionsNotFoundInCtx Code = "PERMISSIONS_NOT_FOUND_IN_CTX"
)
