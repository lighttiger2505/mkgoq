package markdown

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func getTestContent(t *testing.T, fpath string) string {
	file, err := os.OpenFile(fpath, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("cannot open file, %s", err.Error())
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("cannot read file, %s", err.Error())
	}
	return string(b)
}

func TestParseHeader(t *testing.T) {
	text := getTestContent(t, "./testdata/header_test.md")
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []Header
	}{
		{
			name: "normal",
			args: args{
				text: text,
			},
			want: []Header{
				Header{
					Path: "",
					Name: "",
					Poss: Positions{
						St: Position{
							Line: 1,
							Row:  1,
						},
						En: Position{
							Line: 1,
							Row:  21,
						},
					},
					RowString: "# The largest heading",
					Value:     "",
				},
				Header{
					Path: "",
					Name: "",
					Poss: Positions{
						St: Position{
							Line: 3,
							Row:  1,
						},
						En: Position{
							Line: 3,
							Row:  29,
						},
					},
					RowString: "## The second largest heading",
					Value:     "",
				},
				Header{
					Path: "",
					Name: "",
					Poss: Positions{
						St: Position{
							Line: 5,
							Row:  1,
						},
						En: Position{
							Line: 5,
							Row:  27,
						},
					},
					RowString: "###### The smallest heading",
					Value:     "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseHeader(tt.args.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPositionsInText(t *testing.T) {
	text := getTestContent(t, "./testdata/header_test.md")
	type args struct {
		text    string
		indices [][]int
	}
	tests := []struct {
		name string
		args args
		want []Positions
	}{
		{
			name: "normal",
			args: args{
				text: text,
				indices: [][]int{
					[]int{0, 21},
					[]int{42, 71},
					[]int{99, 126},
				},
			},
			want: []Positions{
				Positions{
					St: Position{
						Line: 1,
						Row:  1,
					},
					En: Position{
						Line: 1,
						Row:  21,
					},
				},
				Positions{
					St: Position{
						Line: 3,
						Row:  1,
					},
					En: Position{
						Line: 3,
						Row:  29,
					},
				},
				Positions{
					St: Position{
						Line: 5,
						Row:  1,
					},
					En: Position{
						Line: 5,
						Row:  27,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPositionsInText(tt.args.text, tt.args.indices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPositionsInText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walkInText(t *testing.T) {
	lines := []string{
		"# one",
		"## two",
		"###### five",
	}
	type args struct {
		lines []string
		st    int
		en    int
	}
	tests := []struct {
		name string
		args args
		want Positions
	}{
		{
			name: "normal",
			args: args{
				lines: lines,
				st:    6,
				en:    11,
			},
			want: Positions{
				St: Position{
					Line: 2,
					Row:  1,
				},
				En: Position{
					Line: 2,
					Row:  6,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := walkInText(tt.args.lines, tt.args.st, tt.args.en); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walkInText() = %v, want %v", got, tt.want)
			}
		})
	}
}
