package calc

import (
	"bingo/pkg/logger"
	"testing"
)

func TestCreateUserCard(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{max: 0}, 0},
		{"25", args{max: 25}, 25},
		{"minus", args{max: -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CreateUserCard(tt.args.max)
			logger.Debug(r)
			if tt.want != len(r) {
				t.Errorf("CreateUserCard() = %v, want %v", len(r), tt.args.max)
			}
		})
	}
}
