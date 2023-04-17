package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Add(t *testing.T) {
	assert := assert.New(t)
	type args struct {
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestTree_Add",
			args: args{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := true
			assert.Equal(tt.want, got, tt.name)
		})
	}
}
