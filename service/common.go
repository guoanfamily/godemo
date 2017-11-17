package service

import (
	"fmt"
	"encoding/json"
	"godemo/common"
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
