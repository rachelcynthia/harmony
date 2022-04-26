package repository

import (
	"context"
	"fmt"
	apiModel "harmony/src/models/api"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuthRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth repository test")
}

const (
	TestLoginQuery = "SELECT NAME, USERNAME, EMAIL FROM USERS WHERE EMAIL = :1 AND PASSWORD = :2"
)

var _ = Describe("Auth Repository Test", func() {
	var (
		authRepository AuthRepository
		dbConnection   *sqlx.DB
		sqlMock        *sqlmock.Sqlmock
		tempContext    context.Context
	)
	BeforeEach(func() {
		dbConn, mock, _ := sqlmock.New()
		dbConnection = sqlx.NewDb(dbConn, "sqlmock")
		sqlMock = &mock
		authRepository = NewAuthRepository(dbConnection)
		tempContext = context.WithValue(context.TODO(), "dummy", "dummy")
	})
	AfterEach(func() {
	})

	Context("Login", func() {
		When("using valid credentials", func() {
			It("returns the user details", func() {
				sqlMock := *sqlMock

				rows := sqlMock.NewRows([]string{"ID", "NAME", "USERNAME", "EMAIL", "PASSWORD"}).
					AddRow("1", "name", "username", "email", "password")

				expectedQuery := fmt.Sprintf(TestLoginQuery)

				sqlMock.ExpectQuery(expectedQuery).WithArgs("email", "password").WillReturnRows(rows)

				expectedDetails := apiModel.User{Name: "name", Username: "username", Email: "email"}

				userDetails, err := authRepository.Login(tempContext, "email", "password")
				Expect(err).Should(BeNil())
				Expect(userDetails).Should(Equal(expectedDetails))
			})
		})
	})
})
