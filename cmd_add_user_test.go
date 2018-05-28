package main

import (
	"strings"
	"testing"
)

func TestAddUserName(t *testing.T) {
	s := addUserCmd{}
	if s.Name() != "add-user" {
		t.Errorf("Field didn't match")
	}
}

func TestAddUserSynopsis(t *testing.T) {
	s := addUserCmd{}
	if s.Synopsis() != "Add a new user." {
		t.Errorf("Field didn't match")
	}
}

func TestAddUserUsage(t *testing.T) {
	s := addUserCmd{}
	if !strings.Contains(s.Usage(), "Add a new") {
		t.Errorf("Field didn't match")
	}
}
