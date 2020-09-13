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
		{args{0, 0}, RedPawn},
		{args{6, 2}, RedPawn},
		{args{0, 6}, BluePawn},
		{args{1, 7}, BluePawn},
		{args{0, 4}, Empty},
		{args{1, 3}, Empty},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d should be %v", tt.args.i, tt.args.j, tt.want), func(t *testing.T) {
			b := NewBoard()
			if got := b.Get(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Get() = %v, isLegal %v", got, tt.want)
			}
		})
	}
}

func Test_checkersBoard_remove_and_add(t *testing.T) {
	type args struct {
		i, j  int
		piece Piece
	}
	tests := []struct {
		args args
	}{
		{args{0, 0, RedPawn}},
		{args{2, 6, BluePawn}},
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

func Test_illegal_moves(t *testing.T) {
	type args struct {
		i, j int
	}
	tests := []struct {
		args args
	}{
		{args{0, 1}},
		{args{1, 0}},
		{args{1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d should be illegal", tt.args.i, tt.args.j), func(t *testing.T) {
			b := NewBoard()
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()
			b.Get(tt.args.i, tt.args.j)
		})
	}
}
