package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"regexp"
	"strings"
	"text/template"
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

const t_short = `insert into {{.TableName}}` + "\n" +
	`values ({{$n := len .Values}}{{range $i, $v:=.Values}}
		{{- $v}}{{if ne (last $i) $n -}},{{- end}}{{end}});`

var fns = template.FuncMap{
	"last": func(x int) int {
		return x + 1
	},
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
	var stmt *Statement
	switch qt {
	case SHORT_INSERT:
		fmt.Println("Valid short insert statement!")
		stmt = sliceShortInsertStmt(sql)
		printStatment(stmt)
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

func sliceShortInsertStmt(sql string) *Statement {

	reTable := regexp.MustCompile("(?i)^insert\\s*into\\s*(\\w+)\\s*values\\s*\\(.+\\);?$")
	arr := reTable.FindStringSubmatch(sql)
	tableName := arr[1]

	re := regexp.MustCompile(`\(.*\)`)
	valueStr := strings.TrimFunc(re.FindString(sql), trimBrackets)
	values := strings.Split(valueStr, ",")
	p := &Statement{queryType: SHORT_INSERT, TableName: tableName, columns: nil, Values: values}
	return p
}

func printStatment(stmt *Statement) {

	t := template.Must(template.New("insertShort").Funcs(fns).Parse(t_short))
	err := t.Execute(os.Stdout, *stmt)
	if err != nil {
		fmt.Println("Error while executing template:", err)
	}
}

func trimBrackets(r rune) bool {
	return r == ')' || r == '('
}

type Statement struct {
	queryType QueryType
	TableName string
	columns   []string
	Values    []string
}
