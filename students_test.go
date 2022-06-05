package coverage

import (
	"os"
	"testing"
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
