package service

import (
	"time"
	"godemo/common"
	"fmt"
	"database/sql"
)

type SubObject struct {
	isCache bool
	Id string
}

type Usertable struct {
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
	name string
	Rolename string
	roleid string
	Sex int
	Status bool
	CreateTime time.Time
}

func DeleteUser() bool{
	var user Usertable
	user.Id = "1"
	common.Db.Delete(&user)

	return true
}
func QueryUserTable() interface{}{
	sql := "select * from usertable"
	return common.Db.QuerybySql(sql)
}
func Select() []UserRole{

	sql := "select u.`name` username,u.sex sex,u.status status,r.`name` rolename,u.create_time create_time from usertable u LEFT JOIN user_role ur on u.id= ur.user_id LEFT JOIN role r on ur.role_id=r.id"
	var userroles []UserRole

	common.Db.QueryStruct(&userroles,sql)

	return userroles
}
type AccountList struct {
	Id string
	Name sql.NullString
	Simple_spell sql.NullString
}
func Acclist() []AccountList{
	var acc []AccountList
	sql := "SELECT * FROM city"
	err := common.Db.Select(&acc,sql)
	if err!=nil{
		fmt.Println(err)
	}
	return acc
}

func Save() bool{

	tx := common.Db.MustBegin()
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("Recovered in testPanic2Error", r)
			tx.Rollback()
		}
	}()
	tx.MustExec("insert INTO usertable (id,`name`) values(3,'wang')")
	tx.MustExec("insert INTO usertable (id,`name`) values(4,'wang')")
	tx.MustExec("insert INTO usertable (id,`name`) values(3,'wang')")
	tx.Commit()
	return true
}