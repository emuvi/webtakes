package lib

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const (
	GetText  = "GetText"
	GetClass = "GetClass"
	GetAttr  = "GetAttr"
	GetTag   = "GetTag"
)

type Save struct {
	GetWhat string
	Prepend string
	Append  string
}

const (
	CheckTag   = "CheckTag"
	CheckAttr  = "CheckAttr"
	CheckClass = "CheckClass"
	CheckText  = "CheckText"
)

const (
	AtActual   = "AtActual"
	AtParent   = "AtParent"
	AtAncestry = "AtAncestry"
	AtBrothers = "AtBrothers"
)

const (
	SeemsLikeAs  = "SeemsLikeAs"
	SeemsNotLike = "SeemsNotLike"
)

const (
	ThatIsEquals = "ThatIsEquals"
	ThatContains = "ThatContains"
	ThatStarts   = "ThatStarts"
	ThatEnds     = "ThatEnds"
	ThatPatterns = "ThatPatterns"
)

const (
	TieAnd = "TieAnd"
	TieOr  = "TieOr"
	TieNew = "TieNew"
)

type Clause struct {
	Saves []Save
	Check string
	HasAt string
	Seems string
	Thats string
	Which string
	TieBy string
}

type Criteria struct {
	Input   string
	Clauses []Clause
	Output  string
}

func Startup() (*Criteria, *Clause, *Save) {
	criteria := Criteria{}
	clause := criteria.NewClause()
	save := clause.NewSave()
	return &criteria, clause, save
}

func (criteria *Criteria) NewClause() *Clause {
	newClause := Clause{
		Saves: []Save{},
		Check: CheckTag,
		HasAt: AtActual,
		Seems: SeemsLikeAs,
		Thats: ThatIsEquals,
		Which: "",
		TieBy: TieAnd,
	}
	criteria.Clauses = append(criteria.Clauses, newClause)
	return &newClause
}

func (c *Clause) NewSave() *Save {
	newSave := Save{
		GetWhat: GetText,
	}
	c.Saves = append(c.Saves, newSave)
	return &newSave
}

func Take(criteria *Criteria) {
	resp, err := http.Get(criteria.Input)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(criteria.Output)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	GetContents(resp.Body, file)
	PutReferences(file, criteria.Input)
}

func GetContents(fromBody io.ReadCloser, toFile *os.File) {
	defer fromBody.Close()
	tokens := html.NewTokenizer(fromBody)
	for {
		kind := tokens.Next()
		if kind == html.ErrorToken {
			return
		}
		token := tokens.Token()
		switch {
		case kind == html.StartTagToken:
			toFile.WriteString(token.Data)
		case kind == html.TextToken:
			toFile.WriteString(" ")
			toFile.WriteString(strings.TrimSpace(token.Data))
		case kind == html.EndTagToken:
			toFile.WriteString("\n")
		}
	}
}

func PutReferences(file *os.File, input string) {
	file.WriteString("\n")
	file.WriteString("\nWebTakes Reference")
	file.WriteString("\n")
	file.WriteString("\n- From: <")
	file.WriteString(input)
	file.WriteString(">")
	file.WriteString("\n- When: ")
	file.WriteString(time.Now().UTC().Format(time.RFC3339))
	file.WriteString("\n")
}
