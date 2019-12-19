package vsq

import (
	"os"
	"path/filepath"
	"testing"
)

func testFilePath() string {
	return filepath.Join(os.TempDir(), "vsq_test.json")
}

func removeTestFile() {
	if _, err := os.Stat(testFilePath()); err == nil {
		os.Remove(testFilePath())
	}
}

func TestMain(m *testing.M) {
	exit := m.Run()
	removeTestFile()
	os.Exit(exit)
}
