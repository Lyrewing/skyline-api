package services

import (
	"skyline-api/models"
	"skyline-api/repositories"
)

type UserService struct {
}

var userRepository repositories.IUserRepository

func init() {
	userRepository = repositories.NewUserRepository()
}

func (s *UserService) Login(email string, password string) bool {
	user := &models.User{Name: "", Password: ""}
	if userRepository.QueryByStruct(user) != nil {
		return true
	} else {
		return false
	}
}

func (s *UserService) Add(user *models.User) {
	userRepository.Create(user)
}

func (s *UserService) GetUsers(index uint, size uint, users *[]models.User) int {
	total := 0
	users, total = userRepository.QueryPage(index, size)
	return total

}
