package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	rows, cols int
	data       []int
}

func New(in string) (*Matrix, error) {
	rows := strings.Split(in, "\n")
	matrix := Matrix{rows: len(rows), cols: -1}
	for _, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		if matrix.cols == -1 {
			matrix.cols = len(cols)
		} else if matrix.cols != len(cols) {
			return nil, fmt.Errorf("Rows need to be the same length")
		}
		for _, char := range cols {
			num, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			matrix.data = append(matrix.data, num)
		}
	}
	return &matrix, nil
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, m.rows)
	for r := 0; r < m.rows; r++ {
		rows[r] = make([]int, m.cols)
		for c := 0; c < m.cols; c++ {
			rows[r][c] = m.data[r*m.cols+c]
		}
	}
	return rows
}

func (m Matrix) Cols() [][]int {
	cols := make([][]int, m.cols)
	for c := 0; c < m.cols; c++ {
		cols[c] = make([]int, m.rows)
		for r := 0; r < m.rows; r++ {
			cols[c][r] = m.data[r*m.cols+c]
		}
	}
	return cols
}

func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || m.rows <= r || c < 0 || m.cols <= c {
		return false
	}
	m.data[r*m.cols+c] = val
	return true
}
