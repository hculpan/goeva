package util

import (
	"testing"
)

func TestNewFloat(t *testing.T) {
	value := NewFloatValue(2.0)
	if value.FloatValue != 2.0 {
		t.Errorf("expected value 2.0, found %f", value.FloatValue)
	}
	if value.ValueType != Float {
		t.Errorf("expected type Float, found %d", value.ValueType)
	}
}

func TestNewInteger(t *testing.T) {
	value := NewIntegerValue(2)
	if value.IntegerValue != 2 {
		t.Errorf("expected value 2, found %d", value.IntegerValue)
	}
	if value.ValueType != Integer {
		t.Errorf("expected type Integer, found %d", value.ValueType)
	}
}

func TestNewString(t *testing.T) {
	value := NewStringValue("hello")
	if value.StringValue != "hello" {
		t.Errorf("expected value 'hello', found '%s'", value.StringValue)
	}
	if value.ValueType != String {
		t.Errorf("expected type String, found %d", value.ValueType)
	}
}

func TestSetIntegerOnInteger(t *testing.T) {
	value := NewIntegerValue(2)
	value.SetInteger(3)
	if value.IntegerValue != 3 {
		t.Errorf("expected value 3, found %d", value.IntegerValue)
	}
	if value.ValueType != Integer {
		t.Errorf("expected type Integer, found %d", value.ValueType)
	}
}

func TestSetIntegerOnFloat(t *testing.T) {
	value := NewFloatValue(2)
	value.SetInteger(3)
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 3.0 {
		t.Errorf("expected value 3.0, found %f", value.FloatValue)
	}
	if value.ValueType != Float {
		t.Errorf("expected type Float, found %d", value.ValueType)
	}
}

func TestSetIntegerOnString(t *testing.T) {
	value := NewStringValue("hello")
	value.SetInteger(3)
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.StringValue != "hello" {
		t.Errorf("expected value 'hello', found %s", value.StringValue)
	}
	if !value.IsString() {
		t.Errorf("expected type String, found %d", value.ValueType)
	}
}

func TestSetFloatOnInteger(t *testing.T) {
	value := NewIntegerValue(2)
	value.SetFloat(3.0)
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 3.0 {
		t.Errorf("expected value 3.0, found %f", value.FloatValue)
	}
	if !value.IsNumber() || !value.IsFloat() {
		t.Errorf("expected type Float, found %d", value.ValueType)
	}
}

func TestSetFloatOnFloat(t *testing.T) {
	value := NewFloatValue(2)
	value.SetFloat(3)
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 3.0 {
		t.Errorf("expected value 3.0, found %f", value.FloatValue)
	}
	if value.ValueType != Float {
		t.Errorf("expected type Float, found %d", value.ValueType)
	}
}

func TestSetFloatOnString(t *testing.T) {
	value := NewStringValue("hello")
	value.SetFloat(3)
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 0.0 {
		t.Errorf("expected value 0.0, found %f", value.FloatValue)
	}
	if value.StringValue != "hello" {
		t.Errorf("expected value 'hello', found %s", value.StringValue)
	}
	if !value.IsString() {
		t.Errorf("expected type String, found %d", value.ValueType)
	}
}

func TestSetStringOnInteger(t *testing.T) {
	value := NewIntegerValue(2)
	value.SetString("hello")
	if value.IntegerValue != 2 {
		t.Errorf("expected value 2, found %d", value.IntegerValue)
	}
	if value.FloatValue != 0.0 {
		t.Errorf("expected value 0.0, found %f", value.FloatValue)
	}
	if !value.IsNumber() || !value.IsInteger() {
		t.Errorf("expected type Integer, found %d", value.ValueType)
	}
}

func TestSetStringOnFloat(t *testing.T) {
	value := NewFloatValue(2.0)
	value.SetString("hello")
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 2.0 {
		t.Errorf("expected value 2.0, found %f", value.FloatValue)
	}
	if !value.IsNumber() || !value.IsFloat() {
		t.Errorf("expected type Float, found %d", value.ValueType)
	}
}

func TestSetStringOnString(t *testing.T) {
	value := NewStringValue("hello")
	value.SetString("hello world")
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 0.0 {
		t.Errorf("expected value 0.0, found %f", value.FloatValue)
	}
	if value.StringValue != "hello world" {
		t.Errorf("expected value 'hello world', found %s", value.StringValue)
	}
	if !value.IsString() {
		t.Errorf("expected type String, found %d", value.ValueType)
	}
}

func TestAddIntegerInteger(t *testing.T) {
	value := NewIntegerValue(1)
	value.AddInteger(10)
	if value.IntegerValue != 11 {
		t.Errorf("expected value 11, found %d", value.IntegerValue)
	}
	if !value.IsInteger() {
		t.Errorf("expected type Integer, found %s", value.ValueType.String())
	}
}

func TestDivIntegerInteger(t *testing.T) {
	value := NewIntegerValue(12)
	value.DivInteger(3)
	if value.IntegerValue != 4 {
		t.Errorf("expected value 4, found %d", value.IntegerValue)
	}
	if !value.IsInteger() {
		t.Errorf("expected type Integer, found %s", value.ValueType.String())
	}
}
