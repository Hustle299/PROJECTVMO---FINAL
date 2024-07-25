package domain

import "time"

type Request struct {
	ID          int    `gorm:"primaryKey"`
	UserID      uint   `gorm:"index"`
	Type        string `gorm:"not null"`
	Status      int    `gorm:"not null"`
	RejectNotes string
	VerifierID  int       `gorm:"index"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
