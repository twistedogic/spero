package odd

import (
	"reflect"
	"testing"
)

func TestParseOutcome(t *testing.T) {
	type testStruct struct {
		H  string `json:"H" odd:"type=had"`
		D  string `json:"D" odd:"type=had"`
		A  string `json:"A" odd:"type=had"`
		ID string `json:"ID" odd:"id"`
	}
	cases := []struct {
		input  testStruct
		ID     string
		output []Outcome
	}{
		{
			testStruct{"100@1.1", "100@1.2", "100@1.3", "ID"},
			"ID",
			[]Outcome{
				{OddID: "ID", Type: HAD, Outcome: "H", Odd: 1.1},
				{OddID: "ID", Type: HAD, Outcome: "D", Odd: 1.2},
				{OddID: "ID", Type: HAD, Outcome: "A", Odd: 1.3},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.ID, func(t *testing.T) {
			output := ParseOutcome(tt.input)
			if !reflect.DeepEqual(tt.output, output) {
				t.Errorf("want %v, got %v", tt.output, output)
			}
		})
	}
}

func Test_parseOdd(t *testing.T) {
	type args struct {
		odd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"normal", args{"100@1.1"}, 1.1},
		{"malform but valid", args{"1.1"}, 1.1},
		{"malform", args{"word"}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOdd(tt.args.odd); got != tt.want {
				t.Errorf("parseOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
