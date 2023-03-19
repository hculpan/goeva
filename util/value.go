package util

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/hculpan/goeva/lexer"
)

type ValueTypeIdentifier int64

const (
	Null ValueTypeIdentifier = iota
	Integer
	Float
	String
	Boolean
)

//go:generate stringer -type ValueTypeIdentifier

type Value struct {
	ValueType ValueTypeIdentifier

	IntegerValue int64
	FloatValue   float64
	StringValue  string
	BooleanValue bool
}

func NewValueFromToken(t lexer.Token) *Value {
	if t.Type == lexer.FLOAT {
		n, err := strconv.ParseFloat(t.Literal, 64)
		if err != nil {
			panic(err)
		}
		return NewFloatValue(n)
	} else if t.Type == lexer.INTEGER {
		n, err := strconv.ParseInt(t.Literal, 10, 64)
		if err != nil {
			panic(err)
		}
		return NewIntegerValue(n)
	} else if t.Type == lexer.STRING {
		return NewStringValue(t.Literal)
	} else {
		panic("unimplemented value from token " + t.Type.String())
	}
}

func NewIntegerValue(v int64) *Value {
	return &Value{IntegerValue: v, ValueType: Integer}
}

func NewFloatValue(v float64) *Value {
	return &Value{FloatValue: v, ValueType: Float}
}

func NewStringValue(v string) *Value {
	if len(v) > 0 && v[0] == '"' && v[len(v)-1] == '"' {
		return &Value{StringValue: v[1 : len(v)-1], ValueType: String}
	}
	return &Value{StringValue: v, ValueType: String}
}

func NewBooleanValue(v bool) *Value {
	return &Value{BooleanValue: v, ValueType: Boolean}
}

func (v *Value) String() string {
	switch {
	case v.IsFloat():
		return strconv.FormatFloat(v.FloatValue, 'f', -1, 64)
	case v.IsInteger():
		return strconv.FormatInt(v.IntegerValue, 10)
	case v.IsString():
		return v.StringValue
	case v.IsBoolean():
		return strconv.FormatBool(v.BooleanValue)
	default:
		return "unknown"
	}
}

func (v *Value) SetInteger(val int64) error {
	if v.IsInteger() {
		v.IntegerValue = val
	} else if v.IsFloat() {
		v.FloatValue = float64(val)
	} else {
		return errors.New("invalid data type: cannot set integer on string type")
	}

	return nil
}

func (v *Value) SetFloat(val float64) error {
	if v.IsInteger() {
		v.ValueType = Float
		v.IntegerValue = 0
		v.FloatValue = val
	} else if v.IsFloat() {
		v.FloatValue = val
	} else {
		return errors.New("invalid data type: cannot set float on string type")
	}

	return nil
}

func (v *Value) SetString(val string) error {
	if !v.IsString() {
		return errors.New("invalid data type: cannot set string on non-string type")
	} else {
		v.StringValue = val
	}

	return nil
}

func (v *Value) SetBoolean(val bool) error {
	if !v.IsBoolean() {
		return errors.New("invalid data type: cannot set boolean on non-boolean type")
	} else {
		v.BooleanValue = val
	}

	return nil
}

func (v *Value) IsNumber() bool {
	return (v.IsFloat() || v.IsInteger())
}

func (v *Value) IsInteger() bool {
	return (v.ValueType == Integer)
}

func (v *Value) IsFloat() bool {
	return (v.ValueType == Float)
}

func (v *Value) IsString() bool {
	return (v.ValueType == String)
}

func (v *Value) IsBoolean() bool {
	return (v.ValueType == Boolean)
}

func (v *Value) AddInteger(val int64) {
	if v.IsInteger() {
		v.IntegerValue += val
	} else if v.IsFloat() {
		v.FloatValue += float64(val)
	} else {
		panic(fmt.Sprintf("cannot add int to type %s", v.ValueType.String()))
	}
}

func (v *Value) AddFloat(val float64) {
	if v.IsInteger() {
		v.FloatValue = float64(v.IntegerValue) + val
		v.IntegerValue = 0
		v.ValueType = Float
	} else if v.IsFloat() {
		v.FloatValue += val
	} else {
		panic(fmt.Sprintf("cannot add float to type %s", v.ValueType.String()))
	}
}

func (v *Value) AddString(val string) {
	if v.IsString() {
		v.StringValue += val
	} else {
		panic(fmt.Sprintf("cannot add string to type %s", v.ValueType.String()))
	}
}

func (v *Value) SubInteger(val int64) {
	if v.IsInteger() {
		v.IntegerValue -= val
	} else if v.IsFloat() {
		v.FloatValue -= float64(val)
	} else {
		panic(fmt.Sprintf("cannot sub int from type %s", v.ValueType.String()))
	}
}

func (v *Value) SubFloat(val float64) {
	if v.IsInteger() {
		v.FloatValue = float64(v.IntegerValue) - val
		v.IntegerValue = 0
		v.ValueType = Float
	} else if v.IsFloat() {
		v.FloatValue -= val
	} else {
		panic(fmt.Sprintf("cannot sub float from type %s", v.ValueType.String()))
	}
}

func (v *Value) MultInteger(val int64) {
	if v.IsInteger() {
		v.IntegerValue *= val
	} else if v.IsFloat() {
		v.FloatValue *= float64(val)
	} else {
		panic(fmt.Sprintf("cannot multiply int to type %s", v.ValueType.String()))
	}
}

func (v *Value) MultFloat(val float64) {
	if v.IsInteger() {
		v.FloatValue = float64(v.IntegerValue) * val
		v.IntegerValue = 0
		v.ValueType = Float
	} else if v.IsFloat() {
		v.FloatValue *= val
	} else {
		panic(fmt.Sprintf("cannot multiply float from type %s", v.ValueType.String()))
	}
}

func (v *Value) DivInteger(val int64) {
	if v.IsInteger() {
		n := float64(v.IntegerValue) / float64(val)
		if n == math.Floor(n) {
			v.IntegerValue = int64(math.Floor(n))
		} else {
			v.FloatValue = n
			v.IntegerValue = 0
			v.ValueType = Float
		}
	} else if v.IsFloat() {
		v.FloatValue /= float64(val)
	} else {
		panic(fmt.Sprintf("cannot divide int into type %s", v.ValueType.String()))
	}
}

func (v *Value) DivFloat(val float64) {
	if v.IsInteger() {
		v.FloatValue = float64(v.IntegerValue) / val
		v.IntegerValue = 0
		v.ValueType = Float
	} else if v.IsFloat() {
		v.FloatValue /= val
	} else {
		panic(fmt.Sprintf("cannot divide float into type %s", v.ValueType.String()))
	}
}

func (v1 *Value) Add(v2 *Value) {
	switch v2.ValueType {
	case Integer:
		v1.AddInteger(v2.IntegerValue)
	case Float:
		v1.AddFloat(v2.FloatValue)
	case String:
		v1.AddString(v2.StringValue)
	}
}

func (v1 *Value) Sub(v2 *Value) {
	switch v2.ValueType {
	case Integer:
		v1.SubInteger(v2.IntegerValue)
	case Float:
		v1.SubFloat(v2.FloatValue)
	}
}

func (v1 *Value) Mult(v2 *Value) {
	switch v2.ValueType {
	case Integer:
		v1.MultInteger(v2.IntegerValue)
	case Float:
		v1.MultFloat(v2.FloatValue)
	}
}

func (v1 *Value) Div(v2 *Value) {
	switch v2.ValueType {
	case Integer:
		v1.DivInteger(v2.IntegerValue)
	case Float:
		v1.DivFloat(v2.FloatValue)
	}
}
