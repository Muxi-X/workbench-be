package errno

import err "muxi-workbench/pkg/err"

var (
	// Common errors
	OK                     = &err.Errno{Code: 0, Message: "OK"}
	ErrDatabase            = &err.Errno{Code: 10001, Message: "Database error"}
	ErrBind                = &err.Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrDuplicateStatusLike = &err.Errno{Code: 10003, Message: "The user had like this status."}
	ErrNoStatusLikeRecord  = &err.Errno{Code: 10004, Message: "It doesn't have any records for this user_status like "}
)
