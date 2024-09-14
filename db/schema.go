package database

import (
	"time"
)

type Customer struct {
	ID        uint   `json:id gorm:"primaryKey;default:auto_random()"`
	Username  string `json:username`
	Password  string `json:password`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tokens    []Token
}

type Token struct {
	ID         uint   `json:id gorm:"primaryKey;default:auto_random()"`
	Token      string `json:token`
	CustomerID uint   `json:customer_id`
	CreatedAt  time.Time
	ExpiresAt  time.Time
}
