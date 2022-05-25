package main

import (
	"reflect"
	"testing"
	"webtakes/lib"
)

type TestCase struct {
	args     []string
	criteria *lib.Criteria
}

func testCaseWithPaths() *TestCase {
	args := []string{"webtakes", "-i", "www.test.com", "-o", "output/dir"}
	criteria, _, _, _ := lib.Startup()
	criteria.OfInput = "www.test.com"
	criteria.ToOutput = "output/dir"
	return &TestCase{args, criteria}
}

func testCaseWithSave() *TestCase {
	args := []string{"webtakes", "-i", "www.test.com", "-o", "output/dir",
		lib.GetClass, "Prepend", "prepend with", "Append", "append with"}
	criteria, _, _, save := lib.Startup()
	criteria.OfInput = "www.test.com"
	criteria.ToOutput = "output/dir"
	save.GetWhat = lib.GetClass
	save.Prepend = "prepend with"
	save.Append = "append with"
	return &TestCase{args, criteria}
}

func testCaseAllBasics() *TestCase {
	args := []string{"webtakes", "-i", "www.test.com", "-o", "output/dir",
		lib.GetClass, "Prepend", "prepend with", "Append", "append with",
		lib.CheckText, lib.AtParent, lib.SeemsNotLike, lib.ThatStarts, "Which", "by which", lib.TieOr,
	}
	criteria, _, clause, save := lib.Startup()
	criteria.OfInput = "www.test.com"
	criteria.ToOutput = "output/dir"
	save.GetWhat = lib.GetClass
	save.Prepend = "prepend with"
	save.Append = "append with"
	clause.Check = lib.CheckText
	clause.HasAt = lib.AtParent
	clause.Seems = lib.SeemsNotLike
	clause.Thats = lib.ThatStarts
	clause.Which = "by which"
	clause.TieBy = lib.TieOr
	return &TestCase{args, criteria}
}

func testCaseMultipleSaves() *TestCase {
	args := []string{"webtakes", "-i", "www.test.com", "-o", "output/dir",
		lib.GetClass, "Prepend", "prepend with", "Append", "append with",
		lib.SaveNew, lib.GetAttr, "Prepend", "attr pre with", "Append", "attr pos with"}
	criteria, extract, _, save := lib.Startup()
	criteria.OfInput = "www.test.com"
	criteria.ToOutput = "output/dir"
	save.GetWhat = lib.GetClass
	save.Prepend = "prepend with"
	save.Append = "append with"
	save = extract.NewSave()
	save.GetWhat = lib.GetAttr
	save.Prepend = "attr pre with"
	save.Append = "attr pos with"
	return &TestCase{args, criteria}
}

func testCaseMultipleSavesAndClauses() *TestCase {
	args := []string{"webtakes", "-i", "www.test.com", "-o", "output/dir",
		lib.GetClass, "Prepend", "prepend with", "Append", "append with",
		lib.SaveNew, lib.GetAttr, "Prepend", "attr pre with", "Append", "attr pos with",
		lib.CheckText, lib.AtParent, lib.SeemsNotLike, lib.ThatStarts, "Which", "by which", lib.TieOr,
		lib.CheckTag, lib.AtBrothers, lib.SeemsLikeAs, lib.ThatContains, "Which", "contains this",
	}
	criteria, extract, clause, save := lib.Startup()
	criteria.OfInput = "www.test.com"
	criteria.ToOutput = "output/dir"
	save.GetWhat = lib.GetClass
	save.Prepend = "prepend with"
	save.Append = "append with"
	save = extract.NewSave()
	save.GetWhat = lib.GetAttr
	save.Prepend = "attr pre with"
	save.Append = "attr pos with"
	clause.Check = lib.CheckText
	clause.HasAt = lib.AtParent
	clause.Seems = lib.SeemsNotLike
	clause.Thats = lib.ThatStarts
	clause.Which = "by which"
	clause.TieBy = lib.TieOr
	clause = extract.NewClause()
	clause.Check = lib.CheckTag
	clause.HasAt = lib.AtBrothers
	clause.Seems = lib.SeemsLikeAs
	clause.Thats = lib.ThatContains
	clause.Which = "contains this"
	return &TestCase{args, criteria}
}

func testCaseMultipleExtracts() *TestCase {
	args := []string{"webtakes", "-i", "www.test.com", "-o", "output/dir",
		lib.GetClass, "Prepend", "prepend with", "Append", "append with",
		lib.SaveNew, lib.GetAttr, "Prepend", "attr pre with", "Append", "attr pos with",
		lib.CheckText, lib.AtParent, lib.SeemsNotLike, lib.ThatStarts, "Which", "by which", lib.TieOr,
		lib.CheckTag, lib.AtBrothers, lib.SeemsLikeAs, lib.ThatContains, "Which", "contains this", lib.TieNew,
		lib.GetText, "Prepend", "text after", "Append", "text before",
	}
	criteria, extract, clause, save := lib.Startup()
	criteria.OfInput = "www.test.com"
	criteria.ToOutput = "output/dir"
	save.GetWhat = lib.GetClass
	save.Prepend = "prepend with"
	save.Append = "append with"
	save = extract.NewSave()
	save.GetWhat = lib.GetAttr
	save.Prepend = "attr pre with"
	save.Append = "attr pos with"
	clause.Check = lib.CheckText
	clause.HasAt = lib.AtParent
	clause.Seems = lib.SeemsNotLike
	clause.Thats = lib.ThatStarts
	clause.Which = "by which"
	clause.TieBy = lib.TieOr
	clause = extract.NewClause()
	clause.Check = lib.CheckTag
	clause.HasAt = lib.AtBrothers
	clause.Seems = lib.SeemsLikeAs
	clause.Thats = lib.ThatContains
	clause.Which = "contains this"
	_, _, save = criteria.NewExtract()
	save.GetWhat = lib.GetText
	save.Prepend = "text after"
	save.Append = "text before"
	return &TestCase{args, criteria}
}

func TestParse(t *testing.T) {
	testCases := []TestCase{
		*testCaseWithPaths(),
		*testCaseWithSave(),
		*testCaseAllBasics(),
		*testCaseMultipleSaves(),
		*testCaseMultipleSavesAndClauses(),
		*testCaseMultipleExtracts(),
	}
	for _, testCase := range testCases {
		criteriaExpect := testCase.criteria
		criteriaResult := Parse(testCase.args)
		if reflect.DeepEqual(criteriaExpect, criteriaResult) == true {
			t.Errorf("Parse of (%v) result was %v and was expected %v", testCase.args, criteriaResult, criteriaExpect)
		}
	}
}
