package service_test

import (
	"bioskuy/api/v1/user/dto"
	"bioskuy/api/v1/user/entity"
	"bioskuy/api/v1/user/mock/authmock"
	"bioskuy/api/v1/user/mock/repomock"
	"bioskuy/api/v1/user/service"
	"bioskuy/exception"
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	mockRepo    *repomock.UserRepository
	mockAuth    *authmock.AuthMock
	userService service.UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	// Tentukan path absolut untuk file .env
	envPath := "C:\\enigmacamp\\Incubation Batch #27\\Final-Project\\bioskuy\\.env"
	fmt.Println("Loading .env from:", envPath) // Print the path to ensure it is correct

	// Muat file .env
	err := godotenv.Load(envPath)
	if err != nil {
		suite.T().Fatal("Error loading .env file:", err)
	}

	suite.mockRepo = new(repomock.UserRepository)
	suite.mockAuth = new(authmock.AuthMock)
	validate := validator.New()
	db, _ := sql.Open("postgres", "user=postgres password=12345678 dbname=bioskuy_test sslmode=disable")
	suite.userService = service.NewUserService(suite.mockRepo, validate, db, suite.mockAuth)
}

func (suite *UserServiceTestSuite) TestLogin() {
	ginCtx, _ := gin.CreateTestContext(nil)
	request := dto.CreateUserRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	suite.mockRepo.On("FindByEmail", mock.Anything, mock.Anything, "john.doe@example.com", mock.Anything).Return(entity.User{}, exception.NotFoundError{Message: "User not found"}).Once()
	suite.mockRepo.On("Save", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(entity.User{ID: "new-id", Name: "John Doe", Email: "john.doe@example.com"}, nil).Once()
	suite.mockAuth.On("GenerateToken", mock.Anything, mock.Anything).Return("mock-token", nil).Once()
	suite.mockRepo.On("UpdateToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(entity.User{ID: "new-id", Name: "John Doe", Email: "john.doe@example.com", Token: "mock-token"}, nil).Once()

	response, err := suite.userService.Login(context.Background(), request, ginCtx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "mock-token", response.Token)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
