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

const SaveNew = "SaveNew"

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
	Check string
	HasAt string
	Seems string
	Thats string
	Which string
	TieBy string
}

type Extract struct {
	Saves   []Save
	Clauses []Clause
}

type Criteria struct {
	OfInput  string
	Extracts []Extract
	ToOutput string
}

func Startup() (*Criteria, *Extract, *Clause, *Save) {
	criteria := Criteria{}
	extract, clause, save := criteria.NewExtract()
	return &criteria, extract, clause, save
}

func (c *Criteria) NewExtract() (*Extract, *Clause, *Save) {
	newExtract := Extract{}
	newClause := newExtract.NewClause()
	newSave := newExtract.NewSave()
	c.Extracts = append(c.Extracts, newExtract)
	return &newExtract, newClause, newSave
}

func (e *Extract) NewSave() *Save {
	newSave := Save{
		GetWhat: GetText,
	}
	e.Saves = append(e.Saves, newSave)
	return &newSave
}

func (e *Extract) NewClause() *Clause {
	newClause := Clause{
		Check: CheckTag,
		HasAt: AtActual,
		Seems: SeemsLikeAs,
		Thats: ThatIsEquals,
		Which: "",
		TieBy: TieAnd,
	}
	e.Clauses = append(e.Clauses, newClause)
	return &newClause
}

func Take(criteria *Criteria) {
	resp, err := http.Get(criteria.OfInput)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(criteria.ToOutput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	GetContents(resp.Body, file)
	PutReferences(file, criteria.OfInput)
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
