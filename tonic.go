package tonic

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	// SessionVerifier - Allows us to verify the session
	SessionVerifier interface {
		VerifySession(method, path, session string) bool
	}

	// UserLoader - allows us to load users
	UserLoader interface {
		LoadUser(method, path, session string) *User
	}

	// User - a minimal user abstraction
	User struct {
		ID    string
		Roles []string
	}

	// CookieSessionExtractor - allows us to extract sessions from cookies
	CookieSessionExtractor struct {
		Field string
	}

	// HeaderSessionExtractor - allows us to extract sessions from headers
	HeaderSessionExtractor struct {
		Header string
		Prefix string
	}

	// Result - universal return type
	Result struct {
		Code int         `json:"code"`
		Body interface{} `json:"body"`
		Hook func(*gin.Context)
	}

	// ErrorDTO - how errors are returned
	ErrorDTO struct {
		Message string `json:"message"`
	}
)

// Read - reads the session value from the cookie
func (c *CookieSessionExtractor) Read(ctx *gin.Context) string {
	cookie, _ := ctx.Cookie(c.Field)
	return cookie
}

// Read - reads the session value from the header
func (h *HeaderSessionExtractor) Read(ctx *gin.Context) string {
	session := ctx.GetHeader(h.Header)
	if h.Prefix != "" {
		session = strings.ReplaceAll(session, h.Prefix, "")
	}
	return session
}
