package tag

import (
	"reflect"
	"sort"
	"testing"
)

type obj struct {
	Id int `odd:"id"`
}

func isSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Strings(aCopy)
	sort.Strings(bCopy)
	return reflect.DeepEqual(aCopy, bCopy)
}

func TestGetTaggedFields(t *testing.T) {
	OddTag := "odd"
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"No Tag",
			args{
				struct {
					Name string
				}{
					"hi",
				},
			},
			[]string{},
		}, {
			"flat",
			args{
				obj{11},
			},
			[]string{"Id"},
		}, {
			"nested",
			args{
				struct {
					Obj obj
				}{
					obj{12},
				},
			},
			[]string{"Id"},
		}, {
			"tagged struct",
			args{
				struct {
					Obj obj `odd:"obj"`
				}{
					obj{12},
				},
			},
			[]string{"Id", "Obj"},
		}, {
			"omit empty",
			args{
				struct {
					Test int `odd:"obj"`
				}{},
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTaggedFields(tt.args.s, OddTag)
			names := make([]string, len(got))
			for i, v := range got {
				names[i] = v.Name
			}
			if !isSliceEqual(names, tt.want) {
				t.Errorf("GetTaggedFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
