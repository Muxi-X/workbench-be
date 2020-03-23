package errno

import "muxi-workbench/pkg/err"

var (
	// Common errors
	OK          = &err.Errno{Code: 0, Message: "OK"}
	ErrDatabase = &err.Errno{Code: 10001, Message: "Database error"}
	ErrBind     = &err.Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
)
