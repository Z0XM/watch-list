package error

import (
	"errors"
	"fmt"
	errorsPkg "github.com/pkg/errors"
	"strings"
)

type CustomError struct {
	errorCode     Code
	errorMsg      string
	error         error
	exists        bool
	retryable     bool
	shouldNotify  bool
	loggingParams map[string]interface{}
}

func NewCustomError(errorCode Code, error string, options ...func(*CustomError)) CustomError {
	c := CustomError{
		errorCode:    errorCode,
		errorMsg:     error,
		exists:       true,
		retryable:    false,
		shouldNotify: true,
	}

	for _, option := range options {
		option(&c)
	}

	e := errors.New(fmt.Sprintf("Code: %s | %s", c.errorCode, c.errorMsg))
	c.error = errorsPkg.WithStack(e)
	c.loggingParams = make(map[string]interface{}, 0)
	return c
}

func WithRetryable(retryable bool) func(*CustomError) {
	return func(c *CustomError) {
		c.retryable = retryable
	}
}

func WithShouldNotify(shouldNotify bool) func(*CustomError) {
	return func(c *CustomError) {
		c.shouldNotify = shouldNotify
	}
}

func RequestInvalidError(message string, options ...func(*CustomError)) CustomError {
	c := CustomError{
		errorCode:    RequestInvalid,
		errorMsg:     message,
		exists:       true,
		retryable:    false,
		shouldNotify: true,
	}
	e := errors.New(fmt.Sprintf("Code: %s | %s", c.errorCode, c.errorMsg))
	c.error = errorsPkg.WithStack(e)
	c.loggingParams = make(map[string]interface{}, 0)

	for _, option := range options {
		option(&c)
	}
	return c
}

func (c CustomError) Exists() bool {
	return c.exists
}

func (c CustomError) Log() {
	fmt.Println(c.ToString())
}

func (c CustomError) LoggingParams() map[string]interface{} {
	return c.loggingParams
}

func (c CustomError) ErrorCode() Code {
	return c.errorCode
}

func (c CustomError) ToError() error {
	return c.error
}

func (c CustomError) Error() string {
	return c.error.Error()
}

func (c CustomError) ErrorMessage() string {
	return c.errorMsg
}

func (c CustomError) ShouldNotify() bool {
	return c.shouldNotify
}

func (c CustomError) Retryable() bool {
	return c.retryable
}

func (c CustomError) ToString() string {
	logMsg := fmt.Sprintf("Code: %s, Msg: %s", c.errorCode, c.errorMsg)

	paramStrings := make([]string, 0)
	for key, val := range c.loggingParams {
		paramStrings = append(paramStrings, fmt.Sprintf("%s: {%+v}", strings.ToUpper(key), val))
	}
	return fmt.Sprintf("%s, Params: [%+v]", logMsg, strings.Join(paramStrings, " | "))
}

// WithParam value param should not be a pointer
func (c CustomError) WithParam(key string, val interface{}) CustomError {
	if c.loggingParams == nil {
		c.loggingParams = make(map[string]interface{}, 0)
	}
	c.loggingParams[key] = val
	return c
}

func (c CustomError) ErrorString() string {
	return c.errorMsg
}

func (c CustomError) UserMessage() string {
	return c.errorMsg
}
