package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	return gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func Response(data interface{}) gin.H {
	return gin.H{
		"data":  data,
		"error": nil,
	}
}

func GetErrorCode(err error) int {
	switch err {
	case ErrImageRequired:
		return http.StatusBadRequest
	case ErrForbidden:
		return http.StatusForbidden
	case ErrNotFound, ErrDataNotFound:
		return http.StatusNotFound
	case ErrDataIsEmpty:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func ErrIsBadRequest(err error) bool {
	switch err {
	case ErrInvalidID:
		return true
	case ErrInvalidExtension:
		return true
	case ErrTypeNotValid:
		return true
	case ErrMonthNotValid:
		return true
	case ErrDateDailyNotValid:
		return true
	case ErrDateWeeklyNotValid:
		return true
	default:
		if strings.Contains(err.Error(), ErrInvalidID.Error()) {
			return true
		}
		return false
	}
}
