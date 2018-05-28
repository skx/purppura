package main

import (
	"strings"
	"testing"
)

func TestDElUserName(t *testing.T) {
	s := delUserCmd{}
	if s.Name() != "del-user" {
		t.Errorf("Field didn't match")
	}
}

func TestDelUserSynopsis(t *testing.T) {
	s := delUserCmd{}
	if s.Synopsis() != "Delete an existing user." {
		t.Errorf("Field didn't match")
	}
}

func TestDelUserUsage(t *testing.T) {
	s := delUserCmd{}
	if !strings.Contains(s.Usage(), "Remove a user") {
		t.Errorf("Field didn't match")
	}
}
