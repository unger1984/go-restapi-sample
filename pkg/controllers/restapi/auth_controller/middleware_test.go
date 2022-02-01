package auth_controller

import (
	"awcoding.com/back/pkg/domain/entities"
	mocks "awcoding.com/back/pkg/domain/usecases/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthController_VerifyJWT(t *testing.T) {
	type mockBehavior func(s *mocks.MockAuthCases, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mocks.MockAuthCases, token string) {
				s.EXPECT().GetByToken(token).Return(&entities.User{}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "1",
		},
		{
			name:        "Unauthorized",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mocks.MockAuthCases, token string) {
				s.EXPECT().GetByToken(token).Return(nil, errors.New("invalid token"))
			},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
		{
			name:        "Empty Header",
			headerName:  "Authorization",
			headerValue: "",
			token:       "",
			mockBehavior: func(s *mocks.MockAuthCases, token string) {
				// Skip call
			},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: `{"message":"empty auth_controller header"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mocks.NewMockAuthCases(c)
			testCase.mockBehavior(auth, testCase.token)
			authController := NewAuthController(auth)

			r := gin.New()

			apiGroup := r.Group("/api")
			apiGroup.Use(authController.VerifyJWT)

			apiGroup.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, 1)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/test", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)
			assert.Equalf(t, testCase.expectedStatusCode, w.Code, "Error code, expected %d, got %d", testCase.expectedStatusCode, w.Code)
			assert.Equalf(t, testCase.expectedResponseBody, w.Body.String(), "Error body, expected %s, got %s", testCase.expectedResponseBody, w.Body.String())
		})
	}

}
