package models

type User struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

//获取第一个元素
func (user *User) First() *User {
	DB.First(user)
	return user
}

//获取所有的元素
func (user *User) All() []User {
	var users []User
	DB.Find(&users)
	return users
}

//根据用户名获取用户
func (user *User) FindByName(name string) *User {
	DB.Where("name=?", name).First(user)
	return user
}
