package main

import (
	"os"
	"strings"
	"webtakes/lib"
)

func Parse(args []string) *lib.Criteria {
	index := 1
	length := len(args)
	criteria, clause, save := lib.Startup()
	for index < length {
		thisArg := args[index]
		nextArg := ""
		if index+1 < length {
			nextArg = args[index+1]
		}
		if thisArg == "-i" || thisArg == "--input" {
			criteria.Input = nextArg
			index++
		} else if thisArg == "-o" || thisArg == "--output" {
			criteria.Output = nextArg
			index++
		} else if strings.HasPrefix(thisArg, "Get") {
			save.GetWhat = thisArg
		} else if thisArg == "Prepend" {
			save.Prepend = nextArg
			index++
		} else if thisArg == "Append" {
			save.Append = nextArg
			index++
		} else if thisArg == "SaveToo" {
			save = clause.NewSave()
		} else if strings.HasPrefix(thisArg, "Check") {
			clause.Check = thisArg
		} else if strings.HasPrefix(thisArg, "At") {
			clause.HasAt = thisArg
		} else if strings.HasPrefix(thisArg, "Seems") {
			clause.Seems = thisArg
		} else if strings.HasPrefix(thisArg, "That") {
			clause.Thats = thisArg
		} else if thisArg == "Which" {
			clause.Which = nextArg
			index++
		} else if strings.HasPrefix(thisArg, "Tie") {
			if thisArg == "TieNew" {
				clause = criteria.NewClause()
			} else {
				clause.TieBy = thisArg
			}
		}
		index++
	}
	return criteria
}

func main() {
	lib.Take(Parse(os.Args))
}
