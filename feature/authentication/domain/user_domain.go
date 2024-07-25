package domain

import (
	"time"
)

type User struct {
	ID                 int       `gorm:"primaryKey;autoIncrement"`
	RoleID             int       `gorm:"index;default:null"`
	DepartmentID       int       `gorm:"index;default:null"`
	Email              string    `gorm:"type:varchar(45);not null"`
	Password           string    `gorm:"type:text;not null"`
	Name               string    `gorm:"type:varchar(45);not null"`
	Surname            string    `gorm:"type:varchar(45);not null"`
	Gender             string    `gorm:"type:varchar(20);default:null"`
	DOB                time.Time `gorm:"default:null"`
	Mobile             string    `gorm:"type:varchar(15);default:null"`
	CountryID          int       `gorm:"index;default:null"`
	ResidentCountryID  int       `gorm:"index;default:null"`
	Avatar             string    `gorm:"type:varchar(100);default:null"`
	VerificationStatus int       `gorm:"default:0;comment:0: unverified\n1: verified"`
	Status             int       `gorm:"not null;comment:0: inactive\n1: active"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
