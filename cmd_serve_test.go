package main

import (
	"strings"
	"testing"
)

func TestServeName(t *testing.T) {
	s := serveCmd{}
	if s.Name() != "serve" {
		t.Errorf("Field didn't match")
	}
}

func TestServeSynopsis(t *testing.T) {
	s := serveCmd{}
	if s.Synopsis() != "Launch the HTTP server." {
		t.Errorf("Field didn't match")
	}
}

func TestServeUsage(t *testing.T) {
	s := serveCmd{}
	if !strings.Contains(s.Usage(), "Launch the HTTP server for") {
		t.Errorf("Field didn't match")
	}
}
