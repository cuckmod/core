package core

import (
	"regexp"
	"testing"
	"unicode/utf8"
)

var testConfig = TokenConfig{
	Length:  20,
	Charset: Base64Charset,
}

func TestTokenConfig_Generate(t *testing.T) {
	// Token length.
	if length := utf8.RuneCountInString(testConfig.Generate()); length != testConfig.Length {
		t.Errorf("Generate length = %q, want %q", length, testConfig.Length)
	}

	// Same token.
	if testConfig.Generate() == testConfig.Generate() {
		t.Error("Wrong result for Generate got same result twice")
	}

	// Unexpected character.
	regex := regexp.MustCompile(`[^A-Za-z0-9\-_]`)
	for i := 0; i < 100; i++ {
		if regex.MatchString(testConfig.Generate()) {
			t.Error("Wrong result for Generate got unexpected character")
		}
	}
}

func TestTokenConfig_Verify(t *testing.T) {
	tests := []struct {
		token string
		err   error
	}{
		{
			token: "",
			err:   TokenIsMissing,
		},
		{
			token: "0A1b2C3d4E5f6G7h8I9",
			err:   WrongTokenLength,
		},
		{
			token: "0A1b2C3d4E5f6G7h8I9+",
			err:   TokenIsNotValid,
		},
		{
			token: "0A1b2C3d4E5f6G7h8I9j",
			err:   nil,
		},
	}
	for _, test := range tests {
		t.Run(test.token, func(t *testing.T) {
			if err := testConfig.Verify(test.token); err != test.err {
				t.Errorf("Verify = %q, want %q", err, test.err)
			}
		})
	}
}
