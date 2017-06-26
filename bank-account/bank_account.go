package account

const testVersion = 1

type Account struct {
	balance int64
	status  bool
}

func Open(initialDeposit int64) *Account {
	account := new(Account)
	if initialDeposit >= 0 {
		account.status = true
		account.Deposit(initialDeposit)
		return account
	}
	return nil
}

func (this *Account) Close() (int64, bool) {
	if this.status == true {
		this.status = false
		payout := this.balance
		this.balance = 0
		return payout, true
	}
	return 0, false
}

func (this *Account) Balance() (int64, bool) {
	if this.status == true {
		return this.balance, true
	}
	return 0, false
}

func (this *Account) Deposit(amount int64) (int64, bool) {
	if this.status == true {
		if this.balance+amount >= 0 {
			this.balance += amount
			return this.balance, true
		} else {
			return this.balance, false
		}
	}
	return 0, false
}