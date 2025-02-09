package repositories

import (
	"github.com/yaghoubi-mn/voter/internal/custom_errors"
	"github.com/yaghoubi-mn/voter/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(username string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByUsername(username string) (models.User, error) {

	var user models.User
	if err := r.db.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, custom_errors.RecordNotFound
		}
		return user, err
	}

	return user, nil
}
