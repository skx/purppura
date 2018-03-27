package main

import (
	"testing"
	"time"
)

func TestClear(t *testing.T) {

	out, err := Str2Unix("clear")

	if err != nil {
		t.Errorf("Error parsing 'clear'")
	}
	if out != 0 {
		t.Errorf("Expected '%d' received '%d'", 0, out)
	}
}

func TestRaise(t *testing.T) {

	out, err := Str2Unix("now")
	now := time.Now().Unix()

	if err != nil {
		t.Errorf("Error parsing 'now'")
	}

	if now-out > 1 {
		t.Errorf("'now' didn't result in a current timestamp")
	}
}

func TestError(t *testing.T) {

	//
	// All these will fail
	//
	fails := []string{"moi.kiss",
		"3",
		"m",
		"h",
		"s",
		"33.33",
		"3.5m"}

	for _, str := range fails {

		_, err := Str2Unix(str)
		if err == nil {
			t.Errorf("Expected error parsing '%s' - and didn't get one", str)
		}
	}
}

// Test that the output is a correct
func TestValid(t *testing.T) {
	type TestCase struct {
		Input  string
		Offset int64
	}

	tests := []TestCase{
		{"5m", 300},
		{"5s", 5},
		{"+5m", 300},
		{"+5s", 5},
		{"+5H", 5 * 60 * 60},
		{"+4H", 4 * 60 * 60},
		{"+10m", 600},
		{"+20M", 1200},
		{"+10s", 10},
		{"+12S", 12},
		{"+10h", (60 * 60 * 10)},
	}

	for _, test := range tests {

		out, err := Str2Unix(test.Input)

		if err != nil {
			t.Errorf("Error parsing '%s'", test.Input)
		}

		now := time.Now().Unix()
		diff := out - now

		if diff != test.Offset {
			t.Errorf("'%s' should have differed by %d seconds - got %d instead", test.Input, test.Offset, diff)
		}
	}

}
