package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/storage"
)

// RoleUsecase handles the business logic for roles.
type RoleUsecase struct {
	repo *storage.RoleRepository
}

// NewRoleUsecase creates a new instance of RoleUsecase.
func NewRoleUsecase(repo *storage.RoleRepository) *RoleUsecase {
	return &RoleUsecase{repo: repo}
}

// CreateRole creates a new role using the provided DTO.
func (u *RoleUsecase) CreateRole(input dto.RoleCreateDTO) (*domain.Role, error) {
	role := &domain.Role{
		Name:   input.Name,
		Status: input.Status,
	}
	err := u.repo.Create(role)
	return role, err
}

// GetRoleByID retrieves a role by its ID.
func (u *RoleUsecase) GetRoleByID(id uint) (*domain.Role, error) {
	return u.repo.GetByID(id)
}

// UpdateRole updates a role using the provided DTO.
func (u *RoleUsecase) UpdateRole(id uint, input dto.RoleUpdateDTO) (*domain.Role, error) {
	role, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	role.Name = input.Name
	role.Status = input.Status
	err = u.repo.Update(role)
	return role, err
}

// DeleteRole deletes a role by its ID.
func (u *RoleUsecase) DeleteRole(id uint) error {
	return u.repo.Delete(id)
}
