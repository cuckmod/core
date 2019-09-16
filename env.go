package core

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// StringEnv function set variable to string pointer if environment variable is present.
func StringEnv(v *string, key string) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return
	}
	*v = env
}

// IntEnv function set parsed variable to int pointer if environment variable is present and variable can be parsed.
func IntEnv(v *int, key string) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return
	}
	i, err := strconv.Atoi(env)
	if err != nil {
		Log.WithError(err).WithFields(logrus.Fields{
			"key":      key,
			"variable": env,
		}).Error("can't parse int environment variable")
		return
	}
	*v = i
}
