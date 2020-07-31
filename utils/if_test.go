package utils

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args struct {
		ok         bool
		trueValue  interface{}
		falseValue interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "int", args: args{ok: 1 < 2, trueValue: 1, falseValue: 2}, want: 1},
		{name: "string", args: args{ok: 1 < 2, trueValue: "2", falseValue: "2"}, want: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.ok, tt.args.trueValue, tt.args.falseValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
