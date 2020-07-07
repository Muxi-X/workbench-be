package errno

import (
	err "github.com/Muxi-X/workbench-be/pkg/err"
)

var (
	OK  = &err.Errno{Code: 0, Message: "OK"}

	ErrDatabase = &err.Errno{Code: 10001, Message: "Database error"}

    //group
	ErrPermissionDenied = &err.Errno{Code: 20101, Message: "Permission Denied"}
	)


