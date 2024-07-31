package golearn

type Account struct {
	Id      int `json:"id"`
	Balance int `json:"balance"`
}

// Method to check the balance after a hypothetical deposit
func (account Account) CheckBalance() int {
	return account.Balance
}

func (account *Account) Deposit(amount int) {
	account.Balance += amount
}

// Method to withdraw an amount from the account
func (account *Account) Withdraw(amount int) {
	account.Balance -= amount
}
