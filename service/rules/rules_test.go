package rules_test

import (
	"testing"

	"github.com/emehrkay/rpc/service/rules"
	"github.com/emehrkay/rpc/tests"
)

func TestProcessRecepitsWithDefaultRules(t *testing.T) {
	sk := rules.New()

	for i, c := range tests.Cases {
		res, err := sk.ProcessReceipt(*c.Receipt)
		if err != nil {
			t.Errorf(`unable to process receipt at index: %d to recepit -- %v`, i, err)
		}

		if res != c.Points {
			t.Errorf(`case at index: %d expected %d, but got %d`, i, c.Points, res)
		}
	}
}
