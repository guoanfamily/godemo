package main

import (
	 "github.com/guoanfamily/sqlx"
	 _ "github.com/go-sql-driver/mysql"
	 "godemo/common"
	"fmt"
	 "godemo/router"
	 _ "github.com/go-redis/redis"
)


func init() {
	var err error
	common.Db,err = sqlx.Connect("mysql", "root:@tcp(172.16.4.12:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	common.Db.SetMaxIdleConns(10)
	common.Db.SetMaxOpenConns(100)
	common.Db.Ping()
	fmt.Println("init")
	//redis init
	//common.Rds = redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//
	//pong, err := common.Rds.Ping().Result()
	//fmt.Println(pong, err)
}

//type St struct {
//	Dt time.Time
//	Str string
//}
//type User struct {
//	Name string
//	Age  int
//	Id   string
//}
func main(){
	//tonydon := &User{"TangXiaodong", 100, "0000123"}
	//
	//object := reflect.ValueOf(tonydon)
	//myref := object.Elem()
	//typeOfType := myref.Type()
	//for i:=0; i<myref.NumField(); i++{
	//	field := myref.Field(i)
	//	fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
	//}
	router.Router()
}

