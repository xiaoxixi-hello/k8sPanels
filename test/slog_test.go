package test

import (
	"github.com/gookit/slog"
	"testing"
)

func TestSLog(t *testing.T) {
	slog.Warn("测试1")
	slog.Error("测试2")
}
