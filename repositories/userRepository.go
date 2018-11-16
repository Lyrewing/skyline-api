package repositories

import (
	"skyline-api/models"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	Exec(sql string) bool
	Create(user *models.User)
	DeleteById(uint)
	FindById(uint) *models.User
	FindAll() *[]models.User
	QueryByStruct(*models.User) *[]models.User
	QueryByMap(queryMap map[string]interface{}) *[]models.User
	QueryPage(index uint, size uint) (*[]models.User, uint)
}

type UserRepository struct {
	DbContext *gorm.DB
}

//创建实例
func NewUserRepository() *UserRepository {
	return &UserRepository{DbContext: DB}
}

//执行Sql语句
func (r *UserRepository) Exec(sql string) bool {
	return true
}

//创建user
func (r *UserRepository) Create(user *models.User) {
	r.DbContext.Create(user)
}

//删除通过Id
func (r *UserRepository) DeleteById(id uint) {
	user := &models.User{ID: id}
	r.DbContext.Delete(user)
}

//查询By Id
func (r *UserRepository) FindById(id uint) *models.User {
	user := &models.User{}
	r.DbContext.Where("id = ?", id).Find(user)
	return user
}

func (r *UserRepository) FindAll() *[]models.User {
	users := []models.User{}
	r.DbContext.Find(users)
	return &users
}

func (r *UserRepository) QueryByStruct(user *models.User) *[]models.User {
	users := &[]models.User{}
	r.DbContext.Where(user).Find(users)
	return users
}

func (r *UserRepository) QueryByMap(queryMap map[string]interface{}) *[]models.User {
	users := &[]models.User{}
	r.DbContext.Where(queryMap).Find(users)
	return users
}

func (r *UserRepository) QueryPage(index uint, size uint) (*[]models.User, uint) {
	users := &[]models.User{}
	var count uint
	r.DbContext.Model(&models.User{}).Count(&count)
	r.DbContext.Limit(size).Offset(size * (index - 1)).Order("id desc").Find(users)
	return users, count
}
