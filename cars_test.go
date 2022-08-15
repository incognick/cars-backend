package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CarSuite struct {
	suite.Suite
	DB     *gorm.DB
	router *chi.Mux
	mock   sqlmock.Sqlmock
}

func (s *CarSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.DB.Logger.LogMode(logger.Info)

	s.router = createRouter(s.DB)
}

func TestCarSuite(t *testing.T) {
	suite.Run(t, new(CarSuite))
}

func (suite *CarSuite) TestCreateNotFound() {
	body := `malformed`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusNotFound, res.StatusCode)
}

func (suite *CarSuite) TestCreateInvalidJSON() {
	body := `malformed`
	req := httptest.NewRequest(http.MethodPost, "/cars/foo", strings.NewReader(body))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusBadRequest, res.StatusCode)
}

func (suite *CarSuite) TestCreateMissingFields() {
	body := `
	{
		"make": "Ford"
	}
	`
	req := httptest.NewRequest(http.MethodPost, "/cars/foo", strings.NewReader(body))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}

func (suite *CarSuite) TestCreateInvalidCategory() {
	body := `
	{
		"make": "Tesla",
		"model": "3",
		"package": "Performance",
		"color": "White",
		"year": 2023,
		"category": "Starship",
		"mileage": 1,
		"price_cents": 10000000
	}
	`
	req := httptest.NewRequest(http.MethodPost, "/cars/3i", strings.NewReader(body))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}

func (suite *CarSuite) TestCreateSuccess() {

	suite.mock.ExpectExec(`INSERT INTO "cars"`).
		WithArgs("3i", "Tesla", "3", "Performance", "White", 2023, "Sedan", 1, 10000000).
		WillReturnResult(sqlmock.NewResult(1, 1))

	body := `
	{
		"make": "Tesla",
		"model": "3",
		"package": "Performance",
		"color": "White",
		"year": 2023,
		"category": "Sedan",
		"mileage": 1,
		"price_cents": 10000000
	}
	`
	req := httptest.NewRequest(http.MethodPost, "/cars/3i", strings.NewReader(body))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusCreated, res.StatusCode)
}
