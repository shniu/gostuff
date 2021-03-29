package usecase

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCase(t *testing.T) {
	bootstrap()
}

func TestCreateJournalRequestMarshal(t *testing.T) {
	journal := &CreateJournalRequest{
		Status:       "posted",
		EntryType:    "posted",
		SourceLedger: "posted",
		Reference:    "posted",
		Currency:     "posted",
		Notes:        "posted",
		Transactions: []Transaction{},
	}

	journal.Transactions = append(journal.Transactions, Transaction{Credit: "0.00"})

	bytes, err := json.Marshal(journal)
	fmt.Println("err is ", err, " bytes is", string(bytes))
}
