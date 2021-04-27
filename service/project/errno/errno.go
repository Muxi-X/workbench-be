package errno

import err "muxi-workbench/pkg/err"

var (
	// Common errors
	OK                  = &err.Errno{Code: 0, Message: "OK"}
	ErrDatabase         = &err.Errno{Code: 10001, Message: "Database error"}
	ErrBind             = &err.Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrPermissionDenied = &err.Errno{Code: 10003, Message: "Permission denied."}
	ErrNotFound         = &err.Errno{Code: 10004, Message: "Page not found."}

	// delete children
	ErrFileNotFound    = &err.Errno{Code: 10005, Message: "File not found or had been removed."}
	ErrInvalidIndex    = &err.Errno{Code: 10006, Message: "Invalid file position index."}
	ErrInvalidFileType = &err.Errno{Code: 10007, Message: "Invalid file type."}
)
