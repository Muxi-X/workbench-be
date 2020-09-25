package errno

import (
	"muxi-workbench/pkg/err"
)

var (
	OK = &err.Errno{Code: 0, Message: "OK"}

	ErrDatabase = &err.Errno{Code: 10001, Message: "Database error"}
	ErrClient   = &err.Errno{Code: 10002, Message: "Client error"}

	// group errors
	ErrPermissionDenied = &err.Errno{Code: 20101, Message: "Permission Denied"}

	// invitation errors
	ErrLinkExpiration = &err.Errno{Code: 20201, Message: "Link expiration"}
)
