package errno

import err "muxi-workbench/pkg/err"

var (
	// Common errors
	OK             = &err.Errno{Code: 0, Message: "OK"}
	ErrDatabase    = &err.Errno{Code: 10001, Message: "Database error"}
	ErrBind        = &err.Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrBadRequest  = &err.Errno{Code: 10003, Message: "Request error"}
	ErrUserExisted = &err.Errno{Code: 10004, Message: "User has existed"}
	ErrAuthToken   = &err.Errno{Code: 10005, Message: "Error occurred while handling the auth token"}

	// oauth errors
	ErrRegister          = &err.Errno{Code: 20001, Message: "Error occurred while registering on auth-server"}
	ErrRemoteAccessToken = &err.Errno{Code: 20002, Message: "Error occurred while getting oauth access token from auth-server"}
	ErrLocalAccessToken  = &err.Errno{Code: 20003, Message: "Error occurred while getting oauth access token from local"}
	ErrGetUserInfo       = &err.Errno{Code: 20004, Message: "Error occurred while getting user info from oauth-server by access token"}
)
