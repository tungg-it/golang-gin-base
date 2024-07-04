package logger

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

type MyFormatter struct {
	logrus.TextFormatter // Inherit from TextFormatter
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Define your desired format string
	format := "%s[%s] [%s]: %s\n"

	// Customize message based on entry fields and add colors
	color := ""
	switch entry.Level {
	case logrus.DebugLevel:
		color = "\x1b[36m" // Cyan
	case logrus.InfoLevel:
		color = "\x1b[32m" // Green
	case logrus.WarnLevel:
		color = "\x1b[33m" // Yellow
	case logrus.ErrorLevel:
		color = "\x1b[31m" // Red
	case logrus.FatalLevel:
		color = "\x1b[35m" // Magenta
	}

	// Customize message based on entry fields
	message := fmt.Sprintf(format,
		color,
		strings.ToUpper(entry.Level.String()),
		entry.Time.Format("2006-01-02 15:04:05 MST"),
		entry.Message,
	)

	return []byte(message + "\x1b[0m"), nil
}

func Init() {
	Logger = logrus.New()
	Logger.Formatter = &MyFormatter{}
	// Logger.Formatter = &logrus.TextFormatter{}

	// f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer f.Close()
	// Logger.SetOutput(f)

	Logger.SetLevel(logrus.DebugLevel)
}

func Debug(msg string) {
	Logger.Debug(msg)
}

func Info(msg string) {
	Logger.Info(msg)
}

func Warn(msg string) {
	Logger.Warn(msg)
}

func Error(msg string) {
	Logger.Error(msg)
}

func Fatal(msg string) {
	Logger.Fatal(msg)
}
