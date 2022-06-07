package coverage

import (
	"errors"
	"fmt"
	"os"

	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	// "golang.org/x/tools/go/expect"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestLen(t *testing.T) {
	//create
	testTable := map[string]struct {
		people People
	}{
		"len two": {
			people: []Person{
				{
					firstName: "Zeine",
					lastName:  "Imanmalik",
				},
				{
					firstName: "Zeine2",
					lastName:  "Imanmalik2",
				},
			},
		},
		"empty": {
			people: []Person{},
		},
	}

	//iterate
	for name, testCase := range testTable {
		t.Run(name, func(t *testing.T) {
			num := testCase.people.Len()
			if !assert.Equal(t,num, len(testCase.people) ){
				t.Errorf("Len func problems %v", name)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	//create
	testCase := struct {
		people   People
		expected People
	}{
		people: []Person{
			{
				firstName: "Zeine",
				lastName:  "Imanmalik",
			},
			{
				firstName: "Zeine2",
				lastName:  "Imanmalik2",
			},
		},
		expected: []Person{
			{
				firstName: "Zeine2",
				lastName:  "Imanmalik2",
			},
			{
				firstName: "Zeine",
				lastName:  "Imanmalik",
			},
		},
	}

	//iterate
	testCase.people.Swap(0, 1)
	if testCase.people[0] != testCase.expected[0] || testCase.people[1] != testCase.expected[1] {
		t.Error("The swap operation isnt correct")
	}
}

func TestLess(t *testing.T) {
	// index i j
	//expected
	//the beggining form

	testCases := map[string]struct{
		beforeFunc People
		expected   bool
		i          int
		j          int
	}{
		"correct case": {
			beforeFunc: []Person{
				{
					firstName: "Zeine",
					lastName:  "Imanmalik",
					birthDay:  time.Date(2021, time.April, 11, 21, 34, 01, 0, time.UTC),
				},
				{
					firstName: "Zeine",
					lastName:  "Imanmalik",
					birthDay:  time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
				},
			},
			expected: true,
			i:        0,
			j:        1,
		},
		"second correct case": {
			beforeFunc: []Person{
				{
					firstName: "Another",
					lastName:  "Man",
					birthDay:  time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
				},
				{
					firstName: "Zeine",
					lastName:  "Imanmalik",
					birthDay:  time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
				},
			},
			expected: true,
			i:        0,
			j:        1,
		},
		"third correct case": {
			beforeFunc: []Person{
				{
					firstName: "Zeine",
					lastName:  "Man",
					birthDay:  time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
				},
				{
					firstName: "Zeine",
					lastName:  "Imanmalik",
					birthDay:  time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
				},
			},
			expected: false,
			i:        0,
			j:        1,
		},
	}

	for key, val := range testCases {
		t.Run(key, func(t *testing.T) {
			result := val.beforeFunc.Less(val.i, val.j)
			if !assert.Equal(t, result, val.expected) {
				t.Errorf("Less func problems %v", key)
			}
		})
		
	}
}

func TestNew(t *testing.T) {
	testCases := map[string]struct{
		matrix 	*Matrix
		str 	string
		expErr 	error
	}{
		"correct": {
			matrix: &Matrix{
				rows: 2,
				cols: 1,
				data: []int{1,2},
			},
			str: "1\n2",
			expErr: nil,
		},
		"notNums": {
			matrix: &Matrix{},
			str: "Hola",
			expErr: &strconv.NumError{Func: "Atoi", Num: "Hola", Err: strconv.ErrSyntax},
		},
		"rows not the same": {
			matrix: &Matrix{},
			str: "1 2\n2",
			expErr: errors.New("Rows need to be the same length"),
		},
	}

	for key, val := range testCases {
		t.Run(key, func(t *testing.T) {
			result, err := New(val.str)
			if err != nil {
				if err.Error() != val.expErr.Error() {
					t.Errorf("inccorect")
				}
			} else {
				if !assert.Equal(t, result, val.matrix) {
					t.Errorf("Func inccorect %v", key)
				}
			}
			
		})
		
	}
}

func TestRows(t *testing.T) {
	testCase := struct{
		matrix Matrix
		expResp [][]int
	}{
		matrix: Matrix{
			rows: 2,
			cols: 1,
			data: []int{1,2},
		},
		expResp: [][]int{{1}, {2}},
	}

	result := testCase.matrix.Rows()
	if !assert.Equal(t, result, testCase.expResp) {
		t.Errorf("Not right row func")
	}
}

func TestCols(t *testing.T) {
	testCase := struct{
		matrix Matrix
		expResp [][]int
	}{
		matrix: Matrix{
			rows: 2,
			cols: 1,
			data: []int{1,2},
		},
		expResp: [][]int{{1, 2}},
	}

	result := testCase.matrix.Cols()
	fmt.Println(result)
	if !assert.Equal(t, result, testCase.expResp) {
		t.Errorf("Not right cols func")
	}
}

func TestSet(t *testing.T) {
	testCases := map[string]struct{
		m 				Matrix
		row, col, val 	int
		expResp 		bool
	}{
		"-row": {
			m: Matrix{
				rows: 2,
				cols: 3,
				data: []int{1,2},
			},
			row: -5,
			col: 2,
			val: 5,
			expResp: false,
		},
		"m.row <= row": {
			m: Matrix{
				rows: 2,
				cols: 3,
				data: []int{1,2},
			},
			row: 10,
			col: 2,
			val: 5,
			expResp: false,
			
		},
		"-col": {
			m: Matrix{
				rows: 2,
				cols: 3,
				data: []int{1,2},
			},
			row: 1,
			col: -2,
			val: 5,
			expResp: false,

		},
		"m.col <= col": {
			m: Matrix{
				rows: 2,
				cols: 3,
				data: []int{1,2},
			},
			row: 2,
			col: 5,
			val: 5,
			expResp: false,
			
		},
		"correct": {
			m: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1,1,2,2},
			},
			row: 1,
			col: 1,
			val: 5,
			expResp: true,

		},
	}

	for key, val := range testCases {
		t.Run(key, func(t *testing.T) {
			result := val.m.Set(val.row, val.col,val.val)
			if result != val.expResp {
				t.Errorf("Incorrect")
			}
		})
		
	}

}