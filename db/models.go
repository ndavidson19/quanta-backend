// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type Account struct {
	ID          int64        `json:"id"`
	Owner       string       `json:"owner"`
	Name        string       `json:"name"`
	Balance     int64        `json:"balance"`
	Currency    string       `json:"currency"`
	CreatedAt   time.Time    `json:"created_at"`
	LastUpdated time.Time    `json:"last_updated"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

type AuditLog struct {
	ID                int64         `json:"id"`
	AccountID         int64         `json:"account_id"`
	Action            string        `json:"action"`
	ActionType        string        `json:"action_type"`
	ActionDescription string        `json:"action_description"`
	PerformedBy       string        `json:"performed_by"`
	AffectedRecord    sql.NullInt64 `json:"affected_record"`
	Timestamp         sql.NullTime  `json:"timestamp"`
}

type Deposit struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FinancialInstrument struct {
	ID             int64                 `json:"id"`
	Symbol         string                `json:"symbol"`
	Name           string                `json:"name"`
	InstrumentType string                `json:"instrument_type"`
	Details        pqtype.NullRawMessage `json:"details"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
}

type Notification struct {
	ID               int64     `json:"id"`
	UserID           string    `json:"user_id"`
	Message          string    `json:"message"`
	NotificationDate time.Time `json:"notification_date"`
}

type PortfolioBalance struct {
	ID          int64          `json:"id"`
	AccountID   int64          `json:"account_id"`
	Symbol      sql.NullString `json:"symbol"`
	Amount      sql.NullInt32  `json:"amount"`
	LastUpdated sql.NullTime   `json:"last_updated"`
}

type Trade struct {
	ID        int64  `json:"id"`
	AccountID int64  `json:"account_id"`
	Symbol    string `json:"symbol"`
	// must be positive
	Amount    string       `json:"amount"`
	Price     string       `json:"price"`
	TradeType string       `json:"trade_type"`
	Status    string       `json:"status"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Transaction struct {
	ID                int64         `json:"id"`
	AccountID         int64         `json:"account_id"`
	TransactionType   string        `json:"transaction_type"`
	TransactionAmount string        `json:"transaction_amount"`
	TransactionDate   time.Time     `json:"transaction_date"`
	RelatedTradeID    sql.NullInt64 `json:"related_trade_id"`
}

type User struct {
	Username            string         `json:"username"`
	HashedPassword      string         `json:"hashed_password"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	PhoneNumber         sql.NullString `json:"phone_number"`
	PasswordChangedAt   time.Time      `json:"password_changed_at"`
	CreatedAt           time.Time      `json:"created_at"`
	LastLoginAt         sql.NullTime   `json:"last_login_at"`
	LoginAttempts       sql.NullInt32  `json:"login_attempts"`
	LockedUntil         sql.NullTime   `json:"locked_until"`
	ResetToken          sql.NullString `json:"reset_token"`
	ResetTokenExpiresAt sql.NullTime   `json:"reset_token_expires_at"`
	Role                string         `json:"role"`
}
