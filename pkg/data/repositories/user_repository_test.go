package repositories

import (
	"awcoding.com/back/pkg/domain/entities"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

type UserRepositorySuite struct {
	suite.Suite
	db       *gorm.DB
	dbMocker sqlmock.Sqlmock
}

func (s *UserRepositorySuite) SetupTest() {}

func (s *UserRepositorySuite) BeforeTest(suiteName, testName string) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	s.dbMocker = mock
	s.db = gdb
}

func (s *UserRepositorySuite) AfterTest(_, _ string) {
	db, err := s.db.DB()
	if err != nil {
		panic(err)
	}
	err2 := db.Close()
	if err2 != nil {
		return
	}
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

func (s *UserRepositorySuite) TestUserRepository_GetById() {
	type mock func(m sqlmock.Sqlmock, rows *sqlmock.Rows)
	type test func(t *testing.T, u *entities.User, err error)

	columns := []string{"id", "email", "avatarId", "password", "created_at", "updated_at", "Avatar__id", "Avatar__path", "Avatar__name", "Avatar__type", "Avatar__created_at", "Avatar__updated_at"}

	testTable := []struct {
		name string
		rows *sqlmock.Rows
		mock mock
		test test
	}{
		{
			name: "OK",
			rows: sqlmock.NewRows(columns).AddRow(1, "test@test.ru", 1,
				"123123", time.Now(), time.Now(), 1, "/test.jpg", "Test", "image/jpeg",
				time.Now(), time.Now()),
			mock: func(m sqlmock.Sqlmock, rows *sqlmock.Rows) {
				m.ExpectQuery(`^SELECT (.+) FROM "User" LEFT JOIN "Upload" "Avatar" ON (.+) WHERE`).
					WithArgs(1).WillReturnRows(rows)

			},
			test: func(t *testing.T, u *entities.User, err error) {
				assert.Nilf(t, err, "Error %v", err)
				assert.Equalf(t, "test@test.ru", u.Email, "Error email, expected %s, got %s", "test@test.ru", u.Email)
				assert.Equalf(t, "/test.jpg", u.Avatar.Path, "Error email, expected %s, got %s", "test@test.ru", "/test.jpg", u.Avatar.Path)
			},
		},
		{
			name: "NotFound",
			rows: nil,
			mock: func(m sqlmock.Sqlmock, rows *sqlmock.Rows) {
				m.ExpectQuery(`^SELECT (.+) FROM "User" LEFT JOIN "Upload" "Avatar" ON (.+) WHERE`).
					WithArgs(1).WillReturnRows(sqlmock.NewRows(columns))

			},
			test: func(t *testing.T, u *entities.User, err error) {
				assert.Error(t, err, "No error")
				assert.Equalf(t, "record not found", err.Error(), "Wrong error, expected \"%s\", got \"%s\"", errors.New("record not found"), err)
			},
		},
	}

	for _, testCase := range testTable {
		s.T().Run(testCase.name, func(t *testing.T) {
			testCase.mock(s.dbMocker, testCase.rows)
			r := NewUserRepository(s.db)

			user, err := r.GetById(1)
			testCase.test(t, user, err)
		})
	}
}

func (s *UserRepositorySuite) TestUserRepository_GetByEmailPassword() {
	type mock func(m sqlmock.Sqlmock, rows *sqlmock.Rows)
	type test func(t *testing.T, u *entities.User, err error)

	columns := []string{"id", "email", "avatarId", "password", "created_at", "updated_at", "Avatar__id", "Avatar__path", "Avatar__name", "Avatar__type", "Avatar__created_at", "Avatar__updated_at"}

	testTable := []struct {
		name string
		rows *sqlmock.Rows
		mock mock
		test test
	}{
		{
			name: "OK",
			rows: sqlmock.NewRows(columns).AddRow(1, "test@test.ru", 1,
				"123123", time.Now(), time.Now(), 1, "/test.jpg", "Test", "image/jpeg",
				time.Now(), time.Now()),
			mock: func(m sqlmock.Sqlmock, rows *sqlmock.Rows) {
				m.ExpectQuery(`^SELECT (.+) FROM "User" LEFT JOIN "Upload" "Avatar" ON (.+) WHERE`).
					WithArgs("test@test.ru", "1234").WillReturnRows(rows)

			},
			test: func(t *testing.T, u *entities.User, err error) {
				assert.Nilf(t, err, "Error %v", err)
				assert.Equalf(t, "test@test.ru", u.Email, "Error email, expected %s, got %s", "test@test.ru", u.Email)
				assert.Equalf(t, "/test.jpg", u.Avatar.Path, "Error email, expected %s, got %s", "test@test.ru", "/test.jpg", u.Avatar.Path)
			},
		},
		{
			name: "NotFound",
			rows: nil,
			mock: func(m sqlmock.Sqlmock, rows *sqlmock.Rows) {
				m.ExpectQuery(`^SELECT (.+) FROM "User" LEFT JOIN "Upload" "Avatar" ON (.+) WHERE`).
					WithArgs("test@test.ru", "1234").WillReturnRows(sqlmock.NewRows(columns))

			},
			test: func(t *testing.T, u *entities.User, err error) {
				assert.Error(t, err, "No error")
				assert.Equalf(t, "login and password incorrect", err.Error(), "Wrong error, expected \"%s\", got \"%s\"", errors.New("login and password incorrect"), err)
			},
		},
	}

	for _, testCase := range testTable {
		s.T().Run(testCase.name, func(t *testing.T) {
			testCase.mock(s.dbMocker, testCase.rows)
			r := NewUserRepository(s.db)

			user, err := r.GetByEmailPassword("test@test.ru", "1234")
			testCase.test(t, user, err)
		})
	}
}
