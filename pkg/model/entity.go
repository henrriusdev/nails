package model

import "time"

type BaseEntity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	BaseEntity
	Name     string
	Email    string
	Password string
	RoleID   string
	Role     Role
}

type Role struct {
	BaseEntity
	Name string
}

type Product struct {
	BaseEntity
	Name  string
	Price float64
}

type Customer struct {
	BaseEntity
	Name     string
	Email    string
	Phone    string
	Document string
}

type Appointment struct {
	BaseEntity
	CustomerID  string
	Customer    Customer
	Date        time.Time
	Description string
	UserID      string
	User        User
}
