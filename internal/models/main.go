package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" gorm:"default:CURRENT_TIMESTAMP"`
}

type UserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type Guest struct {
	ID           int       `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Confirmation bool      `json:"confirmation"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime" gorm:"default:CURRENT_TIMESTAMP"`
}

type GuestInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type Gift struct {
	ID          int       `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" gorm:"default:CURRENT_TIMESTAMP"`
}

type GiftResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Link        string    `json:"link"`
	BoughtBy    string    `json:"bought_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GiftInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Link        string  `json:"link"`
}

type BoughtGift struct {
	ID        int       `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	GuestID   int       `json:"guest_id"`
	GiftID    int       `json:"gift_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" gorm:"default:CURRENT_TIMESTAMP"`
}

type Payment struct {
	ID        int       `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	GuestID   int       `json:"guest_id"`
	GiftID    int       `json:"gift_id"`
	PaymentID string    `json:"payment_id"`
	QRCode    string    `json:"qr_code"`
	Link      string    `json:"link"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" gorm:"default:CURRENT_TIMESTAMP"`
}

type PaymentInput struct {
	GuestID int `json:"guest_id" binding:"required"`
	GiftID  int `json:"gift_id" binding:"required"`
}
