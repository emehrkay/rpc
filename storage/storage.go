package storage

import "github.com/google/uuid"

type Storage interface {
	SaveReceipt(receipt Receipt, points uint64) (*ReceiptRecord, error)
	GetReceipt(id uuid.UUID) (*ReceiptRecord, error)
	GetAllReceipts() ([]ReceiptRecord, error)
}
