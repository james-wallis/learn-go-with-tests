package pointersanderrors

import (
	"errors"
	"fmt"
)

// In Go when you call a function or a method the arguments are copied
// To change the actual value we use &myVal to get the memory address for the value
// Pointers let us point to some values and let us change them

type (
	// Bitcoin : a type which represents a Bitcoin
	Bitcoin int

	// Wallet : a struct with a balance of Bitcoins
	Wallet struct {
		balance Bitcoin
	}
)

// ErrInsufficientFunds : amount greater than balance error
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Balance : returns the balance for a Wallet
func (w *Wallet) Balance() Bitcoin {
	// As this is a pointer w.balance is actually (*w).balance
	// It is a struct pointer and automatically dereferenced
	return w.balance
}

// Deposit : increases the balance for a Wallet by the amount given
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw : decreases the balance for a Wallet by the amount given
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// String : returns a string value for Bitcoin (which is an int)
func (b Bitcoin) String() string {
	// This means that when we use fmt.Printf with %s on a Bitcoin it will call this method
	// Instead of printing "20" it will print out "20 BTC" so it is clearer whats going on
	return fmt.Sprintf("%d BTC", b)
}
