package controller

import (
	"bytes"
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

func TestAuthController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Controller test")
}

var _ = Describe("Auth Controller Test", func() {
	var (
		authController AuthController
		authService    *mocks.MockAuthService
		tempContext    *gin.Context
		recorder       *httptest.ResponseRecorder
		mockController *gomock.Controller
	)
	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		mockController = gomock.NewController(GinkgoT())
		tempContext, _ = gin.CreateTestContext(recorder)
		authService = mocks.NewMockAuthService(mockController)
		authController = NewAuthController(authService)
	})
	AfterEach(func() {
		mockController.Finish()
	})
	Context("Login Controller", func() {
		When("service returns success", func() {
			It("should return user details", func() {
				req := `{
					"email": "email",
					"password": "password"
				}`
				serviceReq := apiModel.Login{"email", "password"}

				tempContext.Request = httptest.NewRequest("POST", "/login", bytes.NewBufferString(req))

				authService.EXPECT().Login(tempContext, serviceReq).
					Return(apiModel.User{
						Name:     "name",
						Username: "username",
						Email:    "email",
					}, nil).Times(1)

				authController.Login(tempContext)

				expectedResponse := `{"name":"name","username":"username","email":"email"}`

				fmt.Print(recorder.Body.String())
				Expect(recorder.Code).Should(Equal(http.StatusOK))
				Expect(recorder.Body.String()).Should(Equal(expectedResponse))
			})
		})
	})
})
