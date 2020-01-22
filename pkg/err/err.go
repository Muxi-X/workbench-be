package err

import (
	"encoding/json"
	"net/http"

	errors "github.com/micro/go-micro/errors"
)

type ErrDetail struct {
	Errno
	Cause string `json:"cause"`
}

func BadRequestErr(errno *Errno, cause string) error {
	detail := &ErrDetail{
		*errno,
		cause,
	}

	detailStr, _ := json.Marshal(detail)

	return &errors.Error{
		Id:     "serverName",
		Code:   409,
		Detail: string(detailStr),
		Status: http.StatusText(400),
	}
}

func ServerErr(errno *Errno, cause string) error {
	detail := &ErrDetail{
		*errno,
		cause,
	}

	detailStr, _ := json.Marshal(detail)

	return &errors.Error{
		Id:     "serverName",
		Code:   500,
		Detail: string(detailStr),
		Status: http.StatusText(500),
	}
}
