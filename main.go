package main

import (
	//"errors"
	"fmt"
	"github.com/urfave/cli"
	//"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "qformattter"

	app.Commands = []cli.Command{
		{
			Name:      "fmt",
			Usage:     "formatiere ein SQL",
			ArgsUsage: "[sql]",
			Action:    formatiere,
		},
	}

	app.Usage = "SQLs leserlich formtieren!"
	app.Run(os.Args)
}

func formatiere(c *cli.Context) error {
	if len(c.Args()) == 0 {
		return cli.NewExitError("Fehler: SQL als Parameter angeben!\nUsage: qformatter(.exe) fmt [sql]", 1)
	}
	rawSql := c.Args().First()
	rawSql = strings.Trim(rawSql, " ")
	if strings.HasPrefix(rawSql, "insert") {
		formatiereInsert(rawSql)	
	}
	return nil
}

func formatiereInsert(sql s) {
	rawSqlArr := strings.Split(sql, ",")
	for i := 0; i < len(rawSqlArr); i++ {
		a := strings.Trim(rawSqlArr[i], " ")
		fmt.Printf("%d: %s\n", i, a)
		if strings.Compare(a, "(") == 0 {
			fmt.Println("Found (")
		}
	}
}
