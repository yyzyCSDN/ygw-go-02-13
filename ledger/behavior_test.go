package ledger

import (
	"testing"
)

func TestLedgerPost(t *testing.T) {
	l := NewLedger()
	if err := l.Post("one", 10); err != nil || l.Balance() != 10 {
		t.Fatal("post failed")
	}
}
