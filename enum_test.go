package core

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strings"
	"testing"
)

type testEnum Enum

const (
	foo testEnum = iota
	bar
	baz
)

var testName = []string{
	foo: "foo",
	bar: "bar",
	baz: "baz",
}

func (t testEnum) Name() string {
	if ValidateEnum(&t) != nil {
		return ""
	}
	return (*t.Values())[t]
}

func (t testEnum) Ordinal() int {
	return int(t)
}

func (t testEnum) Values() *[]string {
	return &testName
}

func (t testEnum) MarshalText() ([]byte, error) {
	return MarshalText(&t)
}

func (t *testEnum) UnmarshalText(b []byte) error {
	return UnmarshalText(t, b)
}

func (t testEnum) MarshalJSON() ([]byte, error) {
	return MarshalJSON(&t)
}

func (t *testEnum) UnmarshalJSON(b []byte) error {
	return UnmarshalJSON(t, b)
}

func (t *testEnum) Scan(src interface{}) error {
	return Scan(t, src)
}

func (t testEnum) Value() (driver.Value, error) {
	return Value(&t)
}

func TestMarshalText(t *testing.T) {
	test := bar
	text, err := test.MarshalText()
	if err != nil {
		t.Error(err)
	}
	if string(text) != "bar" {
		t.Log(string(text))
		t.Fail()
	}

	test = -1
	text, err = test.MarshalText()
	if err == nil || err.Error() != "unknown enum" {
		t.Fail()
	}
}

func TestUnmarshalText(t *testing.T) {
	var test testEnum
	err := test.UnmarshalText([]byte("bar"))
	if err != nil {
		t.Error(err)
	}
	if test != 1 {
		t.Log(test)
		t.Fail()
	}

	err = test.UnmarshalText([]byte("beep"))
	if err == nil || err.Error() != "unknown enum" {
		t.Fail()
	}
}

func TestMarshalJSON(t *testing.T) {
	test := bar
	b, err := test.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(b) != "\"bar\"" {
		t.Log(string(b))
		t.Fail()
	}

	b, err = json.Marshal(&test)
	if err != nil {
		t.Error(err)
	}
	if string(b) != "\"bar\"" {
		t.Log(string(b))
		t.Fail()
	}

	test = -1
	b, err = test.MarshalJSON()
	if err == nil || err.Error() != "unknown enum" {
		t.Log(err)
		t.Fail()
	}

	b, err = json.Marshal(&test)
	if err == nil || !strings.HasSuffix(err.Error(), "unknown enum") {
		t.Log(err)
		t.Fail()
	}
}

func TestUnmarshalJSON(t *testing.T) {
	var test testEnum
	err := test.UnmarshalJSON([]byte("\"bar\""))
	if err != nil {
		t.Error(err)
	}
	if test != 1 {
		t.Log(test)
		t.Fail()
	}

	err = json.Unmarshal([]byte("\"bar\""), &test)
	if err != nil {
		t.Error(err)
	}
	if test != 1 {
		t.Log(test)
		t.Fail()
	}

	err = test.UnmarshalJSON([]byte("\"beep\""))
	if err == nil || err.Error() != "unknown enum" {
		t.Log(err)
		t.Fail()
	}
	err = json.Unmarshal([]byte("\"beep\""), &test)
	if err == nil || !strings.HasSuffix(err.Error(), "unknown enum") {
		t.Log(err)
		t.Fail()
	}
}

func TestScan(t *testing.T) {
	var test testEnum
	err := test.Scan(int(1))
	if err != nil {
		t.Error(err)
	}
	if test != 1 {
		t.Log(test)
		t.Fail()
	}

	err = test.Scan("bar")
	if err != nil {
		t.Error(err)
	}
	if test != 1 {
		t.Log(test)
		t.Fail()
	}

	err = test.Scan([]byte("bar"))
	if err != nil {
		t.Error(err)
	}
	if test != 1 {
		t.Log(test)
		t.Fail()
	}

	err = test.Scan("beep")
	if err == nil || err.Error() != "unknown enum" {
		t.Log(err)
		t.Fail()
	}
	err = test.Scan([]byte("beep"))
	if err == nil || err.Error() != "unknown enum" {
		t.Log(err)
		t.Fail()
	}
}

func TestValue(t *testing.T) {
	test := bar
	value, err := test.Value()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(value.([]byte), []byte("bar")) {
		t.Log(value)
		t.Fail()
	}

	test = testEnum(-1)
	value, err = test.Value()
	if err == nil || err.Error() != "unknown enum" {
		t.Error(err)
	}
}
