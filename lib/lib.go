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

type ToSave struct {
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
	TieNew = "TieNew"
)

type Clause struct {
	Saves []ToSave
	Check string
	HasAt string
	Seems string
	That  string
	With  string
	TieBy string
}

type Criteria = []Clause

func Take(input string, output string) {
	resp, err := http.Get(input)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	GetContents(resp.Body, file)
	PutReferences(file, input)
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
