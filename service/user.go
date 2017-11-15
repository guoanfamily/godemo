package service

import (
	"time"
	"godemo/cache"
)

type SubObject struct {
	isCache bool
	Id string
}

type User struct {
	SubObject
	Name         string
	Username     string
	Password     string
	Phone        string
	Sex          string
	Status       string
	CredentialNo string
	Email        string
	AliAccount   string
	Portrait     string
	//OrgId        string
	//Introduce    string
	CreateTime   time.Time
	//CityId       string
	//ValidFlag    string
}

//func (s *SubObject) AfterFind() (err error) {
//	fmt.Println("AfterFind")
//	return
//}
type UserRole struct {
	Username string
	Rolename string
}

func Frist() []UserRole{
	//var users []User

	//var user User
	//user.isCache = true
	//users = append(users,user)
	sql := "select u.`name` username,r.`name` rolename from usertable u LEFT JOIN user_role ur on u.id= ur.user_id LEFT JOIN role r on ur.role_id=r.id where true=true"
	var userroles []UserRole
	cache.CacheSelect(&userroles,sql)
	return userroles
}