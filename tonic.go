package tonic

import (
	"log"
	"strings"

	"github.com/Meduzz/wendy"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type (
	// CookieExtractor - allows us to extract a field from cookies
	CookieExtractor struct {
		Field string
	}

	// HeaderExtractor - allows us to extract a field from headers
	HeaderExtractor struct {
		Header string
		Prefix string
	}

	BodyExtractor struct {
		Field string
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

// ReadGin - reads the field value from the http cookie
func (c *CookieExtractor) ReadGin(ctx *gin.Context) string {
	cookie, _ := ctx.Cookie(c.Field)
	return cookie
}

// ReadGin - reads the field value from the http (gin) header
func (h *HeaderExtractor) ReadGin(ctx *gin.Context) string {
	session := ctx.GetHeader(h.Header)
	if h.Prefix != "" {
		session = strings.ReplaceAll(session, h.Prefix, "")
	}
	return session
}

// ReadWendy - reads the field value from the wendy header
func (h *HeaderExtractor) ReadWendy(request *wendy.Request) string {
	value := request.Headers[h.Header]

	if h.Prefix != "" {
		value = strings.ReplaceAll(value, h.Prefix, "")
	}

	return value
}

// ReadGin - reads a field value from the http (gin) body and add it to the context under "body".
func (b *BodyExtractor) ReadGin(ctx *gin.Context) string {
	raw, err := ctx.GetRawData()

	if err != nil {
		log.Printf("GetRawData threw error: %v\n", err)
		return ""
	}

	ctx.Set("body", raw)

	res := gjson.GetBytes(raw, b.Field)

	return res.String()
}

// ReadWendy - reads the field value from the wendy body
func (b *BodyExtractor) ReadWendy(reqeuest *wendy.Request) string {
	res := gjson.GetBytes(reqeuest.Body.Data, b.Field)
	return res.String()
}
