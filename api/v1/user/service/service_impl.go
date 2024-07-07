package service

import (
	"bioskuy/api/v1/user/repository"
	"bioskuy/auth"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	Repo     repository.UserRepository
	Validate *validator.Validate
	DB *sql.DB
	Jwt auth.Auth
}



// func NewUserService(repo repository.UserRepository, validate *validator.Validate, DB *sql.DB) UserService {
// 	return &userService{repo: repo, validate: validate, DB: DB}
// }

// func (s *userService) Register(ctx context.Context, request dto.CreateUserRequest, c *gin.Context) (dto.UserResponseLoginAndRegister, error){
// 	var UserResponse = dto.UserResponseLoginAndRegister{}

// 	err := s.Validate.Struct(request)
// 	if err != nil {
// 		c.Error(exception.ValidationError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
// 		return  UserResponse, err
// 	}

// 	tx, err := s.DB.Begin()
// 	if err != nil {
// 		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
// 		return  UserResponse, err
// 	}
// 	defer helper.CommitAndRollback(tx, c)


// 	user := entity.User{
// 		Name: request.Name,
// 		Email: request.Email,
// 	}

// 	result, err := s.Repo.Save(ctx, tx, user, c)
// 	if err != nil {
// 		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
// 		return  UserResponse, err
// 	}

// 	Token, err := s.Jwt.GenerateToken(result, c)
// 	if err != nil {
// 		c.Error(exception.InternalServerError{Message: err.Error()}).SetType(gin.ErrorTypePublic)
// 		return  UserResponse, err
// 	}

// 	UserResponse.Token = Token

// 	return UserResponse, nil
// }

// func (s *userService) FindByEmail(ctx context.Context, email string, c *gin.Context) (dto.UserResponse, error){
// 	user, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return dto.UserResponse{}, err
// 	}
// 	return dto.UserResponse{
// 		ID:    user.ID,
// 		Name:  user.Name,
// 		Email: user.Email,
// 	}, nil
// }

// func (s *userService) GetAllUsers() ([]dto.UserResponse, error) {
// 	users, err := s.repo.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var userResponses []dto.UserResponse
// 	for _, user := range users {
// 		userResponses = append(userResponses, dto.UserResponse{
// 			ID:    user.ID,
// 			Name:  user.Name,
// 			Email: user.Email,
// 		})
// 	}
// 	return userResponses, nil
// }

// func (s *userService) UpdateUser(id string, request dto.UpdateUserRequest) (dto.UserResponse, error) {
// 	if err := s.validate.Struct(request); err != nil {
// 		return dto.UserResponse{}, err
// 	}
// 	user, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return dto.UserResponse{}, err
// 	}
// 	user.Name = request.Name
// 	user.Email = request.Email
// 	user, err = s.repo.Update(user)
// 	if err != nil {
// 		return dto.UserResponse{}, err
// 	}
// 	return dto.UserResponse{
// 		ID:    user.ID,
// 		Name:  user.Name,
// 		Email: user.Email,
// 	}, nil
// }

// func (s *userService) DeleteUser(id string) error {
// 	user, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return err
// 	}
// 	return s.repo.Delete(user.ID)
// }
