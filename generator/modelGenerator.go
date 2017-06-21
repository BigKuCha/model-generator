package generator

import (
	"fmt"
	"github.com/bigkucha/model-generator/database"
	"github.com/bigkucha/model-generator/helper"
	"github.com/dave/jennifer/jen"
	"github.com/jinzhu/inflection"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func Generate(c *cli.Context) error {
	dbSns := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		c.String("u"), c.String("p"), c.String("d"))
	db := database.GetDB(dbSns)
	if c.String("t") == "ALL" {
		tables := db.GetDataBySql("show tables")
		for _, table := range tables {
			tableName := table["Tables_in_"+c.String("d")]
			columns := db.GetDataBySql("desc " + tableName)
			generateModel(tableName, columns, c.String("dir"))
		}
	} else {
		columns := db.GetDataBySql("desc " + c.String("t"))
		generateModel(c.String("t"), columns, c.String("dir"))
	}
	return nil
}

func generateModel(tableName string, columns []map[string]string, dir string) {
	var codes []jen.Code
	for _, col := range columns {
		t := col["Type"]
		column := col["Field"]
		var st *jen.Statement
		if column == "id" {
			st = jen.Id("ID").Uint()
		} else {
			st = jen.Id(helper.SnakeCase2CamelCase(column, true))
			getCol(st, t)
		}
		codes = append(codes, st)
	}
	f := jen.NewFilePath(dir)
	f.Type().Id(helper.SnakeCase2CamelCase(inflection.Singular(tableName), true)).Struct(codes...)
	os.MkdirAll(dir, os.ModePerm)
	fileName := dir + "/" + helper.SnakeCase2CamelCase(inflection.Singular(tableName), false) + ".go"
	f.Save(fileName)
	fmt.Println(fileName)
}

func getCol(st *jen.Statement, t string) {
	prefix := strings.Split(t, "(")[0]
	switch prefix {
	case "int", "tinyint", "smallint", "bigint", "mediumint":
		st.Int()
	case "float":
		st.Float32()
	case "varchar":
		st.String()
	case "decimal":
		st.Float32()
	case "date", "time", "timestamp", "year", "datetime":
		st.Qual("time", "Time")
	default:
		st.String()
	}
}
