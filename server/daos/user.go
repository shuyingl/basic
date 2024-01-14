package daos

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDao interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(filter *models.User) (*models.User, error)
	GetUserWithId(userId string) (*models.User, error)
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{
		db: db,
	}
}

func (ud *userDao) CreateUser(user *models.User) (*models.User, error) {
	user.ID = uuid.New()

	return user, ud.db.Create(user).Error
}

func (ud *userDao) GetUser(filter *models.User) (*models.User, error) {
	var user *models.User
	err := ud.db.Where(filter).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ud *userDao) GetUserWithId(userId string) (*models.User, error) {
	var user *models.User
	err := ud.db.First(&user, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
