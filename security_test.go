package core

import (
	"bytes"
	"testing"
)

func TestSalt(t *testing.T) {
	// Salt length.
	want := 20
	if salt := Salt(want); len(salt) != want {
		t.Errorf("Salt length = %q, want %q", len(salt), want)
	}

	// Same salt.
	if bytes.Equal(Salt(want), Salt(want)) {
		t.Error("Wrong result for Salt got same result twice")
	}
}

func TestHash(t *testing.T) {
	hash, salt := Hash([]byte("password"))

	if len(hash) != keyLength {
		t.Errorf("Hash length = %q, want %q", len(hash), keyLength)
	}
	if len(salt) != saltLength {
		t.Errorf("Hash length = %q, want %q", len(salt), saltLength)
	}

	hash2, salt2 := Hash([]byte("password"))
	if bytes.Equal(hash, hash2) {
		t.Error("Wrong result for Hash got same hash twice")
	}
	if bytes.Equal(salt, salt2) {
		t.Error("Wrong result for Hash got same salt twice")
	}
}

func TestEqual(t *testing.T) {
	password := []byte("foo")
	hash, salt := Hash(password)

	if equal := Equal(password, hash, salt); !equal {
		t.Errorf("Equal = %t, want %t", equal, true)
	}

	if equal := Equal([]byte("bar"), hash, salt); equal {
		t.Errorf("Equal = %t, want %t", equal, false)
	}

	if equal := Equal(password, hash, Salt(saltLength)); equal {
		t.Errorf("Equal = %t, want %t", equal, false)
	}
}
