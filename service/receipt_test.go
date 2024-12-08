package service_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/google/uuid"

	"github.com/emehrkay/rpc/service"
	"github.com/emehrkay/rpc/storage"
	"github.com/emehrkay/rpc/tests"
)

var (
	ser *service.Service
)

func init() {
	var err error
	store := storage.NewMemory()
	log := slog.Default()
	ser, err = service.New(store, log)
	if err != nil {
		panic(fmt.Sprintf(`unable to create service -- %v`, err))
	}
}

func TestSave(t *testing.T) {
	for i, c := range tests.Cases {
		resp, err := ser.Receipt.Save(*c.Receipt)
		if err != nil {
			t.Errorf(`unable to save receipt at index: %d to recepit -- %v`, i, err)
		}

		if resp == nil || resp.ID == uuid.Nil {
			t.Errorf(`invalid resp from Save at index: %d`, i)
		}

		if resp.Points != c.Points {
			t.Errorf(`case at index: %d expected %d, but got %d`, i, c.Points, resp.Points)
		}
	}
}

func TestGetByID(t *testing.T) {
	for i, c := range tests.Cases {
		resp, err := ser.Receipt.Save(*c.Receipt)
		if err != nil {
			t.Errorf(`unable to save receipt at index: %d to recepit -- %v`, i, err)
		}

		if resp == nil || resp.ID == uuid.Nil {
			t.Errorf(`invalid resp from Save at index: %d`, i)
		}

		if resp.Points != c.Points {
			t.Errorf(`case at index: %d expected %d, but got %d`, i, c.Points, resp.Points)
		}

		byID, err := ser.Receipt.GetByID(resp.ID)
		if err != nil {
			t.Errorf(`unable to call GetByID(%v) to recepit -- %v`, resp.ID, err)
		}

		if byID.Points != resp.Points || resp.ID != byID.ID {
			t.Errorf(`resp with id: %v is not valid`, byID.ID)
		}
	}

	rec, err := ser.Receipt.GetByID(uuid.New())
	if rec != nil && err == nil {
		t.Errorf(`expected an invalid result from non-existant id`)
	}
}
