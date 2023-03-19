package eval

import "testing"

func TestEval1(t *testing.T) {
	value, err := Eval("1")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !value.IsNumber() || !value.IsInteger() {
		t.Errorf("expected type Integer, found %s", value.ValueType.String())
	}
	if value.IntegerValue != 1 {
		t.Errorf("expected value 1, found %d", value.IntegerValue)
	}
}

func TestEval101_2(t *testing.T) {
	value, err := Eval("101.2")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !value.IsNumber() || !value.IsFloat() {
		t.Errorf("expected type Float, found %s", value.ValueType.String())
	}
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 101.2 {
		t.Errorf("expected value 101.2, found %f", value.FloatValue)
	}
}

func TestEvalHelloWorld(t *testing.T) {
	value, err := Eval(`"hello world"`)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !value.IsString() {
		t.Errorf("expected type String, found %s", value.ValueType.String())
	}
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 0 {
		t.Errorf("expected value 0, found %f", value.FloatValue)
	}
	if value.StringValue != "hello world" {
		t.Errorf("expected value 'hello world', found '%s'", value.StringValue)
	}
}

func TestEvalEmptyString(t *testing.T) {
	value, err := Eval(`""`)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !value.IsString() {
		t.Errorf("expected type String, found %s", value.ValueType.String())
	}
	if value.IntegerValue != 0 {
		t.Errorf("expected value 0, found %d", value.IntegerValue)
	}
	if value.FloatValue != 0 {
		t.Errorf("expected value 0, found %f", value.FloatValue)
	}
	if value.StringValue != "" {
		t.Errorf("expected value empty string, found '%s'", value.StringValue)
	}
}
