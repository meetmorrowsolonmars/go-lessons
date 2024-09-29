package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/meetmorrowsolonmars/go-lessons/testing/coverage/internal/domain"
)

func TestIsEvenNumber(t *testing.T) {
	type args struct {
		number int
	}

	tests := []struct {
		name string
		argv args
		want bool
	}{
		{
			name: "even number",
			argv: args{
				number: 10,
			},
			want: true,
		},
		{
			name: "odd number",
			argv: args{
				number: 11,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := domain.IsEvenNumber(tt.argv.number)

			assert.Equal(t, tt.want, result)
		})
	}
}
