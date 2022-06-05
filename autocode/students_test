package coverage

import (
	"os"
	"testing"
	"time"
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
	testTable := []struct {
		people People
	}{
		{
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
		{
			people: []Person{},
		},
	}

	//iterate
	for _, testCase := range testTable {
		num := testCase.people.Len()
		if num != len(testCase.people) {
			t.Error("Lens are not equal")
		}
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

	testCases := map[string]struct {
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
		result := val.beforeFunc.Less(val.i, val.j)
		if result != val.expected {
			t.Errorf("Less func problems %v", key)
		}
	}
}
