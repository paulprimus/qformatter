package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"regexp"
	"strings"
)

type QueryType int

const (
	SHORT_INSERT QueryType = iota
	LONG_INSERT
	INVALID_QUERY
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
	qt := checkValidInsertStmt(sql)

	switch qt {
	case SHORT_INSERT:
		fmt.Println("Valid short insert statement!")
		sliceShortInsertStmt(sql)
	case LONG_INSERT:
		fmt.Println("Valid long insert statement!")
	case INVALID_QUERY:
		fmt.Println("Not a valid insert statement!")
	}
}

func checkValidInsertStmt(sql string) QueryType {
	var patternShort = regexp.MustCompile(`(?i)^insert\s*into\s*\w+\s*values\s*\(.+\);?$`)
	var patternLong = regexp.MustCompile(`(?i)^insert\s*into\s*\w+\s*\(.+\)\s*values\s*\(.*\);?$`)
	if patternShort.MatchString(sql) == true {
		return SHORT_INSERT
	} else if patternLong.MatchString(sql) == true {
		return LONG_INSERT
	} else {
		return INVALID_QUERY
	}
}

func sliceShortInsertStmt(sql string) {
	insertStmt := &Statement{queryType: QueryType.SHORT_INSERT, tableName: "testTable", columns: []string{"spalte1,", "spalte2"}, []string{"wert1", "wert2"}}
	fmt.Println(insertStmt)
}

type Statement struct {
	queryType QueryType
	tableName string
	columns   []string
	values    []string
}
