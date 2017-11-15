package cache

import (
	"godemo/common"
	"regexp"
	"encoding/json"
	"crypto/md5"
	"fmt"
	"database/sql"
)

const FIND_TABLE_NAME_REGSTR = "(?i)\\s+(from|join)\\s+`*(\\w+)`*\\s+"
const PREFIX_TABLE_NAME = "table_"
func CacheSelect(dest interface{}, query string, args ...interface{}){
	var firstTableName string
	reg := regexp.MustCompile(FIND_TABLE_NAME_REGSTR)
	tableNames := reg.FindAllStringSubmatch(query,-1)
	sqlvarsstr:= fmt.Sprint(args)
	haskey := md5.Sum([]byte(query+ sqlvarsstr))
	haskeystr := fmt.Sprintf("%x", haskey) //将[]byte转成16进制
	if len(tableNames)>0 {
		firstTableName =PREFIX_TABLE_NAME + tableNames[0][2]
		fmt.Println(firstTableName)
		if (common.Rds.HExists(firstTableName, haskeystr).Val()) {
			//get values from redis
			redisValue, _ := common.Rds.HGet(firstTableName, haskeystr).Bytes()
			json.Unmarshal(redisValue, dest)
		}else {
			common.Db.Select(dest, query, args ...)
			//set redis value
			jsonValue, _ := json.Marshal(dest)
			common.Rds.HSet(firstTableName, haskeystr, jsonValue)
		}
	}
}

func CacheMustExec(query string, args ...interface{}) sql.Result{
	reg := regexp.MustCompile(FIND_TABLE_NAME_REGSTR)
	tableNames := reg.FindAllStringSubmatch(query,-1)
	for _,name:=range tableNames {
		delTable := PREFIX_TABLE_NAME + name[2]
		common.Rds.Del(delTable)
	}
	return common.Db.MustExec(query,args ...)
}

func get(dest interface{}, query string, args ...interface{}) error {
	var firstTableName string
	var err error
	reg := regexp.MustCompile(FIND_TABLE_NAME_REGSTR)
	tableNames := reg.FindAllStringSubmatch(query,-1)
	sqlvarsstr:= fmt.Sprint(args)
	haskey := md5.Sum([]byte(query+ sqlvarsstr))
	haskeystr := fmt.Sprintf("%x", haskey) //将[]byte转成16进制
	if len(tableNames)>0 {
		firstTableName =PREFIX_TABLE_NAME + tableNames[0][2]
		if (common.Rds.HExists(firstTableName, haskeystr).Val()) {
			//get values from redis
			redisValue, _ := common.Rds.HGet(firstTableName, haskeystr).Bytes()
			json.Unmarshal(redisValue, dest)
		}else {
			err= common.Db.Get(dest,query,args ...)
			//set redis value
			jsonValue, _ := json.Marshal(dest)
			common.Rds.HSet(firstTableName, haskeystr, jsonValue)
		}
	}
	return err
}