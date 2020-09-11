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
		{args{1,0}, BluePawn},
		{args{7,2}, BluePawn},
		{args{0,7}, RedPawn},
		{args{0,5}, RedPawn},
		{args{0,3}, Empty},
		{args{0,4}, Empty},
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

func Test_checkersBoard_remove_and_add(t *testing.T) {
	type args struct {
		i,j int
		piece Piece
	}
	tests := []struct {
		args args
	}{
		{args{1, 0, RedPawn}},
		{args{4, 3, BluePawn}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d should be removed", tt.args.i, tt.args.j), func(t *testing.T) {
			b := NewBoard()
			if b.Remove(tt.args.i, tt.args.j); b.Get(tt.args.i, tt.args.j) != Empty {
				t.Errorf("%d,%d should be Empty", tt.args.i, tt.args.j)
			}

			if b.Add(tt.args.i, tt.args.j, tt.args.piece); b.Get(tt.args.i, tt.args.j) != tt.args.piece {
				t.Errorf("%d,%d should be %v", tt.args.i, tt.args.j, tt.args.piece)
			}
		})
	}
}