package auth_controller

import (
	"awcoding.com/back/pkg/domain/entities"
	mocks "awcoding.com/back/pkg/domain/usecases/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthController_signIn(t *testing.T) {
	type mockBehavior func(s *mocks.MockAuthCases, login string, password string)

	tastTable := []struct {
		name                 string
		inputBody            string
		input                SignInInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"login":"login","password":"password"}`,
			input:     SignInInput{Login: "login", Password: "password"},
			mockBehavior: func(s *mocks.MockAuthCases, login string, password string) {
				s.EXPECT().SignIn(login, password).Return(&entities.Auth{Token: "token"}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"token":"token","user":null}`,
		},
		{
			name:      "Unauthorized",
			inputBody: `{"login":"login","password":"password"}`,
			input:     SignInInput{Login: "login", Password: "password"},
			mockBehavior: func(s *mocks.MockAuthCases, login string, password string) {
				s.EXPECT().SignIn(login, password).Return(nil, errors.New("login and password incorrect"))
			},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: `{"message":"login and password incorrect"}`,
		},
		{
			name:      "BadRequest",
			inputBody: `{"login1":"login","password1":"password"}`,
			input:     SignInInput{Login: "login", Password: "password"},
			mockBehavior: func(s *mocks.MockAuthCases, login string, password string) {
				// Skip call
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "InternalServerError",
			inputBody: `{"login":"login","password":"password"}`,
			input:     SignInInput{Login: "login", Password: "password"},
			mockBehavior: func(s *mocks.MockAuthCases, login string, password string) {
				s.EXPECT().SignIn(login, password).Return(nil, errors.New("iternal server error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"iternal server error"}`,
		},
	}

	for _, testCase := range tastTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mocks.NewMockAuthCases(c)
			testCase.mockBehavior(auth, testCase.input.Login, testCase.input.Password)

			authController := NewAuthController(auth)

			r := gin.New()
			authApi := r.Group("/auth")
			authController.NewRoutes(authApi)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)
			assert.Equalf(t, testCase.expectedStatusCode, w.Code, "Error code, expected %d, got %d", testCase.expectedStatusCode, w.Code)
			assert.Equalf(t, testCase.expectedResponseBody, w.Body.String(), "Error body, expected %s, got %s", testCase.expectedResponseBody, w.Body.String())
		})
	}
}
