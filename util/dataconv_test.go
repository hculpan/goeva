package util

import "testing"

func TestIsInteger(t *testing.T) {
	if !IsInteger("1") {
		t.Errorf("expected true for integer check of '1', got false")
	}
	if !IsInteger("1901231") {
		t.Errorf("expected true for integer check of '1901231', got false")
	}
	if !IsInteger("-1901231") {
		t.Errorf("expected true for integer check of '-1901231', got false")
	}
}

func TestIsFloat(t *testing.T) {
	if !IsFloat("1.0") {
		t.Errorf("expected true for integer check of '1.0', got false")
	}
	if !IsFloat("1901231.013") {
		t.Errorf("expected true for integer check of '1901231.013', got false")
	}
	if !IsFloat("-1901231.87662") {
		t.Errorf("expected true for integer check of '-1901231.87662', got false")
	}
}
