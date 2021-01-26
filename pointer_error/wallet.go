package wallet

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// レシーバは関数内でコピーされて渡される。メンバ変数のアドレスは実行ごとに変わる
// Walletへのポインタとして指定することで、取得先がポインタになる

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// 内部的には (*w).balance となっている
	return w.balance
}
