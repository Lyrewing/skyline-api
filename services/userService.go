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
