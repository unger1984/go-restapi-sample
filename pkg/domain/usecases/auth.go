package usecases

import (
	"awcoding.com/back/pkg/domain/entities"
	"awcoding.com/back/pkg/infrastructure/config"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//go:generate mockgen --source=auth.go --destination=mocks/auth_mock.go

type AuthCases interface {
	SignIn(login string, password string) (*entities.Auth, error)
	GetByToken(token string) (*entities.User, error)
}

type authCases struct {
	userService UserCases
	cfg         *config.Config
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthCases(userService UserCases, cfg *config.Config) *authCases {
	return &authCases{userService: userService, cfg: cfg}
}

func (s *authCases) SignIn(login string, password string) (*entities.Auth, error) {
	hashedPassword := s.generatePasswordHash(password)
	user, err := s.userService.GetByEmailPassword(login, hashedPassword)
	if err != nil {
		return nil, errors.New("login and password incorrect")
	}
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &entities.Auth{Token: token, User: user}, nil
}

func (s *authCases) GetByToken(token string) (*entities.User, error) {
	id, err := s.parseToken(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userService.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authCases) generateToken(user *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(s.cfg.Secret.TokenKey))
}

func (s *authCases) parseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.Secret.TokenKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *authCases) generatePasswordHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.Secret.Salt)))
}
