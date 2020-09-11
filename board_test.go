package checkers

import (
	"fmt"
	"testing"
)

func Test_board_get(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		args args
		want Piece
	}{
		{args{1,0}, BLUE},
		{args{7,2}, BLUE},
		{args{0,7}, RED},
		{args{0,5}, RED},
		{args{0,3}, EMPTY},
		{args{0,4}, EMPTY},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d should be %v", tt.args.i, tt.args.j, tt.want), func(t *testing.T) {
			b := NewBoard()
			if got := b.Get(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkersBoard_remove(t *testing.T) {
	type args struct {
		i,j int
		piece Piece
	}
	tests := []struct {
		args args
	}{
		{args{1, 0, RED}},
		{args{4, 3, BLUE}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d should be removed", tt.args.i, tt.args.j), func(t *testing.T) {
			b := NewBoard()
			if b.Remove(tt.args.i, tt.args.j); b.Get(tt.args.i, tt.args.j) != EMPTY {
				t.Errorf("%d,%d should be EMPTY", tt.args.i, tt.args.j)
			}

			if b.Add(tt.args.i, tt.args.j, tt.args.piece); b.Get(tt.args.i, tt.args.j) != tt.args.piece {
				t.Errorf("%d,%d should be %v", tt.args.i, tt.args.j, tt.args.piece)
			}
		})
	}
}