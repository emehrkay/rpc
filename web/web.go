package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/emehrkay/rpc/service"
	"github.com/emehrkay/rpc/storage"
)

func New(port string, rpc *service.Service, ruoter *http.ServeMux) (*Server, error) {
	server := &Server{
		port:   port,
		rpc:    rpc,
		router: ruoter,
		log:    rpc.Log,
	}

	return server, nil
}

type Server struct {
	rpc    *service.Service
	router *http.ServeMux
	log    *slog.Logger
	port   string
}

func (s *Server) Run() error {
	fmt.Printf("SERVER STARTED: %s\n", s.port)
	s.Routes()

	return http.ListenAndServe(s.port, s.router)
}

func (s *Server) Routes() {
	s.router.HandleFunc("POST /receipts/process", s.RecepitSave)
	s.router.HandleFunc("GET /receipts/{id}/points", s.RecepitGetPointsByID)
	s.router.HandleFunc("GET /receipts", s.RecepitGetAll)
}

// repackage all errors as something safe to return to users
func (s *Server) HandleError(w http.ResponseWriter, err error) {
	notFound := storage.ErrNotFound
	if errors.Is(err, notFound) {
		e := HttpError{
			OrignialError: err,
			StatusCode:    http.StatusNotFound,
			Message:       http.StatusText(http.StatusNotFound),
		}
		s.RespondJson(w, e.StatusCode, e.WebError())
		return
	}

	httpe := HttpError{}
	if errors.As(err, &httpe) {
		e := err.(HttpError)
		s.RespondJson(w, e.StatusCode, e.WebError())
		return
	}

	e := HttpError{
		OrignialError: err,
		StatusCode:    http.StatusInternalServerError,
	}
	s.log.Error("request error", "err", e.Error())
	s.RespondJson(w, e.StatusCode, e.WebError())
}

func (s *Server) RespondJson(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
