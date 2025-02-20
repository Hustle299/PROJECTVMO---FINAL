package usecase

import (
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	"github.com/golang-jwt/jwt/v4"
)

type UserUsecaseInterface interface {
	Login(req dto.LoginUserRequest) (*dto.LoginUserTokenResponse, string)
	RegisterUser(req dto.RegisterUserRequest) (*dto.RegisterUserResponse, string)
}

type UserUsecase struct {
	repo      storage.AuthenticationSrore
	secretKey string
}

func NewUserUsecase(repo storage.AuthenticationSrore, secretKey string) *UserUsecase {
	return &UserUsecase{repo: repo,
		secretKey: secretKey}
}
func (u *UserUsecase) Login(req dto.LoginUserRequest) (*dto.LoginUserTokenResponse, string) {
	user, msg := u.repo.GetUserByEmail(req.Email, req.Password)
	if user != nil {
		//Create claim
		claims := jwt.MapClaims{
			"userId": user.ID,
			"roleId": user.RoleID,
			"exp":    time.Now().Add(time.Hour * 72).Unix(),
		}
		//Create token with claim and secret key
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//Sign token with key
		tokenString, err := token.SignedString([]byte(u.secretKey))
		if err != nil {
			return nil, "Could not generate token"
		}
		//Login success return token to user
		return &dto.LoginUserTokenResponse{
			Token: tokenString,
		}, ""
	}
	return nil, msg
}

func (u *UserUsecase) RegisterUser(req dto.RegisterUserRequest) (*dto.RegisterUserResponse, string) {
	// check existed user
	user, _ := u.repo.GetUserByEmail(req.Email, "")
	if user != nil {
		return nil, "User existed"
	}
	// register user
	registerUser, err := u.repo.RegisterUser(&req)
	if err != nil {
		return nil, "Register failed"
	}

	return registerUser, ""
}
