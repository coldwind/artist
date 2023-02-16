package ilog

import (
	"testing"

	"go.uber.org/zap"
)

func TestDebug(t *testing.T) {
	Start("/tmp/", "debug", true, true)
	Debug("debug msg", zap.Int64("value", 12345))
}
