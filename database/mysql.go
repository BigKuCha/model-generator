package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type mdb struct {
	db *sql.DB
}

func GetDB(dbSns string) *mdb {
	db, err := sql.Open("mysql", dbSns)
	if err != nil {
		fmt.Println("连接数据库失败", err.Error())
		os.Exit(1)
	}
	return &mdb{db: db}
}

func (d *mdb) GetDataBySql(sql string) []map[string]string {
	//查询数据库
	query, err := d.db.Query(sql)
	if err != nil {
		fmt.Println("查询数据库失败", err.Error())
		os.Exit(1)
	}
	defer query.Close()

	//读出查询出的列字段名
	cols, _ := query.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}

	//
	var results []map[string]string
	//results := make(map[string]string)
	i := 0
	for query.Next() { //循环，让游标往下推
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			os.Exit(1)
		}

		row := make(map[string]string) //每行数据

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row) //装入结果集中
		i++
	}

	//查询出来的数组
	//d.db.Close() //用完关闭
	return results
}
