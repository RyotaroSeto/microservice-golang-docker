package pkg

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{a: 1, b: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Sum(tt.args.a, tt.args.b))
		})
	}
}
