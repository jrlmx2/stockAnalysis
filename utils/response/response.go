package response

import (
	"fmt"
	"net/http"
)

type Respondable interface {
	toJSON() []byte
}

type response struct {
	writer http.ResponseWriter
	code   int
}

func (r *response) Respond(i Respondable) {
	code, err := r.writer.Write(i.toJSON())
	if err != nil {
		fmt.Printf("Responding failed with %d, %s", code, err)
	}
}

func Prepare(code int, writer http.ResponseWriter) *response {
	writer.WriteHeader(code)

	return &response{
		writer: writer,
	}
}
