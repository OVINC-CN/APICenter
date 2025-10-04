package home

import (
	"net/http"

	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
)

var MySQLConnFailed = &apiMixin.APIError{
	Code: http.StatusInternalServerError,
	Msg:  "database connection failed",
}

var RedisConnFailed = &apiMixin.APIError{
	Code: http.StatusInternalServerError,
	Msg:  "redis connection failed",
}
