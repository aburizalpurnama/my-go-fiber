package repository

import (
	"github.com/aburizalpurnama/my-go-fiber/model/domain"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(user domain.User) (id string, err error)
	}

	userRepositoryImpl struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) CreateUser(user domain.User) (id string, err error) {
	result := r.db.Create(&user)
	id = user.Id
	err = result.Error
	if err != nil {
		log.Infof("Failed to create user, err:%s", err)
		return
	}

	return
}
