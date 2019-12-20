package main

import (
	"github.com/BigKuCha/model-generator/generator"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "generate model for jinzhu/gorm"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username,u",
			Value: "root",
			Usage: "Username of mysql",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "",
			Usage: "Password of mysql",
		},
		cli.StringFlag{
			Name:  "database, d",
			Value: "",
			Usage: "select database",
		},
		cli.StringFlag{
			Name:  "table, t",
			Usage: "table name",
			Value: "ALL",
		},
		cli.StringFlag{
			Name:  "dir",
			Usage: "path which models will be stored",
			Value: "models",
		},
	}
	app.Action = generator.Generate
	app.Run(os.Args)
}
