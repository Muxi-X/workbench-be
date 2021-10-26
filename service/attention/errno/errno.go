package errno

import "muxi-workbench/pkg/err"

var (
	// Common errors
	OK                = &err.Errno{Code: 0, Message: "OK"}
	ErrDatabase       = &err.Errno{Code: 10001, Message: "Database error"}
	ErrBind           = &err.Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrFormatList     = &err.Errno{Code: 10003, Message: "Error occurred while format data list"}
	ErrPublishMsg     = &err.Errno{Code: 10004, Message: "Error occurred while publishing message"}
	ErrJsonMarshal    = &err.Errno{Code: 10005, Message: "Error occurred while marshaling json"}
	ErrGetDataFromRPC = &err.Errno{Code: 10006, Message: "Error occurred in rpc"}
)
