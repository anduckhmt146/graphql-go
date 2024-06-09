package services

import (
	"github.com/anduckhmt146/graphql-api/internal/model"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers() ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(name string, age int) (*model.User, error)
	UpdateUser(id int, name string, age int) (*model.User, error)
	DeleteUser(id int) (int64, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	var user model.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *userService) CreateUser(name string, age int) (*model.User, error) {
	user := model.User{Name: name, Age: age}
	result := s.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *userService) UpdateUser(id int, name string, age int) (*model.User, error) {
	var user model.User
	s.db.First(&user, id)
	if name != "" {
		user.Name = name
	}
	if age != 0 {
		user.Age = age
	}
	result := s.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *userService) DeleteUser(id int) (int64, error) {
	result := s.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
