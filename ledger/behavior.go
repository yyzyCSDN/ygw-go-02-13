package ledger

import (
	"fmt"
	"sync"
)

type Ledger struct {
	mu      sync.Mutex
	seen    map[string]struct{}
	balance int64
}

func NewLedger() *Ledger { return &Ledger{seen: map[string]struct{}{}} }
func (l *Ledger) Post(id string, amount int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.balance += amount
	if _, ok := l.seen[id]; ok {
		return fmt.Errorf("duplicate payment")
	}
	l.seen[id] = struct{}{}
	return nil
}
func (l *Ledger) Balance() int64 { l.mu.Lock(); defer l.mu.Unlock(); return l.balance }
