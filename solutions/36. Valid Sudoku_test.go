package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSudoku(t *testing.T) {
	fnNames := map[string]func(board [][]byte) bool{
		"isValidSudoku": isValidSudoku,
	}

	type TestCase struct {
		Name     string
		Board    [][]byte
		Expected bool
	}

	tcs := []TestCase{
		{
			Name: "1_leetCode_Valid",
			Board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			Expected: true,
		},
		{
			Name: "2_leetCode_DuplicateInColumn_I1_J1_E8",
			Board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			Expected: false,
		},
		{
			Name: "3_DuplicateInSquare_I2_J3_E3",
			Board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '3', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			Expected: false,
		},
		{
			Name: "4_leetCode_DuplicateInColumn_I2_J2_E6",
			Board: [][]byte{
				{'.', '.', '5', '.', '.', '.', '.', '.', '.'},
				{'1', '.', '.', '2', '.', '.', '.', '.', '.'},
				{'.', '.', '6', '.', '.', '3', '.', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'3', '.', '1', '5', '2', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '4', '.'},
				{'.', '.', '6', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '9', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			Expected: false,
		},
		{
			Name: "4_leetCode_DuplicateInRow_I2_J2_E6",
			Board: [][]byte{
				{'.', '.', '5', '.', '.', '.', '.', '.', '.'},
				{'1', '.', '.', '2', '.', '.', '.', '.', '.'},
				{'.', '.', '6', '.', '.', '3', '6', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'3', '.', '1', '5', '2', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '4', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '9', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			Expected: false,
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := fn(tc.Board)

				assert.Equal(t, tc.Expected, out)
			})
		}
	}
}