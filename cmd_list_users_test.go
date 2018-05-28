package main

import (
	"strings"
	"testing"
)

func TestListUsersName(t *testing.T) {
	s := listUsersCmd{}
	if s.Name() != "list-users" {
		t.Errorf("Field didn't match")
	}
}

func TestListUsersSynopsis(t *testing.T) {
	s := listUsersCmd{}
	if s.Synopsis() != "List our existing users." {
		t.Errorf("Field didn't match")
	}
}

func TestListUsersUsage(t *testing.T) {
	s := listUsersCmd{}
	if !strings.Contains(s.Usage(), "Show all the existing") {
		t.Errorf("Field didn't match")
	}
}
