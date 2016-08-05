package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"pandora/modules/logger"
	"pandora/vars"
	"strconv"
	"strings"
)

var currDb *sql.DB= nil;

func getDb() *sql.DB {
	//logger.D(vars.Conf.GetString("mysqlURL"))
	if(currDb!=nil){
		return currDb;
	}
	db, err := sql.Open("mysql", vars.Conf.GetString("mysqlURL"))
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	checkErr(err)
	currDb=db
	return db
}

//插入demo
func Insert(sql string, args ...interface{}) int64 {
	db := getDb()
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(args...)
	if err != nil {
		logger.W("db", err)
		return -1
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.E("db", err)
		return -1
	}
	stmt.Close()
	return id
}

func Test() {
	logger.D("db", "testing")
}
func GetRealSql(sql string, args ...interface{}) string {
	str := sql
	for _, v := range args {
		vs := fmt.Sprint(v)
		vs = strings.Trim(vs, "\n")
		vs = strings.TrimSpace(vs)
		_, e := strconv.ParseInt(vs, 10, 64)
		if e != nil {
			vs = "'" + vs + "'"
		}
		str = strings.Replace(str, "?", vs, 1)
	}
	return str
}

//查询demo
func Query(sql string, args ...interface{}) *sql.Rows {
	db := getDb()
	//logger.D("db", "querying:", getRealSql(sql, args...))
	stmt, err := db.Prepare(sql)
	checkErr(err)
	rows, err := stmt.Query(args...)
	checkErr(err)
	//	//普通demo
	//	//for rows.Next() {
	//	//    var userId int
	//	//    var userName string
	//	//    var userAge int
	//	//    var userSex int
	//
	//	//    rows.Columns()
	//	//    err = rows.Scan(&userId, &userName, &userAge, &userSex)
	//	//    checkErr(err)
	//
	//	//    fmt.Println(userId)
	//	//    fmt.Println(userName)
	//	//    fmt.Println(userAge)
	//	//    fmt.Println(userSex)
	//	//}
	//
	//	//字典类型
	//	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	//	columns, _ := rows.Columns()
	//	scanArgs := make([]interface{}, len(columns))
	//	values := make([]interface{}, len(columns))
	//	for i := range values {
	//		scanArgs[i] = &values[i]
	//	}
	//
	//	for rows.Next() {
	//		//将行数据保存到record字典
	//		err = rows.Scan(scanArgs...)
	//		record := make(map[string]string)
	//		for i, col := range values {
	//			if col != nil {
	//				record[columns[i]] = string(col.([]byte))
	//			}
	//		}
	//		fmt.Println(record)
	//	}

	stmt.Close()
	return rows
}


//获取唯一结果，如count 等
func QueryUnigueResult(sql string, args ...interface{}) string {
	r := Query(sql, args...)
	var s string
	for r.Next() {
		r.Scan(&s)
	}
	return s
}
func QuerySqlToJSON(jsonKeys []string, sql string, args ...interface{}) string {
	r := Query(sql, args...)
	return convertRowsToJSON(r, jsonKeys)
}
func ConvertQueryToJSON(rows *sql.Rows, jsonKeys ...string) string {
	return convertRowsToJSON(rows, jsonKeys)
}
func convertRowsToJSON(rows *sql.Rows, jsonKeys []string) string {
	cols, _ := rows.Columns()
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	values := make([][]byte, len(cols))
	for i := range values {
		dest[i] = &values[i]
	}
	//var data [][]map[string]interface{}
	var results []map[string]interface{}
	for rows.Next() {
		rows.Scan(dest...)
		//logger.D("db",dest)
		m := make(map[string]interface{})
		for i, k := range jsonKeys {
			m[k] = string(values[i])

		}
		results = append(results, m)
	}
	if len(results) == 0 {
		return "[]"
	}
	s, _ := json.Marshal(results)
	rows.Close()
	return string(s)
}

//更新数据
func Update(sql string, args ...interface{}) int64 {
	db := getDb()
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(args...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	stmt.Close()
	return num
}

//删除数据
func Remove(sql string, args ...interface{}) int64 {
	db := getDb()

	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(args...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	stmt.Close()
	return num
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
}
