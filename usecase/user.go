package usecase

import (
	"fmt"

	"github.com/aburizalpurnama/my-go-fiber/model/domain"
	"github.com/aburizalpurnama/my-go-fiber/model/request"
	"github.com/aburizalpurnama/my-go-fiber/model/respose"
	"github.com/aburizalpurnama/my-go-fiber/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecase interface {
		CreateUser(userReq request.CreateUser) (resp respose.Default, err error)
	}

	userUsecaseImpl struct {
		userRepo repository.UserRepository
	}
)

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo}
}

func (u *userUsecaseImpl) CreateUser(userReq request.CreateUser) (resp respose.Default, err error) {
	var user domain.User
	user.Email = userReq.Email

	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	user.Password = string(hashedBytes)

	id, err := u.userRepo.CreateUser(user)
	if err != nil {
		resp.Code = "69"
		resp.Message = fmt.Sprintf("Failed to create user, err:%s", err)
		return
	}

	resp.Data = respose.UserData{Id: id}
	return
}
