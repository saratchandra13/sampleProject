package enterr

import "fmt"

type CustomCode string

type CustomError struct {
	Code CustomCode
	Msg string
	Err error
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("CustomError: %s, %s", ce.Code, ce.Msg)
}

func NewCustomError(code CustomCode, msg string, err error) *CustomError {
	return &CustomError{
		Code: code,
		Msg: msg,
		Err: err,
	}
}

// custom codes
const (
	AddReviewFailed = CustomCode("errAddReviewFailed")
	FetchReviewFailed = CustomCode("errFetchReviewFailed")
)
