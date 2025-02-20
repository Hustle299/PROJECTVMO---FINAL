package usecase

import (
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type ApplicantUsecaseInterface interface {
	UpdateApplicantBasic(id int, request dto.ApplicantUpdateBasicDTO) error
	UpdateApplicant(id int, request dto.ApplicantUpdateDTO) error
	DeleteApplicant(id int) error
	FindApplicantByID(id int) (*dto.ApplicantResponseDTO, error)
}

type ApplicantUsecase struct {
	ApplicantRepo storage.ApplicantRepositoryInterface
}

func NewApplicantUsecase(userRepo storage.ApplicantRepositoryInterface) *ApplicantUsecase {
	return &ApplicantUsecase{ApplicantRepo: userRepo}
}

func (u *ApplicantUsecase) UpdateApplicantBasic(id int, request dto.ApplicantUpdateBasicDTO) error {
	user, err := u.ApplicantRepo.FindApplicantByID(id)
	if err != nil {
		return err
	}

	user.Email = request.Email
	user.Name = request.Name
	user.Surname = request.Surname

	return u.ApplicantRepo.UpdateApplicantBasic(user)
}

func (u *ApplicantUsecase) UpdateApplicant(id int, request dto.ApplicantUpdateDTO) error {
	user, err := u.ApplicantRepo.FindApplicantByID(id)
	if err != nil {
		return err
	}
	//Thay doi request DOB ve dang time.Time
	dob, err := time.Parse("2006-01-02", request.DOB)
	if err != nil {
		return err
	}
	user.Name = request.Name
	user.Surname = request.Surname
	user.Gender = request.Gender
	user.DOB = dob
	user.Mobile = request.Mobile
	user.CountryID = request.CountryID
	user.ResidentCountryID = request.ResidentCountryID
	user.DepartmentID = request.DepartmentID

	return u.ApplicantRepo.UpdateApplicant(user)
}

func (u *ApplicantUsecase) DeleteApplicant(id int) error {
	return u.ApplicantRepo.DeleteApplicant(id)
}

func (u *ApplicantUsecase) FindApplicantByID(id int) (*dto.ApplicantResponseDTO, error) {
	user, err := u.ApplicantRepo.FindApplicantByID(id)
	if err != nil {
		return nil, err
	}

	response := &dto.ApplicantResponseDTO{
		ID:                 user.ID,
		Email:              user.Email,
		Name:               user.Name,
		Surname:            user.Surname,
		Gender:             user.Gender,
		DOB:                user.DOB.String(),
		Mobile:             user.Mobile,
		RoleID:             user.RoleID,
		VerificationStatus: user.VerificationStatus,
		CountryID:          user.CountryID,
		ResidentCountryID:  user.ResidentCountryID,
		DepartmentID:       user.DepartmentID,
	}
	return response, nil
}
