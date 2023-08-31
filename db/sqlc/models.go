// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Cart struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type CartDetail struct {
	ID            int64 `json:"id"`
	CartID        int64 `json:"cart_id"`
	ProductID     int64 `json:"product_id"`
	QuantityAdded int64 `json:"quantity_added"`
}

type Order struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	TotalPrice int64     `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderDetail struct {
	ID                 int64  `json:"id"`
	ProductID          int64  `json:"product_id"`
	Username           string `json:"username"`
	QuantityOrdered    int64  `json:"quantity_ordered"`
	PriceAtTimeOfOrder int64  `json:"price_at_time_of_order"`
}

type Product struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         int64     `json:"price"`
	StockQuantity int64     `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	Username          string      `json:"username"`
	Email             string      `json:"email"`
	HashedPassword    string      `json:"hashed_password"`
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Address           pgtype.Text `json:"address"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `json:"created_at"`
	IsAdmin           bool        `json:"is_admin"`
}
