package wallet

import "fmt"

type Bitcoin int

// %sでprintした時に実行される
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// レシーバは関数内でコピーされて渡される。メンバ変数のアドレスは実行ごとに変わる
// Walletへのポインタとして指定することで、取得先がポインタになる

func (w *Wallet) Deposit(amount Bitcoin) {
	// 内部的には (*w).balance となっている
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
