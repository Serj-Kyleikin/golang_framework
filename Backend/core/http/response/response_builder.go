package response

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Envelope struct {
	Status  bool        `json:"status"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

type builder struct {
	statusCode int
	env        Envelope
}

var pool = sync.Pool{
	New: func() any {
		return &builder{}
	},
}

func get() *builder {
	b := pool.Get().(*builder)

	b.statusCode = 0
	b.env.Status = false
	b.env.Message = nil
	b.env.Data = nil
	return b
}

func put(b *builder) {

	b.env.Message = nil
	b.env.Data = nil
	pool.Put(b)
}

func Status(code int) *builder {
	b := get()
	return b.Status(code)
}

func Message(msg string) *builder {
	b := get()
	return b.Message(msg)
}

func Data(data interface{}) *builder {
	b := get()
	return b.Data(data)
}

func (b *builder) Status(code int) *builder {
	if isValidHTTPStatus(code) {
		b.statusCode = code
	}
	return b
}

func (b *builder) Message(msg string) *builder {
	b.env.Message = &msg
	return b
}

func (b *builder) Data(data interface{}) *builder {
	b.env.Data = data
	return b
}

func (b *builder) Success(c *gin.Context) {
	code := b.statusCode
	if !isValidHTTPStatus(code) {
		code = http.StatusOK
	}

	b.env.Status = true
	c.JSON(code, b.env)
	put(b)
}

func (b *builder) Fail(c *gin.Context) {
	code := b.statusCode
	if !isValidHTTPStatus(code) {
		code = http.StatusBadRequest
	}

	b.env.Status = false
	c.JSON(code, b.env)
	put(b)
}

func isValidHTTPStatus(code int) bool {
	return code >= 100 && code <= 599
}
