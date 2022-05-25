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
		lib.CheckText, lib.AtParent, lib.SeemsNotLike, lib.ThatStarts,
		"Which", "by which", lib.TieOr,
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

func TestParse(t *testing.T) {
	testCases := []TestCase{
		*testCaseWithPaths(),
		*testCaseWithSave(),
		*testCaseAllBasics(),
		*testCaseMultipleSaves(),
	}
	for _, testCase := range testCases {
		criteriaGot := Parse(testCase.args)
		if reflect.DeepEqual(criteriaGot, testCase.criteria) == false {
			t.Errorf("Parse of (%v) got %v and expected %v", testCase.args, criteriaGot, testCase.criteria)
		}
	}
}
