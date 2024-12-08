package service

import (
	"fmt"

	"github.com/emehrkay/rpc/service/rules"
	"github.com/emehrkay/rpc/storage"
	"github.com/google/uuid"
)

type receiptService struct {
	service *Service
}

func (r *receiptService) Save(receipt storage.Receipt) (*storage.ReceiptRecord, error) {
	score := rules.New()
	points, err := score.ProcessReceipt(receipt)
	if err != nil {
		return nil, fmt.Errorf(`unable to get score for receipt -- %w`, err)
	}

	record, err := r.service.store.SaveReceipt(receipt, points)
	if err != nil {
		return nil, fmt.Errorf(`unable to save receipt to storage -- %w`, err)
	}

	return record, nil
}

func (r *receiptService) GetByID(id uuid.UUID) (*storage.ReceiptRecord, error) {
	record, err := r.service.store.GetReceipt(id)
	if err != nil {
		return nil, fmt.Errorf(`unable to retrieve receipt with id: %s -- %w`, id, err)
	}

	return record, nil
}

func (r *receiptService) GetAll() ([]storage.ReceiptRecord, error) {
	return r.service.store.GetAllReceipts()
}
