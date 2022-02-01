package usecases

import (
	mock_repositories "awcoding.com/back/pkg/data/repositories/mocks"
	"awcoding.com/back/pkg/domain/entities"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetById(t *testing.T) {
	type mockBehavior func(r *mock_repositories.MockUserRepository)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expected     interface{}
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock_repositories.MockUserRepository) {
				r.EXPECT().GetById(1).Return(&entities.User{Email: "test@test.ru"}, nil)
			},
			expected: "test@test.ru",
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := mock_repositories.NewMockUserRepository(ctrl)
			testCase.mockBehavior(r)

			uc := NewUserCases(r)

			usr, err := uc.GetById(1)
			assert.Nil(t, err, "Error user")
			assert.Equalf(t, testCase.expected, usr.Email, "Error user.Email, expected %s, got %s", testCase.expected, usr.Email)
		})
	}
}

func TestGetByEmailPassword(t *testing.T) {
	type mockBehavior func(r *mock_repositories.MockUserRepository)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expected     func(usr *entities.User, err error)
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock_repositories.MockUserRepository) {
				r.EXPECT().GetByEmailPassword("test@test.ru", "123").Return(&entities.User{Email: "test@test.ru"}, nil)
			},
			expected: func(usr *entities.User, err error) {
				assert.Nil(t, err, "Error user")
				assert.Equalf(t, "test@test.ru", usr.Email, "Error user.Email, expected %s, got %s", "test@test.ru", usr.Email)
			},
		},
		{
			name: "NotFound",
			mockBehavior: func(r *mock_repositories.MockUserRepository) {
				r.EXPECT().GetByEmailPassword("test@test.ru", "123").Return(nil, errors.New("Server error"))
			},
			expected: func(usr *entities.User, err error) {
				assert.NotNil(t, err, "Error expected error, got nil")
				assert.Equalf(t, "Server error", err.Error(), "Error, expected %s, got %s", "Server error", err.Error())
			},
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := mock_repositories.NewMockUserRepository(ctrl)
			testCase.mockBehavior(r)

			uc := NewUserCases(r)

			usr, err := uc.GetByEmailPassword("test@test.ru", "123")
			testCase.expected(usr, err)
		})
	}
}
