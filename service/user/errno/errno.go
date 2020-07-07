package errno

import err "muxi-workbench/pkg/err"

var (
	// Common errors
	OK             = &err.Errno{Code: 0, Message: "OK"}
	ErrDatabase    = &err.Errno{Code: 10001, Message: "Database error"}
	ErrBind        = &err.Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrAccessToken = &err.Errno{Code: 10003, Message: "Error occurred while getting access token from oauth2 server"}
	ErrAuthToken   = &err.Errno{Code: 10004, Message: "Error occurred while handling the auth token"}
	ErrBadRequest  = &err.Errno{Code: 10005, Message: "Request error"}
)
