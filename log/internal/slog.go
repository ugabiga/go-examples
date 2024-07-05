package internal

import (
	"errors"
	"log/slog"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	slogformatter "github.com/samber/slog-formatter"
)

const (
	LogAttrKeyError = "error"
)

func RunSlog() {
	logger := newSlogLogger()
	logger.Debug("Debug message",
		slog.String("string", "value"),
		slog.Int("int", 1),
		slog.Bool("bool", true),
	)

	err := errors.New("error message")
	logger.Error("Error message", slog.Any(LogAttrKeyError, err))

}

func newSlogLogger() *slog.Logger {
	logOptions := slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &logOptions)

	logger := slog.New(
		slogformatter.NewFormatterHandler(
			errorFormatter(LogAttrKeyError),
		)(
			logHandler,
		),
	)

	return logger
}

func errorFormatter(fieldName string) slogformatter.Formatter {
	return slogformatter.FormatByFieldType(fieldName, func(err error) slog.Value {
		return slog.GroupValue([]slog.Attr{
			slog.String("message", err.Error()),
			slog.String("type", reflect.TypeOf(err).String()),
			slog.Any("stackTrace", stackTrace()),
		}...)
	})
}

type stackTraceInfo struct {
	Function string `json:"function"`
	File     string `json:"file"`
}

var reStackTrace = regexp.MustCompile(`log/slog.*\n`)

func removeAfterPlusZeroX(input string) string {
	idx := strings.Index(input, "+0x")
	if idx != -1 {
		return input[:idx]
	}
	return input
}

func stackTrace() []stackTraceInfo {
	seperatorToken := "---"
	var st []stackTraceInfo

	traceLine := stackTraceLine()
	if traceLine == "" {
		return nil
	}

	traceLine = strings.Replace(traceLine, "\n\t", seperatorToken, -1)

	traceLines := strings.Split(traceLine, "\n")
	for _, line := range traceLines {
		line = removeAfterPlusZeroX(line)
		line = strings.Replace(line, " ", "", -1)

		t := strings.Split(line, seperatorToken)
		if len(t) > 1 {
			st = append(st, stackTraceInfo{
				Function: t[0],
				File:     t[1],
			})
		}
	}
	return st
}

func stackTraceLine() string {
	stackInfo := make([]byte, 1024*1024)

	if stackSize := runtime.Stack(stackInfo, false); stackSize > 0 {
		traceLines := reStackTrace.Split(string(stackInfo[:stackSize]), -1)
		if len(traceLines) > 0 {
			return traceLines[len(traceLines)-1]
		}
	}

	return ""
}
