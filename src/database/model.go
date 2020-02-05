package database

import "database/sql"

type Login struct {
	ID       string `json:"id"`
	Bank     string `json:"bank"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *Login) GetLoginByBankAndUsername(db *sql.DB) error {
	return db.QueryRow(
		"SELECT bank, username, password FROM login WHERE bank=$1 AND username=$2",
		login.Bank,
		login.Username,
	).Scan(&login.Bank, &login.Username, &login.Password)
}

func (login *Login) CreateLogin(db *sql.DB) error {
	return db.QueryRow(
		"INSERT INTO logins(id, bank, username, password) VALUES($1, $2, $3) RETURNING id",
		login.ID,
		login.Username,
		login.Password,
		login.Bank,
	).Scan(&login.ID)
}

type Account struct {
	ID      string `json:"id"`
	LoginID string

	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (account *Account) CreateAccount(db *sql.DB, loginID string) error {
	return db.QueryRow(
		"INSERT INTO accounts(id, login_id, name, balance) VALUES($1, $2, $3, $4) RETURNING id",
		account.ID,
		account.LoginID,
		account.Name,
		account.Balance,
	).Scan(&account.ID)
}
