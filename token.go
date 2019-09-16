package core

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// TokenConfig is a struct with parameters for token generating and verification.
type TokenConfig struct {
	Length  int
	Charset string
}

const (
	CommonCharset = "abcdefghijklmnopqrstuvwxyz0123456789"
	CaseCharset   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	Base64Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

var (
	// TokenIsMissing is a used as a return value from Verify to indicate that the token is missing.
	TokenIsMissing = errors.New("token is missing")

	// WrongTokenLength is a used as a return value from Verify to indicate that the token length not equals to config length.
	WrongTokenLength = errors.New("wrong token length")

	// TokenIsNotValid is a used as a return value from Verify to indicate that the token is missing.
	TokenIsNotValid = errors.New("token is not valid")
)

// Generate function returns generated token from config settings.
func (c TokenConfig) Generate() string {
	bytes := Salt(c.Length)
	for i, b := range bytes {
		bytes[i] = c.Charset[b%byte(len(c.Charset))]
	}
	return string(bytes)
}

// Verify function returns error if token is missing or does not match the config.
func (c TokenConfig) Verify(token string) error {
	if token == "" {
		return TokenIsMissing
	}
	if utf8.RuneCountInString(token) != c.Length {
		return WrongTokenLength
	}
	for _, r := range token {
		if strings.IndexRune(c.Charset, r) == -1 {
			return TokenIsNotValid
		}
	}
	return nil
}
