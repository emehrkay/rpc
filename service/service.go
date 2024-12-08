package service

import (
	"log/slog"

	"github.com/emehrkay/rpc/storage"
)

func New(store storage.Storage, log *slog.Logger) (*Service, error) {
	serv := &Service{
		store: store,
		Log:   log,
	}

	serv.Receipt = &receiptService{
		service: serv,
	}

	return serv, nil
}

type Service struct {
	store   storage.Storage
	Log     *slog.Logger
	Receipt *receiptService
}
