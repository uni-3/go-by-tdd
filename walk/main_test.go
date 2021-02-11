package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

// Test .
func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with two string field",
			Input: struct {
				Name string
				City string
			}{"chris", "london"},
			ExpectedCalls: []string{"chris", "london"},
		},
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
				Age  int
			}{"chris", 33},
			ExpectedCalls: []string{"chris"},
		},
		{
			Name: "nested fields",
			Input: Person{
				"chris",
				Profile{Age: 33, City: "london"},
			},
			ExpectedCalls: []string{"chris", "london"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}

		})
	}

}
