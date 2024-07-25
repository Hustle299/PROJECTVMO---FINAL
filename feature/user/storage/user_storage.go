package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"

	"gorm.io/gorm"
)

type ApplicantRepositoryInterface interface {
	UpdateApplicantBasic(user *domain.User) error
	UpdateApplicant(user *domain.User) error
	DeleteApplicant(id int) error
	FindApplicantByID(id int) (*domain.User, error)
}

type ApplicantRepository struct {
	DB *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) *ApplicantRepository {
	return &ApplicantRepository{DB: db}
}

func (r *ApplicantRepository) UpdateApplicantBasic(user *domain.User) error {
	return r.DB.Save(user).Error
}

func (r *ApplicantRepository) UpdateApplicant(user *domain.User) error {
	return r.DB.Save(user).Error
}

func (r *ApplicantRepository) DeleteApplicant(id int) error {
	return r.DB.Delete(&domain.User{}, id).Error
}

func (r *ApplicantRepository) FindApplicantByID(id int) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
