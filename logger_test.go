package core

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevelValue_String(t *testing.T) {
	tests := []struct {
		level  LevelValue
		string string
	}{
		{
			level:  LevelValue(logrus.TraceLevel),
			string: "trace",
		},
		{
			level:  LevelValue(logrus.DebugLevel),
			string: "debug",
		},
		{
			level:  LevelValue(logrus.InfoLevel),
			string: "info",
		},
		{
			level:  LevelValue(logrus.WarnLevel),
			string: "warning",
		},
		{
			level:  LevelValue(logrus.ErrorLevel),
			string: "error",
		},
		{
			level:  LevelValue(logrus.FatalLevel),
			string: "fatal",
		},
		{
			level:  LevelValue(logrus.PanicLevel),
			string: "panic",
		},
	}
	for _, test := range tests {
		t.Run(test.string, func(t *testing.T) {
			if test.level.String() != test.string {
				t.Errorf("String = %q, want %q", test.level.String(), test.string)
			}
		})
	}
}

func TestLevelValue_Set(t *testing.T) {
	tests := []struct {
		level  LevelValue
		string string
	}{
		{
			level:  LevelValue(logrus.TraceLevel),
			string: "trace",
		},
		{
			level:  LevelValue(logrus.DebugLevel),
			string: "debug",
		},
		{
			level:  LevelValue(logrus.InfoLevel),
			string: "info",
		},
		{
			level:  LevelValue(logrus.WarnLevel),
			string: "warn",
		},
		{
			level:  LevelValue(logrus.WarnLevel),
			string: "warning",
		},
		{
			level:  LevelValue(logrus.ErrorLevel),
			string: "error",
		},
		{
			level:  LevelValue(logrus.FatalLevel),
			string: "fatal",
		},
		{
			level:  LevelValue(logrus.PanicLevel),
			string: "panic",
		},
	}
	for _, test := range tests {
		t.Run(test.string, func(t *testing.T) {
			var level LevelValue
			err := level.Set(test.string)
			if err != nil {
				t.Error(err)
			}
			if level != test.level {
				t.Errorf("Set = %q, want %q", level.String(), test.level.String())
			}
		})
	}
}
