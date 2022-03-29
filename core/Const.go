package core

var (
	SchemeHttp  = "http"
	SchemeHttps = "https"

	MethodGet    = "GET"
	MethodPut    = "PUT"
	MethodPost   = "POST"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
	MethodHead   = "HEAD"
)

type RefreshStatus int

const (
	RefreshStatusRuning RefreshStatus = iota
	RefreshStatusSuccess
	RefreshStatusFailed
)
