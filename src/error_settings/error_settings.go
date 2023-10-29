package error_settings

import (
	"net/http"

	"github.com/pkg/errors"
)

// ref: https://qiita.com/immrshc/items/13199f420ebaf0f0c37c#%E3%82%A8%E3%83%A9%E3%83%BC%E3%81%AB%E5%BF%9C%E3%81%98%E3%81%A6%E3%82%B9%E3%83%86%E3%83%BC%E3%82%BF%E3%82%B9%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E9%81%B8%E3%81%B6

// ErrorType エラーの種類
type ErrorType uint

const (
	Unknown ErrorType = iota
	BadRequest
	Unauthenticated
	Forbidden
	NotFound
)

// ErrorTypeを返すインターフェース
type typeGetter interface {
	Type() ErrorType
}

// ErrorTypeを持つ構造体
type customError struct {
	errorType     ErrorType
	originalError error
}

// New 指定したErrorTypeを持つcustomErrorを返す
func (et ErrorType) New(message string) error {
	return customError{errorType: et, originalError: errors.New(message)}
}

// Wrap 指定したErrorTypeと与えられたメッセージを持つcustomErrorにWrapする
func (et ErrorType) Wrap(err error, message string) error {
	return customError{errorType: et, originalError: errors.Wrap(err, message)}
}

// Error errorインターフェースを実装する
func (e customError) Error() string {
	return e.originalError.Error()
}

// Type typeGetterインターフェースを実装する
func (e customError) Type() ErrorType {
	return e.errorType
}

// Wrap 受け取ったerrorがErrorTypeを持つ場合はそれを引き継いで与えられたエラーメッセージを持つcustomErrorにWrapする
func Wrap(err error, message string) error {
	we := errors.Wrap(err, message)
	if ce, ok := err.(typeGetter); ok {
		return customError{errorType: ce.Type(), originalError: we}
	}
	return customError{errorType: Unknown, originalError: we}
}

// Cause errors.CauseのWrapper
func Cause(err error) error {
	return errors.Cause(err)
}

// GetType ErrorTypeを持つ場合はそれを返し、無ければUnknownを返す
func GetType(err error) ErrorType {
	if e, ok := err.(typeGetter); ok {
		return e.Type()
	} else {
		return Unknown
	}
}

func StatusCode(err error) int {
	switch GetType(err) {
	case BadRequest:
		return http.StatusBadRequest // 400
	case Unauthenticated:
		return http.StatusUnauthorized // 401
	case Forbidden:
		return http.StatusForbidden // 403
	case NotFound:
		return http.StatusNotFound // 404
	default:
		return http.StatusInternalServerError // 500
	}
}
