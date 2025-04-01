package sl

import (
	"log/slog"
	"runtime"
	"strconv"
	"strings"
)

const (
	maxStackDepth = 5
)

type StackTracer interface {
	StackTrace() []string
}

func Err(err error, withStack bool) slog.Attr {
	if err == nil {
		return slog.Attr{}
	}

	if !withStack {
		return slog.String("error", err.Error())
	}

	if st, ok := err.(StackTracer); ok {
		return slog.Group("error",
			slog.String("msg", err.Error()),
			slog.Any("stack", st.StackTrace()),
		)
	}

	return slog.Group("error",
		slog.String("msg", err.Error()),
		slog.Any("stack", getStackTrace()),
	)
}

func getStackTrace() []string {
	var stack []string
	for i := 2; i <= maxStackDepth; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stack = append(stack, formatFrame(pc, file, line))
	}
	return stack
}

func formatFrame(pc uintptr, file string, line int) string {
	fn := runtime.FuncForPC(pc)
	return strings.TrimPrefix(fn.Name(), "your_project/") + "() " + file + ":" + strconv.Itoa(line)
}
