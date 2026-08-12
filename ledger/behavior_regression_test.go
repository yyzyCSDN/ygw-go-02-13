package ledger

import (
	"testing"
)

func TestRegressionBehavior(t *testing.T) {
	l := NewLedger()
	_ = l.Post("one", 10)
	if err := l.Post("one", 10); err == nil {
		t.Fatal("expected duplicate error")
	}
	if l.Balance() != 10 {
		t.Fatalf("balance=%d; want 10", l.Balance())
	}
}
