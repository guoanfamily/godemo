package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"godemo/common"
	"fmt"
	"godemo/router"
	"regexp"
	"github.com/go-redis/redis"
)

type Point struct {
	X,Y float64
}

func (p Point)Distance() float64{
	return p.X+p.Y
	//return math.Hypot(q.X-p.X, q.Y-p.Y)
}
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


func main(){
	//r:=regexp.MustCompile(`\s+from\s+(\w+)\s+`)
	r2:= regexp.MustCompile("(?i)\\s+(from|join)\\s+`*(\\w+)`*\\s+")
	str :="select u.`name` username,r.`name` rolename from `user` u LEFT JOIN user_role ur on u.id= ur.user_id LEFT JOIN role r on ur.role_id=r.id"
	//str = strings.ToLower(str)
	//tablesa := r.FindAllStringSubmatch(str,-1)
	tablesb := r2.FindAllStringSubmatch(str,-1)

	fmt.Println(tablesb)
	fmt.Println(len(tablesb))
	for _,name:=range tablesb {
		fmt.Println(name[2])
	}

		//var p = new(Point)
	//p.X = 0.1
	//p.Y = 2
	//d :=p.Distance()
	//fmt.Println(d)

	//ch := make(chan int)
	//
	//go func(){
	//	for x:=0;x<10;x++{
	//		fmt.Println("gofunc");
	//		ch <-x;
	//
	//	}
	//	time.Sleep(2*time.Second)
	//	ch<- 20
	//	close(ch)
	//}()
	//go func(){
	//	for y:=range ch {
	//		fmt.Println("main func")
	//		fmt.Println(y)
	//	}
	//}()
	//fmt.Println("end")
	//strings.Contains("ab","a")

	//redis测试代码
	//err := common.Rds.Set("mykey", "superWang",0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//username, err := common.Rds.Get("mykey").Result()
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Printf("Get mykey: %v \n", username)
	//}
	router.Router()

}