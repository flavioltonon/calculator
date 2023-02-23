package http

import (
	"fmt"
	"net/http"

	"calculator/application/http/controller"
	"calculator/application/http/middleware"
	"calculator/application/http/router"
	"calculator/domain/services"
	"calculator/infrastructure/logging/zap"
	"calculator/shared/logging"
)

type Server struct {
	core   *http.Server
	logger logging.Logger
}

func NewServer(host string, port int) (*Server, error) {
	var (
		adder      = services.NewAdder()
		subtractor = services.NewSubtractor()
		multiplier = services.NewMultiplier()
		divider    = services.NewDivider()
		service    = services.NewCalculator(adder, subtractor, multiplier, divider)
		controller = controller.New(service)
	)

	logger, err := zap.NewLogger()
	if err != nil {
		return nil, err
	}

	router := router.New(controller)
	router.AddMiddleware(middleware.Log(logger))

	return &Server{
		core: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", host, port),
			Handler: router,
		},
		logger: logger,
	}, nil
}

func (s *Server) ListenAndServe() error {
	s.logger.Info("Serving API", logging.String("address", s.core.Addr))
	return s.core.ListenAndServe()
}
