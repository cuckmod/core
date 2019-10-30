package core

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
)

type (
	Enum   int
	Enumer interface {
		Name() string
		Ordinal() int
		Values() *[]string

		// Text
		MarshalText() ([]byte, error)
		UnmarshalText([]byte) error

		// JSON
		MarshalJSON() ([]byte, error)
		UnmarshalJSON(b []byte) error

		// SQL
		Scan(interface{}) error
		Value() (driver.Value, error)
	}
)

func ValidateEnum(e Enumer) (err error) {
	if e.Ordinal() < 0 || e.Ordinal() >= len(*e.Values()) {
		return errors.New("unknown enum")
	}
	return
}

func MarshalText(e Enumer) (text []byte, err error) {
	err = ValidateEnum(e)
	if err != nil {
		return
	}
	return []byte(e.Name()), nil
}

func UnmarshalText(e interface{}, text []byte) (err error) {
	for k, v := range *((e).(Enumer)).Values() {
		if v == string(text) {
			reflect.ValueOf(e).Elem().SetInt(int64(k))
			return nil
		}
	}
	return errors.New("unknown enum")
}

func MarshalJSON(e Enumer) ([]byte, error) {
	text, err := MarshalText(e)
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func UnmarshalJSON(e interface{}, b []byte) error {
	var un string
	if err := json.Unmarshal(b, &un); err != nil {
		return err
	}
	return UnmarshalText(e, []byte(un))
}

func Scan(e interface{}, src interface{}) (err error) {
	switch t := src.(type) {
	case int, int8, int16, int32, int64:
		i := reflect.ValueOf(t).Convert(reflect.ValueOf(int64(0)).Type()).Int()
		reflect.ValueOf(e).Elem().SetInt(i)
	case string:
		if err := UnmarshalText(e, []byte(t)); err != nil {
			return err
		}
	case []byte:
		if err := UnmarshalText(e, t); err != nil {
			return err
		}
	default:
		return errors.New("unknown type for enum: " + reflect.ValueOf(t).Type().String())
	}
	return
}

func Value(e Enumer) (driver.Value, error) {
	text, err := MarshalText(e)
	if err != nil {
		return nil, err
	}
	return text, nil
}
