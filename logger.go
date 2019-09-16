package core

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"unicode/utf8"
)

// Log is global variable.
var Log = log.New()

// LevelValue type is logging level that implements Value interface.
type LevelValue log.Level

// CustomFormatter is a struct that implements Formatter interface.
type CustomFormatter struct {
	Service string
}

// InitLogger function initialize logger.
func InitLogger(filename, service string, level LevelValue) {
	Log.SetLevel(log.Level(level))
	Log.SetFormatter(&CustomFormatter{Service: service})

	logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		Log.Panic(err)
	}
	Log.SetOutput(logfile)
}

// Format function returns formatted log entry that implements Formatter interface.
func (f CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	buffer.WriteString(entry.Time.Format("Jan 02 15:04:05.000"))
	buffer.WriteRune(' ')
	buffer.WriteString(strings.ToUpper(entry.Level.String()))

	padding := 28
	buffer.WriteString(strings.Repeat(" ", padding-utf8.RuneCountInString(buffer.String())))

	if f.Service != "" {
		buffer.WriteString(f.Service + ": ")
	}

	buffer.WriteString(entry.Message)

	if len(entry.Data) > 0 {
		if entry.Message != "" {
			buffer.WriteRune('\n')
			buffer.WriteString(strings.Repeat(" ", padding))
		}

		i := 0
		for k, v := range entry.Data {
			if strings.ContainsRune(k, ' ') || strings.ContainsRune(k, '=') {
				buffer.WriteString(fmt.Sprintf(`"%s"`, k))
			}
			buffer.WriteRune('=')
			value := fmt.Sprintf(`%v`, v)
			if strings.ContainsRune(value, ' ') || strings.ContainsRune(value, '=') {
				buffer.WriteString(fmt.Sprintf(`"%v"`, value))
			}
			if i < len(entry.Data)-1 {
				buffer.WriteRune(' ')
			}
			i++
		}
	}
	buffer.WriteRune('\n')
	return buffer.Bytes(), nil
}

func (l LevelValue) String() string {
	return log.Level(l).String()
}

// Set function that implements Value interface.
func (l *LevelValue) Set(s string) error {
	level, err := log.ParseLevel(strings.ToLower(s))
	if err != nil {
		return err
	}
	*l = LevelValue(level)
	return nil
}
