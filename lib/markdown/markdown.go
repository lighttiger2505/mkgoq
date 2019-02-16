package markdown

import (
	"regexp"
	"strings"
)

type Header struct {
	Path      string
	Name      string
	StPos     Position
	EnPos     Position
	RowString string
	Value     string
}

type Position struct {
	Line int
	Row  int
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

		stPos, enPos := walkInText(lines, st, en-1)
		header := Header{
			Path:      "",
			Name:      "",
			StPos:     stPos,
			EnPos:     enPos,
			RowString: str,
			Value:     "",
		}
		results = append(results, header)
	}
	return results
}

func walkInText(lines []string, st, en int) (stPos Position, enPos Position) {
	walkIndex := 0
	for i, line := range lines {
		lineIndex := i + 1
		for rowIndex := 1; rowIndex < len(line)+1; rowIndex++ {
			if st == walkIndex {
				stPos = Position{
					Row:  rowIndex,
					Line: lineIndex,
				}
			}
			if en == walkIndex {
				enPos = Position{
					Row:  rowIndex,
					Line: lineIndex,
				}
			}
			walkIndex++
		}
		// count line break
		walkIndex++
	}
	return
}

