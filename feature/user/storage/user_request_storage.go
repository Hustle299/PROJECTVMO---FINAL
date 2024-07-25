package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"

	"gorm.io/gorm"
)

type ApplicantRequestRepositoryInterface interface {
	CreateApplicantRequest(request *domain.Request) error
}

type ApplicantRequestRepository struct {
	DB *gorm.DB
}

func NewApplicantRequestRepository(db *gorm.DB) *ApplicantRequestRepository {
	return &ApplicantRequestRepository{DB: db}
}

func (r *ApplicantRequestRepository) CreateApplicantRequest(request *domain.Request) error {
	return r.DB.Create(request).Error
}
