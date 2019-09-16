package core

import (
	"crypto/rand"
	"crypto/subtle"
	"golang.org/x/crypto/argon2"
)

// Argon2 encoding parameters.
var (
	saltLength = 16
	time       = 2
	memory     = 16 * 1024
	threads    = 1
	keyLength  = 32
)

func init() {
	// Load encoding parameters.
	IntEnv(&saltLength, "CUCKMOD_ARGON2_SALT_LENGTH")
	IntEnv(&time, "CUCKMOD_ARGON2_TIME")
	IntEnv(&memory, "CUCKMOD_ARGON2_MEMORY")
	IntEnv(&threads, "CUCKMOD_ARGON2_THREADS")
	IntEnv(&keyLength, "CUCKMOD_ARGON2_KEY_LENGTH")
}

// Salt function returns the random salt.
func Salt(length int) []byte {
	salt := make([]byte, length)
	_, _ = rand.Read(salt)
	return salt
}

// Hash function returns Argon2 hash of the password and salt.
func Hash(password []byte) (hash, salt []byte) {
	salt = Salt(saltLength)
	return argon2.IDKey(password, salt, uint32(time), uint32(memory), uint8(threads), uint32(keyLength)), salt
}

// Equal function returns true if encoded password and salt equals hash.
func Equal(password, hash, salt []byte) bool {
	otherHash := argon2.IDKey(password, salt, uint32(time), uint32(memory), uint8(threads), uint32(keyLength))
	return subtle.ConstantTimeCompare(hash, otherHash) == 1
}
