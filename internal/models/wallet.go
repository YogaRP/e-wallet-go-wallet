package models

import "time"

type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id" gorm:"column:user_id;unique"`
	Balance   float64   `json:"balance" gorm:"column:balance;type:decimal(15, 2)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID                    int
	WalletID              int     `gorm:"column:wallet_id"`
	Amount                float64 `gorm:"column:amount;type:decimal(15, 2)"`
	WalletTransactionType string  `gorm:"column:wallet_transaction_type;type:ENUM('CREDIT', 'DEBIT')"`
	Reference             string  `gorm:"column:type;type:varchar(100)"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}
