package auth

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
	ID     uint32
	Role   uint32
	TeamID uint32
}

// Parse parses the token, and returns the context if the token is valid.
func Parse(tokenString string) (*Context, error) {
	t, err := token.ResolveToken(tokenString)
	if err != nil {
		return nil, err
	}

	return &Context{
		ID:     t.ID,
		Role:   t.Role,
		TeamID: t.TeamID,
	}, nil
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return nil, ErrMissingHeader
	}

	return Parse(header)
}
