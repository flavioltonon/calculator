package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator/application/http/dto"
	"calculator/domain/services"
	"calculator/domain/valueobject"

	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite

	controller *Controller
}

func (s *ControllerTestSuite) SetupSuite() {
	s.controller = &Controller{
		calculator: services.NewCalculator(
			services.NewAdder(),
			services.NewSubtractor(),
			services.NewMultiplier(),
			services.NewDivider(),
		),
	}
}

func TestController(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}

func (s *ControllerTestSuite) TestController_Add() {
	type args struct {
		r func() *http.Request
	}

	type want struct {
		statusCode int
		err        bool
		result     valueobject.Number
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "When the request body is invalid, a 400 HTTP status code should be returned",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodPost, "/add", bytes.NewBuffer([]byte(`foo`)))
				},
			},
			want: want{
				statusCode: http.StatusBadRequest,
				err:        true,
			},
		},
		{
			name: "When the request body is valid, a 200 HTTP status code and a Calculation should be returned",
			args: args{
				r: func() *http.Request {
					b, _ := json.Marshal(dto.Operands{
						X: 1,
						Y: 2,
					})
					return httptest.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(b))
				},
			},
			want: want{
				statusCode: http.StatusOK,
				result:     3,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			recorder := httptest.NewRecorder()
			s.controller.Add(recorder, tt.args.r())
			s.Equal("application/json", recorder.Header().Get("Content-Type"))
			s.Equal(tt.want.statusCode, recorder.Code)
		})
	}
}

func (s *ControllerTestSuite) TestController_Subtract() {
	type args struct {
		r func() *http.Request
	}

	type want struct {
		statusCode int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "When the request body is invalid, a 400 HTTP status code should be returned",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodPost, "/subtract", bytes.NewBuffer([]byte(`foo`)))
				},
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "When the request body is valid, a 200 HTTP status code and a Calculation should be returned",
			args: args{
				r: func() *http.Request {
					b, _ := json.Marshal(dto.Operands{
						X: 1,
						Y: 2,
					})
					return httptest.NewRequest(http.MethodPost, "/subtract", bytes.NewBuffer(b))
				},
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			recorder := httptest.NewRecorder()
			s.controller.Subtract(recorder, tt.args.r())
			s.Equal(recorder.Header().Get("Content-Type"), "application/json")
			s.Equal(recorder.Code, tt.want.statusCode)
		})
	}
}

func (s *ControllerTestSuite) TestController_Multiply() {
	type args struct {
		r func() *http.Request
	}

	type want struct {
		statusCode int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "When the request body is invalid, a 400 HTTP status code should be returned",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodPost, "/multiply", bytes.NewBuffer([]byte(`foo`)))
				},
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "When the request body is valid, a 200 HTTP status code and a Calculation should be returned",
			args: args{
				r: func() *http.Request {
					b, _ := json.Marshal(dto.Operands{
						X: 1,
						Y: 2,
					})
					return httptest.NewRequest(http.MethodPost, "/multiply", bytes.NewBuffer(b))
				},
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			recorder := httptest.NewRecorder()
			s.controller.Multiply(recorder, tt.args.r())
			s.Equal(recorder.Header().Get("Content-Type"), "application/json")
			s.Equal(recorder.Code, tt.want.statusCode)
		})
	}
}

func (s *ControllerTestSuite) TestController_Divide() {
	type args struct {
		r func() *http.Request
	}

	type want struct {
		statusCode int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "When the request body is invalid, a 400 HTTP status code should be returned",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodPost, "/divide", bytes.NewBuffer([]byte(`foo`)))
				},
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "When the request body is valid, a 200 HTTP status code and a Calculation should be returned",
			args: args{
				r: func() *http.Request {
					b, _ := json.Marshal(dto.Operands{
						X: 1,
						Y: 2,
					})
					return httptest.NewRequest(http.MethodPost, "/divide", bytes.NewBuffer(b))
				},
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			recorder := httptest.NewRecorder()
			s.controller.Divide(recorder, tt.args.r())
			s.Equal(recorder.Header().Get("Content-Type"), "application/json")
			s.Equal(recorder.Code, tt.want.statusCode)
		})
	}
}
