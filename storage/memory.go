package storage

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("memory storage: no records found")
)

func NewMemory() *memory {
	return &memory{
		records: make(map[uuid.UUID]*ReceiptRecord),
		mu:      &sync.Mutex{},
	}
}

type memory struct {
	records map[uuid.UUID]*ReceiptRecord
	mu      *sync.Mutex
}

func (m *memory) SaveReceipt(receipt Receipt, points uint64) (*ReceiptRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	rec := &ReceiptRecord{
		ID:      uuid.New(),
		Receipt: receipt,
		Points:  points,
	}
	m.records[rec.ID] = rec

	return rec, nil
}

func (m *memory) GetReceipt(id uuid.UUID) (*ReceiptRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if rec, ok := m.records[id]; ok {
		return rec, nil
	}

	return nil, ErrNotFound
}

func (m *memory) GetAllReceipts() ([]ReceiptRecord, error) {
	records := make([]ReceiptRecord, len(m.records))
	i := 0
	for _, rec := range m.records {
		records[i] = *rec
		i += 1
	}

	return records, nil
}
