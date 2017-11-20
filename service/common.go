package service

import (
	"fmt"
	"encoding/json"
	"godemo/common"
	"reflect"
	"strings"
)

func QuerybySql(sql string,args ...interface{}) []interface{}{
	var retmaps []interface{}
	rows,err:=common.Db.Query(sql,args ...)
	if(err!=nil){
		fmt.Println(err)
	}
	var values []interface{}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	values = make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}
	for rows.Next(){
		rows.Scan(values...)
		v := make(map[string]interface{})
		for i:= range values{
			v[columns[i]] =fmt.Sprintf("%s",*(values[i].(*interface{})))
		}
		retmaps = append(retmaps, v)
	}
	alldata ,err:= json.Marshal(retmaps)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(alldata))
	return retmaps
}
/**
 * 查询返回结构体数组
 */
func Query(dict interface{},sql string,args...interface{}) error{
	structArray := reflect.ValueOf(dict).Elem() //传入值为结构化数组指针，需获取他的值&[]user

	rows,err:=common.Db.Query(sql,args ...)
	if(err!=nil){
		return err
	}
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	//构建数据库返回数组，接收每行结果
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}
	for rows.Next(){
		rows.Scan(values...)
		var dictStruct = reflect.New(structArray.Type().Elem())
		//进行数据库返回结果到struct的映射转换,目前统一影射为string,待增加类型映射
		for i:= range values{
			structName := strFirstToUpper(columns[i])
			var structValue string
			if rv :=*(values[i].(*interface{}));rv!=nil {
				structValue = fmt.Sprintf("%s",rv)
			}else {
				structValue = ""
			}
			dictStruct.Elem().FieldByName(structName).SetString(structValue)
		}
		structArray.Set(reflect.Append(structArray,dictStruct.Elem()))
	}
	return nil
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
