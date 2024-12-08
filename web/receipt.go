package web

import (
	"encoding/json"
	"net/http"

	"github.com/emehrkay/rpc/storage"
	"github.com/google/uuid"
)

type ReceiptCreatedResponse struct {
	ID uuid.UUID `json:"id"`
}

type ReceiptPointsResponse struct {
	Points uint64 `json:"points"`
}

type ReceiptsResponse struct {
	Receipts []storage.ReceiptRecord `json:"recepits"`
}

func (s *Server) RecepitSave(w http.ResponseWriter, r *http.Request) {
	var rec storage.Receipt
	if r.Body == nil || r.Body == http.NoBody {
		s.HandleError(w, RequestBodyRequiredErr)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rec)
	if err != nil {
		e := RequestBodyInvalidErr
		e.OrignialError = err
		s.HandleError(w, e)
		return
	}

	resp, err := s.rpc.Receipt.Save(rec)
	if err != nil {
		s.HandleError(w, err)
		return
	}

	// the api.yaml file expects this to be a 200 and not a 201
	s.RespondJson(w, http.StatusOK, ReceiptCreatedResponse{
		ID: resp.ID,
	})
}

func (s *Server) RecepitGetPointsByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		e := InvalidUUIDErr
		e.OrignialError = err
		s.HandleError(w, e)
		return
	}

	record, err := s.rpc.Receipt.GetByID(id)
	if err != nil {
		s.HandleError(w, err)
		return
	}

	s.RespondJson(w, http.StatusOK, ReceiptPointsResponse{
		Points: record.Points,
	})
}

func (s *Server) RecepitGetAll(w http.ResponseWriter, r *http.Request) {
	receipts, err := s.rpc.Receipt.GetAll()
	if err != nil {
		s.HandleError(w, err)
		return
	}

	s.RespondJson(w, http.StatusOK, ReceiptsResponse{
		Receipts: receipts,
	})
}
