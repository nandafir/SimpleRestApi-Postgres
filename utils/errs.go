package utils

import "errors"

var (
	// sql error
	ErrDuplicateKey = errors.New("duplicate key")

	// common
	ErrForbidden    = errors.New("ErrForbidden")
	ErrUnauthorized = errors.New("ErrUnauthorized")

	// validate
	ErrIsRequired         = errors.New("is required")
	ErrPasswordDoestMatch = errors.New("ErrPasswordDoestMatch")
	ErrInvalidTimeBasis   = errors.New("ErrInvalidTimeBasis")
	ErrInvalidID          = errors.New("ErrInvalidID")
	ErrInvalidParams      = errors.New("ErrInvalidParams")
	ErrAtLeastZero        = errors.New("must at least 0")
	ErrMustBeFilled       = errors.New("must be filled")

	// Time
	ErrUploadTimeExeeded = errors.New("ErrUploadTimeExeeded")
	ErrDateCannotBeBlank = errors.New("ErrDateCannotBeBlank")
	ErrInvalidSchedule   = errors.New("ErrInvalidSchedule")
	ErrInvalidDate       = errors.New("ErrInvalidDate")
	ErrDateRange         = errors.New("date-from cannot be greater than date-to")

	// File
	ErrInvalidExtension  = errors.New("ErrInvalidExtension")
	ErrCsvFileNotValid   = errors.New("ErrCsvFileNotValid")
	ErrExcelFileNotValid = errors.New("ErrExcelFileNotValid")
	ErrFileMustPDF       = errors.New("ErrFileMustPDF")

	// not found
	ErrNotFound          = errors.New("ErrNotFound")
	ErrDataNotFound      = errors.New("data not found")
	ErrUserNotFound      = errors.New("ErrUserNotFound")
	ErrDataIsEmpty       = errors.New("ErrDataIsEmpty")
	ErrTokenNotExist     = errors.New("ErrTokenNotExist")
	ErrUserEmailNotFound = errors.New("ErrUserEmailNotFound")

	// required
	ErrImageRequired    = errors.New("ErrImageRequired")
	ErrFileRequired     = errors.New("ErrFileRequired")
	ErrIDRequired       = errors.New("ErrIDRequired")
	ErrCacheKeyRequired = errors.New("ErrCacheKeyRequired")
	ErrUsernameRequired = errors.New("ErrUsernameRequired")
	ErrUploadMapImage   = errors.New("ErrImageNeeded")
	ErrDateNeeded       = errors.New("ErrDateNeeded")
	ErrMonthNeeded      = errors.New("ErrMonthNeeded")
	ErrTypeNeeded       = errors.New("ErrTypeNeeded")
	ErrHourNeeded       = errors.New("ErrHourNeeded")
	ErrYearNeeded       = errors.New("ErrYearNeeded")
	ErrPeriodIsRequired = errors.New("ErrPeriodNeeded")

	ErrTypeNotValid       = errors.New("TypeNotValid")
	ErrMonthNotValid      = errors.New("MonthNotValid")
	ErrDateDailyNotValid  = errors.New("ErrorDateDailyNotValid")
	ErrDateWeeklyNotValid = errors.New("ErrorDateWeeklyNotValid")
	ErrFieldIsWrongResult = errors.New("ErrorFieldIsWrong")
)
