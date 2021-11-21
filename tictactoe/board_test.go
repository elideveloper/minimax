package tictactoe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWin(t *testing.T) {

	testCases := []struct {
		name     string
		board    [5][5]int
		plID     int
		needResp bool
	}{
		{
			board: [5][5]int{
				{1, 0, 0, 1, 0},
				{0, 0, 0, 1, 0},
				{1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1},
			},
			plID:     1,
			needResp: false,
		},
		{
			board: [5][5]int{
				{0, 0, 0, 1, 1},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 0},
				{1, 0, 0, 1, 1},
			},
			plID:     1,
			needResp: false,
		},
		{
			board: [5][5]int{
				{0, 1, 0, 1, 1},
				{0, 1, 0, 1, 1},
				{0, 1, 0, 0, 1},
				{0, 0, 0, 1, 0},
				{1, 0, 0, 1, 1},
			},
			plID:     1,
			needResp: false,
		},
		{
			name: "vertical 4",
			board: [5][5]int{
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
				{1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0},
				{1, 0, 0, 0, 1},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "vertical 5",
			board: [5][5]int{
				{1, 0, 0, 1, 1},
				{0, 1, 0, 1, 1},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 1, 0},
				{1, 0, 0, 1, 1},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "horizontal 4",
			board: [5][5]int{
				{0, 0, 0, 1, 0},
				{0, 1, 1, 1, 1},
				{1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "diag left 4",
			board: [5][5]int{
				{0, 0, 0, 1, 0},
				{0, 1, 0, 1, 1},
				{1, 0, 1, 1, 1},
				{0, 1, 0, 0, 0},
				{1, 0, 0, 1, 1},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "diag right 4",
			board: [5][5]int{
				{1, 0, 0, 1, 1},
				{0, 1, 0, 0, 1},
				{0, 0, 1, 1, 0},
				{0, 0, 0, 1, 0},
				{1, 0, 0, 1, 0},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "diag right up 4",
			board: [5][5]int{
				{0, 1, 0, 1, 1},
				{0, 1, 1, 0, 1},
				{0, 0, 1, 1, 0},
				{0, 0, 0, 1, 1},
				{1, 0, 0, 1, 0},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "diag right down 4",
			board: [5][5]int{
				{0, 0, 0, 1, 1},
				{1, 1, 1, 0, 1},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 1},
				{1, 0, 0, 1, 0},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "diag left down 4",
			board: [5][5]int{
				{0, 0, 0, 1, 1},
				{1, 0, 1, 0, 1},
				{0, 1, 0, 1, 0},
				{0, 0, 1, 0, 1},
				{0, 1, 0, 1, 0},
			},
			plID:     1,
			needResp: true,
		},
		{
			name: "diag left up 4",
			board: [5][5]int{
				{0, 0, 0, 1, 1},
				{0, 1, 1, 0, 1},
				{0, 1, 0, 0, 0},
				{1, 0, 1, 0, 1},
				{1, 0, 0, 0, 0},
			},
			plID:     1,
			needResp: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBoard(5, 4)
			b.matrix = tc.board
			resp := b.IsWin(tc.plID)
			assert.Equal(t, tc.needResp, resp)
		})
	}
}
