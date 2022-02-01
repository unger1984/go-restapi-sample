package usecases

import (
	mock_repositories "awcoding.com/back/pkg/data/repositories/mocks"
	"awcoding.com/back/pkg/domain/entities"
	"awcoding.com/back/pkg/infrastructure/config"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignIn(t *testing.T) {
	type mockBehavior func(r *mock_repositories.MockUserRepository)

	testTable := []struct {
		name         string
		login        string
		password     string
		mockBehavior mockBehavior
		expected     func(auth *entities.Auth, err error)
	}{
		{
			name:     "OK",
			login:    "test@test.ru",
			password: "password",
			mockBehavior: func(r *mock_repositories.MockUserRepository) {
				hashedPassword := NewAuthCases(r, &config.Config{}).generatePasswordHash("password")
				r.EXPECT().GetByEmailPassword("test@test.ru", hashedPassword).Return(&entities.User{Email: "test@test.ru"}, nil)
			},
			expected: func(auth *entities.Auth, err error) {
				assert.Nil(t, err, "Error auth")
				assert.Regexpf(t, "^eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.", auth.Token,
					"Error user.Email, expected %s, got %s", "^eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.",
					auth.Token,
				)
			},
		},
		{
			name:     "Incorrect",
			login:    "test@test.ru",
			password: "password",
			mockBehavior: func(r *mock_repositories.MockUserRepository) {
				hashedPassword := NewAuthCases(r, &config.Config{}).generatePasswordHash("password")
				r.EXPECT().GetByEmailPassword("test@test.ru", hashedPassword).Return(nil, errors.New("Some error"))
			},
			expected: func(auth *entities.Auth, err error) {
				assert.NotNil(t, err, "Error auth")
				assert.Equalf(t, "login and password incorrect", err.Error(),
					"Error expected %s, got %s",
					"login and password incorrect", err.Error(),
				)
			},
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := mock_repositories.NewMockUserRepository(ctrl)
			testCase.mockBehavior(r)

			ac := NewAuthCases(r, &config.Config{})

			auth, err := ac.SignIn(testCase.login, testCase.password)
			testCase.expected(auth, err)
		})
	}
}

func TestGetByToken(t *testing.T) {
	type mockBehavior func(r *mock_repositories.MockUserRepository)

	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDM4MDA1NzksImlhdCI6MTY0Mzc1NzM3OSwidXNlcl9pZCI6MH0.abeSkdD2A79C9YujQ9mzAljlTbHCZI_ixPmRQb-qhIY"

	testTable := []struct {
		name         string
		token        string
		mockBehavior mockBehavior
		expected     func(usr *entities.User, err error)
	}{
		{
			name:  "OK",
			token: token,
			mockBehavior: func(r *mock_repositories.MockUserRepository) {
				r.EXPECT().GetById(0).Return(&entities.User{Email: "test@test.ru"}, nil)
			},
			expected: func(usr *entities.User, err error) {
				assert.Nil(t, err, "Error auth")
				assert.Equalf(t, "test@test.ru", usr.Email,
					"Error user.Email, expected %s, got %s",
					"test@test.ru", usr.Email,
				)
			},
		},
		{
			name:  "Validation",
			token: "asdsadasdasdasd",
			mockBehavior: func(r *mock_repositories.MockUserRepository) {
				//r.EXPECT().GetById(1).Return(&entities.User{Email: "test@test.ru"}, nil)
			},
			expected: func(usr *entities.User, err error) {
				assert.NotNil(t, err, "Error auth")
				assert.Equalf(t, "token contains an invalid number of segments", err.Error(),
					"Error user.Email, expected %s, got %s",
					"token contains an invalid number of segments", err.Error(),
				)
			},
		},
		//{
		//	name:     "Incorrect",
		//	login:    "test@test.ru",
		//	password: "password",
		//	mockBehavior: func(r *mock_repositories.MockUserRepository) {
		//		hashedPassword := NewAuthCases(r, &config.Config{}).generatePasswordHash("password")
		//		r.EXPECT().GetByEmailPassword("test@test.ru", hashedPassword).Return(nil, errors.New("Some error"))
		//	},
		//	expected: func(auth *entities.Auth, err error) {
		//		assert.NotNil(t, err, "Error auth")
		//		assert.Equalf(t, "login and password incorrect", err.Error(),
		//			"Error expected %s, got %s",
		//			"login and password incorrect", err.Error(),
		//		)
		//	},
		//},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := mock_repositories.NewMockUserRepository(ctrl)
			testCase.mockBehavior(r)

			ac := NewAuthCases(r, &config.Config{})

			usr, err := ac.GetByToken(testCase.token)
			testCase.expected(usr, err)
		})
	}
}
