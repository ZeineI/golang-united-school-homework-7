package coverage

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
This file consists of two parts, separated from each other by "/".
The first part consists of a structure with implemented methods Len, Less and Swap,
which means the implementation of the sort.Interface interface.
The second part consists of implementing a matrix to store numbers in rows and columns.
Your goal to write tests for this code and achieve as much coverage as possible.
*/

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

//People is a slice of type person
type People []Person

//Len returns the length of input value
func (p People) Len() int {
	return len(p)
}

//Less reports when it is necessary to sort element i before element j
func (p People) Less(i, j int) bool {
	if p[i].birthDay.Unix() == p[j].birthDay.Unix() {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		}
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthDay.Unix() > p[j].birthDay.Unix()
}

//Swap swaps the elements with indexes i and j.
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

///////////////////////////////////////////////////////////////////////////////////////////////////////

//Matrix stores an array of numbers in rows and columns
type Matrix struct {
	rows, cols int
	data       []int
}

/*New creates a matrix from a string.*/
func New(str string) (*Matrix, error) {
	rows := strings.Split(str, "\n")
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

/*Rows gets the matrix represented in rows.*/
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

/*Cols gets the matrix represented in columns*/
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

/*Set sets the value of the matrix at point row, col.*/
func (m *Matrix) Set(row, col, value int) bool {
	if row < 0 || m.rows <= row || col < 0 || m.cols <= col {
		return false
	}
	m.data[row*m.cols+col] = value
	return true
}

