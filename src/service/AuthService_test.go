package service

import (
	"context"
	"fmt"
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
				fmt.Print("hello")
				Expect(err).Should(BeNil())
				Expect(userDetails).Should(Equal(apiModel.User{
					Name:     "name",
					Username: "username",
					Email:    "email",
				}))
			})
		})
	})
})
