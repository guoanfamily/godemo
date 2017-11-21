package main

import (
	 "github.com/jmoiron/sqlx"
	 _ "github.com/go-sql-driver/mysql"
	 "godemo/common"
	"fmt"
	 "godemo/router"
	 "github.com/go-redis/redis"
	"time"
	"reflect"
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
	//redis init
	common.Rds = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := common.Rds.Ping().Result()
	fmt.Println(pong, err)
}

type St struct {
	dt time.Time
	str string
}

func main(){
	var st St
	rt:= reflect.TypeOf(st.dt)
	//rk := rt.Kind()
	fmt.Println(rt.Name())
	router.Router()
}

