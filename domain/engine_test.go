package domain

import (
	"encoding/json"
	"testing"
)

func TestPaymentBookLifecycle(t *testing.T) {
	engine := NewPaymentBook()
	if err := engine.OpenAccount(); err != nil {
		t.Fatal(err)
	}
	if err := engine.RegisterPayment(PaymentBookRecord{ID: "primary", Quantity: 4, Labels: map[string]string{"zone": "a"}}); err != nil {
		t.Fatal(err)
	}
	if err := engine.PostPayment("primary", 3); err != nil {
		t.Fatal(err)
	}
	if err := engine.ReversePayment("primary", 2); err != nil {
		t.Fatal(err)
	}
	if got := engine.CountBalance(); got != 5 {
		t.Fatalf("count = %d; want 5", got)
	}
	if err := engine.FreezeAccount(); err != nil {
		t.Fatal(err)
	}
}

func TestPaymentBookPrioritiesAndExport(t *testing.T) {
	engine := NewPaymentBook()
	_ = engine.RegisterPayment(PaymentBookRecord{ID: "low", Quantity: 1})
	_ = engine.RegisterPayment(PaymentBookRecord{ID: "high", Quantity: 2})
	if err := engine.PrioritizePayment("high", 9); err != nil {
		t.Fatal(err)
	}
	values := engine.List()
	if len(values) != 2 || values[0].ID != "high" {
		t.Fatalf("unexpected order: %#v", values)
	}
	values[0].Labels = map[string]string{"changed": "yes"}
	data, err := engine.ExportPayments()
	if err != nil || !json.Valid(data) {
		t.Fatalf("invalid export: %s, %v", data, err)
	}
}

func TestPaymentBookRejectsInvalidOperations(t *testing.T) {
	engine := NewPaymentBook()
	if err := engine.RegisterPayment(PaymentBookRecord{}); err == nil {
		t.Fatal("expected blank id error")
	}
	if err := engine.RegisterPayment(PaymentBookRecord{ID: "one", Quantity: 1}); err != nil {
		t.Fatal(err)
	}
	if err := engine.RegisterPayment(PaymentBookRecord{ID: "one"}); err == nil {
		t.Fatal("expected duplicate error")
	}
	if err := engine.ReversePayment("one", 2); err == nil {
		t.Fatal("expected insufficient quantity error")
	}
	if err := engine.PrioritizePayment("missing", 1); err == nil {
		t.Fatal("expected missing record error")
	}
}
