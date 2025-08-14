package controller

import (
	"log/slog"
	"net/http"
	"orderService/internal/usecase"
	"orderService/pkg/logger"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *mux.Router
	u      usecase.UsecaseInterface
	l      logger.LoggerInterface
}

func New(u usecase.UsecaseInterface, l logger.LoggerInterface) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		u:      u,
		l:      l,
	}
	s.Router.HandleFunc("/", s.HomeHandler).Methods("GET")
	s.Router.HandleFunc("/order/{order_uid}", s.GetOrderByOrderUID).Methods("GET")
	s.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fs := http.FileServer(http.Dir("static"))
	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return s
}

func (s *Server) Run(port string) {
	s.l.Info("server is started", slog.String("port", port))
	if err := http.ListenAndServe("0.0.0.0:"+port, s.Router); err != nil {
		s.l.Error("Run", slog.Any("error", "fatal error"))
		return

	}

}
