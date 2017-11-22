package service

import (
	"fmt"
	"encoding/json"
	"godemo/common"
	"reflect"
	"strings"
	"strconv"
)

/**
 * 查询直接返回json结构数组
 * sql=执行sql
 * args=sql参数
 */
func QuerybySql(sql string, args ...interface{}) []interface{} {
	var retmaps []interface{}
	rows, err := common.Db.Query(sql, args ...)
	if (err != nil) {
		fmt.Println(err)
	}
	//延时关闭Rows
	defer rows.Close()
	var values []interface{}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	values = make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}
	for rows.Next() {
		rows.Scan(values...)
		v := make(map[string]interface{})
		for i := range values {
			v[columns[i]] = fmt.Sprintf("%s", *(values[i].(*interface{})))
		}
		retmaps = append(retmaps, v)
	}
	alldata, err := json.Marshal(retmaps)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(alldata))
	return retmaps
}

/**
 * 查询返回结构体数组
 * dict= 传入实体类数组地址
 * sql=执行sql
 * args=sql参数
 */
func Query(dict interface{}, sql string, args ...interface{}) error {
	structArray := reflect.ValueOf(dict).Elem() //传入值为结构化数组指针，需获取他的值&[]user

	rows, err := common.Db.Query(sql, args ...)
	if (err != nil) {
		return err
	}
	//延时关闭Rows
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	//构建数据库返回数组，接收每行结果
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}
	for rows.Next() {
		rows.Scan(values...)
		var dictStruct = reflect.New(structArray.Type().Elem()) //UserRole
		//进行数据库返回结果到struct的映射转换,目前统一影射为string,待增加类型映射
		for i := range values {
			structName := strFirstToUpper(columns[i]) //Username
			sv := dictStruct.Elem().FieldByName(structName) //userrole.Username
			rv := *(values[i].(*interface{}))
			FilterDbValue(sv, rv)
		}
		structArray.Set(reflect.Append(structArray, dictStruct.Elem()))
	}
	return nil
}

func Saves() {

}
func update(){

}
func insert(dest interface{}){
	var id string
	id="1 or 1=1"
	common.Db.MustExec("select * from user where id=?",id)
}
/**
 * 字符串首字母转化为大写 ios_bbbbbbbb -> IosBbbbbbbbb
 */
func strFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

/*
数据库类型映射
 */
func FilterDbValue(sv reflect.Value, rv interface{}) error {
	valueType := sv.Type()
	fmt.Println(valueType.Kind())
	switch valueType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rv == nil {
			sv.SetInt(0)
		} else {
			value := reflect.ValueOf(rv)
			str := byteString(value.Bytes())
			ivalue, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			sv.SetInt(int64(ivalue))
		}
	case reflect.Float32, reflect.Float64:
		if rv == nil {
			sv.SetFloat(0)
		} else {
			value := reflect.ValueOf(rv)
			str := byteString(value.Bytes())
			fvalue, err := strconv.ParseFloat(str, 64)
			if err != nil {
				panic(err)
			}
			sv.SetFloat(fvalue)
		}
	case reflect.String:
		if rv == nil {
			sv.SetString("")
		} else {
			value := reflect.ValueOf(rv)
			sv.SetString(string(value.Bytes()))
		}
	case reflect.Bool:
		if rv == nil {
			sv.SetBool(false)
		} else {
			value := reflect.ValueOf(rv)
			b := value.Bytes()[0]
			if b == 0 || b == 48 {
				sv.SetBool(false)
			} else {
				sv.SetBool(true)
			}
		}
	case reflect.Struct:
		if rv!=nil{
			value := reflect.ValueOf(rv)
			sv.Set(value)
		}
	}
	return nil
}
/*
字节转字符串，防止0问题
 */
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
