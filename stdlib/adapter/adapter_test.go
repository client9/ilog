package adapter

import (
	"github.com/client9/ilog"
	"log"
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
	logger := New(log.New(os.Stdout, "XX ", 0))
	testLogger(t, logger)
}
