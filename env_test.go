package core

import (
	"os"
	"strconv"
	"testing"
)

func TestExistStringEnv(t *testing.T) {
	key := "CUCKMOD_TEST_STRING_ENV"
	value := "foo"
	err := os.Setenv(key, value)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err := os.Unsetenv(key)
		if err != nil {
			t.Error(err)
		}
	}()
	s := "bar"
	StringEnv(&s, key)
	if s != value {
		t.Errorf("StringEnv = %q, want %q", s, value)
	}
}

func TestNotExistStringEnv(t *testing.T) {
	key := "CUCKMOD_TEST_STRING_ENV"
	err := os.Unsetenv(key)
	if err != nil {
		t.Error(err)
	}
	want := "foo"
	s := want
	StringEnv(&s, key)
	if s != want {
		t.Errorf("StringEnv = %q, want %q", s, want)
	}
}

func TestExistIntEnv(t *testing.T) {
	key := "CUCKMOD_TEST_INT_ENV"
	value := 123
	err := os.Setenv(key, strconv.Itoa(value))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err := os.Unsetenv(key)
		if err != nil {
			t.Error(err)
		}
	}()
	i := 321
	IntEnv(&i, key)
	if i != value {
		t.Errorf("StringEnv = %q, want %q", i, value)
	}
}

func TestNotExistIntEnv(t *testing.T) {
	key := "CUCKMOD_TEST_INT_ENV"
	err := os.Unsetenv(key)
	if err != nil {
		t.Error(err)
	}
	want := 123
	i := want
	IntEnv(&i, key)
	if i != want {
		t.Errorf("StringEnv = %q, want %q", i, want)
	}
}
