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
		{
			Name: "pointer",
			Input: &Person{
				"chris",
				Profile{Age: 33, City: "london"},
			},
			ExpectedCalls: []string{"chris", "london"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{Age: 33, City: "london"},
				{Age: 34, City: "reykjavik"},
			},
			ExpectedCalls: []string{"london", "reykjavik"},
		},
		{
			Name: "array",
			Input: [2]Profile{
				{Age: 33, City: "london"},
				{Age: 34, City: "reykjavik"},
			},
			ExpectedCalls: []string{"london", "reykjavik"},
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

	// mapは順序保証しないので、上のテスト方法ではgotの中身が変わる
	t.Run("with map", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"baz": "boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "bar")
		assertContains(t, got, "boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "berlin"}
			aChannel <- Profile{34, "katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"berlin", "katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{
					33, "berlin",
				},
				Profile{
					34, "katowice",
				}
		}

		var got []string
		want := []string{"berlin", "katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
