package repository

import (
	"context"
	"fmt"
	apiModel "harmony/src/models/api"
	"regexp"
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
	TestLoginQuery    = "SELECT * FROM USERS WHERE EMAIL = $1 AND PASSWORD = $2"
	TestRegisterQuery = "INSERT INTO USERS(NAME, USERNAME, EMAIL, PASSWORD) VALUES ($1, $2, $3, $4)"
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

				rows := sqlMock.NewRows([]string{"id", "name", "username", "email", "password"}).
					AddRow("1", "name", "username", "email", "password")

				expectedQuery := fmt.Sprintf(TestLoginQuery)

				sqlMock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("email", "password").WillReturnRows(rows)

				expectedDetails := apiModel.User{Name: "name", Username: "username", Email: "email"}

				userDetails, err := authRepository.Login(tempContext, "email", "password")
				Expect(err).Should(BeNil())
				Expect(userDetails).Should(Equal(expectedDetails))
			})
		})
		When("context cancelled", func() {
			It("returns error", func() {
				sqlMock := *sqlMock

				rows := sqlMock.NewRows([]string{"id", "name", "username", "email", "password"}).
					AddRow("1", "name", "username", "email", "password")

				expectedQuery := fmt.Sprintf(TestLoginQuery)
				sqlMock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("email", "password").WillReturnRows(rows)

				ctx, cancel := context.WithCancel(context.TODO())
				cancel()

				_, err := authRepository.Login(ctx, "email", "password")
				Expect(err).Should(Not(BeNil()))
			})
		})
	})
	Context("Register", func() {
		When("query is successfull", func() {
			It("returns no error", func() {
				sqlMock := *sqlMock

				expectedQuery := fmt.Sprintf(TestRegisterQuery)

				sqlMock.ExpectExec(regexp.QuoteMeta(expectedQuery)).WithArgs("name", "username", "email", "password").
					WillReturnResult(sqlmock.NewResult(1, 0))

				err := authRepository.Register(tempContext, apiModel.Register{
					Name: "name", Username: "username", Email: "email", Password: "password", ConfirmPassword: "password",
				})
				Expect(err).Should(BeNil())
			})
		})
		When("context cancelled", func() {
			It("returns error", func() {
				sqlMock := *sqlMock

				expectedQuery := fmt.Sprintf(TestRegisterQuery)
				sqlMock.ExpectExec(regexp.QuoteMeta(expectedQuery)).WithArgs("name", "username", "email", "password").
					WillReturnResult(sqlmock.NewResult(1, 0))

				ctx, cancel := context.WithCancel(context.TODO())
				cancel()

				err := authRepository.Register(ctx, apiModel.Register{
					Name: "name", Username: "username", Email: "email", Password: "password", ConfirmPassword: "password",
				})
				Expect(err).Should(Not(BeNil()))
			})
		})
	})
})
