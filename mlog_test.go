package mlog

import (
	"os"
	"testing"
)

func TestDebug(t *testing.T) {
	Start("Debug", "")

	Debug("Debug log")
	Info("info log")
	Warning("warning log")

	Error("error log")

	// Fatalf("fatalf log")
	Stop()
}

func TestInfo(t *testing.T) {
	Start("Info", "")

	Debug("Debug log")
	Info("info log")
	Warning("warning log")

	Error("error log")

	// Fatalf("fatalf log")
	Stop()
}

func TestWarning(t *testing.T) {
	Start("Warn", "")

	Debug("Debug log")
	Info("info log")
	Warning("warning log")

	Error("error log")

	// Fatalf("fatalf log")
	Stop()
}

func TestError(t *testing.T) {
	Start("Error", "")

	Debug("Debug log")
	Info("info log")
	Warning("warning log")

	Error("error log")

	// Fatalf("fatalf log")
	Stop()
}

func TestStartEx(t *testing.T) {
	path := "./test"
	os.RemoveAll(path)

	os.Mkdir(path, 0777)
	fileName := path + "/startex"

	StartEx("Info", fileName, 10, 2)

	Info("Test 1")
	Info("Test 2")

	if _, err := os.Stat(fileName + ".1"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName + ".2"); err == nil {
		t.Fatal(err)
	}

	Info("Test 3")

	if _, err := os.Stat(fileName + ".2"); err != nil {
		t.Fatal(err)
	}

	Info("Test 4")

	if _, err := os.Stat(fileName + ".3"); err == nil {
		t.Fatal(err)
	}

	Stop()

	os.RemoveAll(path)
}

func TestRotatingFileHandler(t *testing.T) {
	path := "./test_log"
	os.RemoveAll(path)

	os.Mkdir(path, 0777)
	fileName := path + "/test"

	h, err := NewRotatingFileHandler(fileName, 10, 2)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 10)

	h.Write(buf)

	h.Write(buf)

	if _, err := os.Stat(fileName + ".1"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName + ".2"); err == nil {
		t.Fatal(err)
	}

	h.Write(buf)
	if _, err := os.Stat(fileName + ".2"); err != nil {
		t.Fatal(err)
	}

	h.Close()

	os.RemoveAll(path)
}
