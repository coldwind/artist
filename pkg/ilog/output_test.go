package ilog

import (
	"testing"

	"go.uber.org/zap"
)

func TestInfo(t *testing.T) {
	Start("/tmp/", "test.log", true, false)
	Info("test msg", zap.Any("1", "2"))
}
