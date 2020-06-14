package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/airabinovich/memequotes_back/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

//SupportLogger is a type for a log with tags
type SupportLogger struct {
	*logrus.Entry
}

var log = logrus.Logger{
	Out: os.Stdout,
	Formatter: &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000-07:00",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
		},
	},
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
}

//NewLogger creates a creditsLogger instance
func NewLogger(tags map[string]interface{}) *SupportLogger {
	c := config.Conf
	level, err := logrus.ParseLevel(c.GetString("log.level", logrus.DebugLevel.String()))
	if err != nil {
		log.Println("Malformed log level configuration. Use configuration by default logging level INFO", err)
		level = logrus.DebugLevel
	}

	log.SetLevel(level)
	sl := &SupportLogger{log.WithFields(tags)}
	sl.Logger.SetOutput(&lumberjack.Logger{
		Filename: c.GetString("log.file_path"),
		MaxAge:   int(c.GetInt32("log.max_age", 1)),
	})
	return sl
}

// Print logs an interface, the related tags, and a formatted stack trace of the goroutine that calls it.
func (log *SupportLogger) Print(e interface{}) {
	log.Printf("%s: %s", e, debug.Stack())
}

// Debug logs a message and the related tags, with a DEBUG level.
func (log *SupportLogger) Debug(message string, tags ...string) {
	log.WithFields(fieldsFromString(tags...)).Debug(message)
}

// Info logs a message and the related tags, with an INFO level.
func (log *SupportLogger) Info(message string, tags ...string) {
	log.WithFields(fieldsFromString(tags...)).Info(message)
}

// Warn logs a message and the related tags, with a WARN level.
func (log *SupportLogger) Warn(message string, tags ...string) {
	log.WithFields(fieldsFromString(tags...)).Warn(message)
}

// Error logs a message and the related tags, with an ERROR level.
func (log *SupportLogger) Error(message string, err error, tags ...string) {
	message = fmt.Sprintf("%s - ERROR: %v", message, err)
	log.WithFields(fieldsFromString(tags...)).Error(message)
}

// Panic logs a message and the related tags, with an ERROR level. Then panics with the same error
func (log *SupportLogger) Panic(message string, err error, tags ...string) {
	message = fmt.Sprintf("%s - ERROR: %v", message, err)
	log.WithFields(fieldsFromString(tags...)).Error(message)
	panic(err)
}

func fieldsFromString(tags ...string) logrus.Fields {
	fields := make(logrus.Fields)

	for _, value := range tags {
		values := strings.Split(value, ":")

		fields[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
	}

	return fields
}

