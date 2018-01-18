package vsq

import (
	"os"
	"testing"
)

const testFilePath = "/tmp/vsq_test.json"

func removeTestFile() {
	if _, err := os.Stat(testFilePath); err == nil {
		os.Remove(testFilePath)
	}
}

func TestMain(m *testing.M) {
	exit := m.Run()
	removeTestFile()
	os.Exit(exit)
}
