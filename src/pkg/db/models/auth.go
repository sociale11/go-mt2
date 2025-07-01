package models

import (
	"time"
)

type Account struct {
	Id              int32 `gorm:"primaryKey"`
	Username        string
	Password        string
	Email           string
	AccountStatusId int
	DeleteCode      string
	LastLogin       time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type AccountStatus struct {
	Id           int32 `gorm:"primaryKey"`
	ClientStatus string
	AllowLogin   bool
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AccountStatuses []AccountStatus
type Accounts []Account
