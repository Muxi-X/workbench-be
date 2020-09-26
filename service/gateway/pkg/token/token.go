package token

import (
	"errors"

	"muxi-workbench/pkg/token"

	"github.com/gin-gonic/gin"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token.
type Context struct {
	ID uint32
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string) (*Context, error) {
	id, err := token.ResolveToken(tokenString)
	if err != nil {
		return &Context{}, err
	}

	return &Context{ID: id}, nil
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	return Parse(header)
}
