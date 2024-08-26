package main

import (
	"cron-parser/parser"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cron-parser \"*/15 0 1,15 * 1-5 /usr/bin/find\"")
		return
	}

	cronExpr := os.Args[1]
	fields := strings.Fields(cronExpr)

	if len(fields) != 6 {
		fmt.Println("Invalid cron expression format.")
		return
	}

	cronParser, err := parser.Validate(fields)
	if err != nil {
		fmt.Println(err)
		return
	}

	cronParser.Print()

}
