package service

import (
	"context"
	"harmony/src/mocks"
	apiModel "harmony/src/models/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuthService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Service test")
}

var _ = Describe("Auth Service test", func() {
	var (
		mockController     *gomock.Controller
		authService        AuthService
		mockAuthRepository *mocks.MockAuthRepository
		tempContext        *gin.Context
		recorder           *httptest.ResponseRecorder
	)
	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		mockAuthRepository = mocks.NewMockAuthRepository(mockController)
		authService = NewAuthService(mockAuthRepository)
		recorder = httptest.NewRecorder()
		tempContext, _ = gin.CreateTestContext(recorder)
		tempContext.Request, _ = http.NewRequest("GET", "url", nil)
	})
	AfterEach(func() {
		mockController.Finish()
	})
	Context("Login Service", func() {
		When("db query is successful", func() {
			It("should return user details", func() {
				mockAuthRepository.EXPECT().Login(context.Background(), "email", "password").
					Return(apiModel.User{
						Name:     "name",
						Username: "username",
						Email:    "email",
					}, nil).Times(1)

				userDetails, err := authService.Login(tempContext, apiModel.Login{Email: "email", Password: "password"})
				Expect(err).Should(BeNil())
				Expect(userDetails).Should(Equal(apiModel.User{
					Name:     "name",
					Username: "username",
					Email:    "email",
				}))
			})
		})
	})
	Context("Register Service", func() {
		When("db query is successful", func() {
			It("should return name", func() {
				mockAuthRepository.EXPECT().Register(context.Background(), apiModel.Register{
					Name:            "name",
					Username:        "username",
					Email:           "email",
					Password:        "password",
					ConfirmPassword: "password",
				}).Return(nil).Times(1)

				name, err := authService.Register(tempContext, apiModel.Register{
					Name:            "name",
					Username:        "username",
					Email:           "email",
					Password:        "password",
					ConfirmPassword: "password",
				})
				Expect(err).Should(BeNil())
				Expect(name).Should(Equal("name"))
			})
		})
	})
})
