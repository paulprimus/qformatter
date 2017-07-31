package main

import (
	//"errors"
	"fmt"
	"github.com/urfave/cli"
	//"log"
	"os"
	"strings"
	"regexp"
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
	} else {
		formatiereInsert(rawSql)	
	}
	return nil
}

func formatiereInsert(sql string) {
	checkValidInsertStmt(sql)

	/*rawSqlArr := strings.Split(sql, ",")
	for i := 0; i < len(rawSqlArr); i++ {
		a := strings.Trim(rawSqlArr[i], " ")
		fmt.Printf("%d: %s\n", i, a)
		if strings.Compare(a, "(") == 0 {
			fmt.Println("Found (")			
		}
	}*/
}

func checkValidInsertStmt(sql string) {
	var pattern = regexp.MustCompile(`^insert|INSERT\sinto|INTO\s\w+.*`)
	if pattern.MatchString(sql) == true {
		fmt.Println("Valid insert statement!")
	} else {
		fmt.Println("Not a valid insert statement!")
	}
	fmt.Println(sql)
}