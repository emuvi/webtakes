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

func testCase1() *TestCase {
	args := []string{"webtakes", "-i", "www.test.com", "-o", "output/dir"}
	criteria, _, _ := lib.Startup()
	criteria.Input = "www.test.com"
	criteria.Output = "output/dir"
	return &TestCase{args, criteria}
}

func TestParse(t *testing.T) {
	testCases := []TestCase{
		*testCase1(),
	}
	for _, testCase := range testCases {
		criteriaGot := Parse(testCase.args)
		if reflect.DeepEqual(criteriaGot, testCase.criteria) == false {
			t.Errorf("Parse of (%v) got %v and expected %v", testCase.args, criteriaGot, testCase.criteria)
		}
	}
}
