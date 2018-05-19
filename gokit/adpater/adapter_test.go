package adapter

import (
	"github.com/client9/ilog"
	"github.com/go-kit/kit/log"
	"os"
	"testing"
)

func testLogger(t *testing.T, logger ilog.Logger) {
	if err := logger.Log("foo", "bar"); err != nil {
		t.Fatalf("err=%q", err)
	}
	if err := logger.Log("foo"); err != nil {
		t.Fatalf("err=%q", err)
	}
}
func TestAdapt(t *testing.T) {
	logger := New(log.NewLogfmtLogger(os.Stdout))
	testLogger(t, logger)
}
