package common

import (
	"fmt"
	"strconv"
)

type CustomError map[string]string

func (e CustomError) Error() string {
	return fmt.Sprintf("%s(%v)", e["name"], e)
}

func (e CustomError) HTTPError() HTTPError {
	httpStatus, err := strconv.Atoi(getOrDefault(e, "http_status", "500"))
	if err != nil {
		httpStatus = 500
	}
	httpCode := getOrDefault(e, "http_code", "")
	httpMessage := getOrDefault(e, "http_message", "")
	return HTTPError{httpStatus, httpCode, httpMessage}
}

func getOrDefault(m map[string]string, key string, dflt string) string {
	value, ok := m[key]
	if ok {
		return value
	} else {
		return dflt
	}
}
