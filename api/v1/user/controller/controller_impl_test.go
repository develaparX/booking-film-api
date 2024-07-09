package controller

import (
	"bioskuy/api/v1/user/service"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewUserController(t *testing.T) {
	type args struct {
		userService service.UserService
	}
	tests := []struct {
		name string
		args args
		want UserController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userController_LoginWithGoogle(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctl  *userController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctl.LoginWithGoogle(tt.args.c)
		})
	}
}

func Test_userController_CallbackFromGoogle(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctl  *userController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctl.CallbackFromGoogle(tt.args.c)
		})
	}
}

func Test_userController_GetUserByID(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctl  *userController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctl.GetUserByID(tt.args.c)
		})
	}
}

func Test_userController_GetAllUsers(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctl  *userController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctl.GetAllUsers(tt.args.c)
		})
	}
}

func Test_userController_UpdateUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctl  *userController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctl.UpdateUser(tt.args.c)
		})
	}
}

func Test_userController_DeleteUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctl  *userController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctl.DeleteUser(tt.args.c)
		})
	}
}
