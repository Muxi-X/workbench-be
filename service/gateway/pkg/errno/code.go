package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrQuery            = &Errno{Code: 10003, Message: "Error occurred while getting url queries."}
	ErrPathParam        = &Errno{Code: 10004, Message: "Error occurred while getting path param."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}

	// auth errors
	ErrTokenInvalid     = &Errno{Code: 20101, Message: "The token was invalid."}
	ErrPermissionDenied = &Errno{Code: 20102, Message: "Permission denied."}
	ErrNotJoined        = &Errno{Code: 20103, Message: "User not join any team."}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20201, Message: "The user was not found."}
	ErrPasswordIncorrect = &Errno{Code: 20202, Message: "The password was incorrect."}

	// feed errors
	ErrFeedList = &Errno{Code: 20301, Message: "Error occurred while getting feed list."}

	// attention errors
	ErrAttentionList = &Errno{Code: 20401, Message: "Error occurred while getting attention list."}

	// project errors
	ErrTrashbinType            = &Errno{Code: 20501, Message: "Invalid trashbin type."}
	ErrNoProjectId             = &Errno{Code: 20502, Message: "Project service must have project_id in query."}
	ErrProjectPermissionDenied = &Errno{Code: 20503, Message: "Permission denied or this project has been deleted."}

	// status errors
	// ...

	// upload errors
	ErrGetFile    = &Errno{Code: 20701, Message: "Error occurred in getting file from FormFile()"}
	ErrUploadFile = &Errno{Code: 20702, Message: "Error occurred in uploading file to oss"}

	// team errors
	// ...

)
