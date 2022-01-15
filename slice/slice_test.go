package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		src interface{}
		m   func(interface{}) interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantDst []interface{}
	}{
		{
			name: "int slice",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				m: func(i interface{}) interface{} {
					return fmt.Sprintf("%d", i)
				},
			},
			wantDst: []interface{}{"1", "2", "3", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDst := Map(tt.args.src, tt.args.m); !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("Map() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	type args struct {
		src interface{}
		fn  func(interface{}) bool
	}
	tests := []struct {
		name    string
		args    args
		wantDst []interface{}
	}{
		{
			name: "int slice",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				fn: func(i interface{}) bool {
					return i.(int)%2 == 0
				},
			},
			wantDst: []interface{}{int(2), int(4)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDst := Select(tt.args.src, tt.args.fn); !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("Select() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}
