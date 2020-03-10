package pointer

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds : 当提取的比特币大于余额时提示的错误
var ErrInsufficientFunds = errors.New("can not withdraw, insufficient funds")

// Bitcoin : 比特币单位
type Bitcoin float64

// Stringer : 可序列化的类型
type Stringer interface {
	String() string
}

// Wallet : 钱包
type Wallet struct {
	balance Bitcoin
}

// Deposit : 设置金额
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance : 获取余额
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw : 提取金额
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
