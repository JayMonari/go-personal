package account

import "sync"

type Account struct {
	balance int64
	closed  bool
	sync.Mutex
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.closed {
		payout = a.balance
		a.closed, a.balance = true, 0
		ok = true
	}
	return
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	a.Lock()
	defer a.Unlock()

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.closed {
		return 0, false
	}
	a.Lock()
	defer a.Unlock()

	if a.balance + amount < 0 {
		return -1, false
	}
	a.balance += amount
	return a.balance, true
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}
