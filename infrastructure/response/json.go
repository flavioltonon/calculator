package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, content Content) {
	if b, err := json.Marshal(Response{
		Kind:    content.Kind(),
		Content: content,
	}); err != nil {
		JSON(w, http.StatusInternalServerError, Error(err))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(b)
	}
}

type Content interface {
	Kind() string
}

type Response struct {
	Kind    string  `json:"kind"`
	Content Content `json:"content"`
}

type ErrorContent struct {
	Message string `json:"message"`
}

func (c ErrorContent) Kind() string {
	return "error"
}

func Error(err error) Content {
	return ErrorContent{
		Message: err.Error(),
	}
}

type SuccessContent struct {
	data any
}

func (c SuccessContent) Kind() string {
	return "success"
}

func Success(data any) Content {
	return SuccessContent{data: data}
}

func (c SuccessContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.data)
}
