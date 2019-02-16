package markdown

import (
	"regexp"
	"strings"
)

type Header struct {
	Path      string
	Name      string
	Poss      Positions
	RowString string
	Value     string
}

type Position struct {
	Line int
	Row  int
}

type Positions struct {
	St Position
	En Position
}

func ParseHeader(text string) []Header {
	results := []Header{}

	r := regexp.MustCompile(`#{1,6}.*`)
	if !r.MatchString(text) {
		return results
	}

	lines := strings.Split(text, "\n")
	strs := r.FindAllString(text, -1)
	indices := r.FindAllStringIndex(text, -1)

	for i := 0; i < len(strs); i++ {
		str := strs[i]
		st := indices[i][0]
		en := indices[i][1]

		poss := walkInText(lines, st, en-1)
		header := Header{
			Path:      "",
			Name:      "",
			Poss:      poss,
			RowString: str,
			Value:     "",
		}
		results = append(results, header)
	}
	return results
}

func getPositionsInText(text string, indices [][]int) []Positions {
	results := []Positions{}
	for _, index := range indices {
		st := index[0]
		en := index[1]
		lines := strings.Split(text, "\n")
		poss := walkInText(lines, st, en-1)
		results = append(results, poss)
	}
	return results
}

func walkInText(lines []string, st, en int) Positions {
	res := Positions{}

	walkIndex := 0

	for i, line := range lines {
		lineIndex := i + 1
		for rowIndex := 1; rowIndex < len(line)+1; rowIndex++ {
			if st == walkIndex {
				res.St.Row = rowIndex
				res.St.Line = lineIndex
			}
			if en == walkIndex {
				res.En.Row = rowIndex
				res.En.Line = lineIndex
			}
			walkIndex++
		}
		// count line break
		walkIndex++
	}
	return res
}
