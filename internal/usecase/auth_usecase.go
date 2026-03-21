package usecase

import (
	"errors"
	"time"

	"golang-clean-arch/internal/entity"
	"golang-clean-arch/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(string, string) (string, error)
	Register(string, string, string) error
}

type authUsecase struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(repo repository.UserRepository, jwtSecret string) AuthUsecase {
	return &authUsecase{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (auth *authUsecase) Login(email string, password string) (string, error) {
	user, err := auth.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(auth.jwtSecret))
	if err != nil {
		return "", errors.New("failed to signing token")
	}

	return signed, nil
}

func (auth *authUsecase) Register(username string, email string, password string) error {
	findUser, _ := auth.repo.GetByEmail(email)
	if findUser != nil {
		return errors.New("email already exist")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash a password")
	}

	user := &entity.User{
		Username: username,
		Email:    email,
		Password: string(hashedPass),
	}

	err = auth.repo.Create(user)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
