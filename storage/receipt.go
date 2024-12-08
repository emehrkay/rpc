package storage

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DateOnly struct {
	time.Time
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("2006-1-2", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}

	d.Time = t

	return nil
}

func (d DateOnly) String() string {
	return fmt.Sprintf("%q", d.Time.Format("2006-1-2"))
}

type TimeOnly struct {
	time.Time
}

func (to TimeOnly) MarshalJSON() ([]byte, error) {
	return []byte(to.String()), nil
}

func (to *TimeOnly) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("15:04", string(b[1:len(b)-1]))
	if err != nil {
		return err
	}

	to.Time = t

	return nil
}

func (to TimeOnly) String() string {
	return fmt.Sprintf("%q", to.Time.Format("15:04"))
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

type Receipt struct {
	Retailer     string   `json:"retailer"`
	PurchaseDate DateOnly `json:"purchaseDate"`
	PurchaseTime TimeOnly `json:"purchaseTime"`
	Items        []Item   `json:"items"`
	Total        float64  `json:"total,string"`
}

type ReceiptRecord struct {
	ID      uuid.UUID `json:"id"`
	Receipt Receipt   `json:"receipt"`
	Points  uint64    `json:"points"`
}
