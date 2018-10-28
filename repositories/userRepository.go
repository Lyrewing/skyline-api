package repositories

import (
	"skyline-api/models"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	Exec(string) bool
	Create(user *models.User) bool
	DeleteById(uint) bool
	FindById(uint) *models.User
	FindAll() *[]models.User
	QueryByStruct(*models.User) *[]models.User
	QueryByMap(queryMap map[string]interface{}) *[]models.User
	QueryPage(index uint, size uint) (*[]models.User, uint)
}

type UserRepository struct {
	DbContext *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DbContext: DB}
}
func (r *UserRepository) Exec(sql string) bool {
	return true
}

func (r *UserRepository) Create(user *models.User) bool {
	return true
}
func (r *UserRepository) DeleteById(id uint) bool {
	return true
}

func (r *UserRepository) FindById(id uint) *models.User {
	return &models.User{ID: 1, Name: "", Age: 12, Sex: 1}
}

func (r *UserRepository) FindAll() *[]models.User {
	users := []models.User{}
	return &users
}

func (r *UserRepository) QueryByStruct(user *models.User) *[]models.User {
	users := []models.User{}
	return &users
}

func (r *UserRepository) QueryByMap(queryMap map[string]interface{}) *[]models.User {
	users := []models.User{}
	return &users
}

func (r *UserRepository) QueryPage(index uint, size uint) (*[]models.User, uint) {
	users := []models.User{}
	return &users, 100
}
