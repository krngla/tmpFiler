package myTempfiler

import (
	"testing"
)

func TestOpenRead(t *testing.T) {
	_, _ = OpenWrite("test.txt", "Hello World\n")
	output, err := OpenRead("test.txt")
	if err != nil {
		t.Error(err)
	}
	if output != "Hello World\n" {
		t.Errorf("expected: \"Hello World\\n\", got: %q", output)
	}
}

func TestOpenWrite(t *testing.T) {
	n, err := OpenWrite("test.txt", "Hello World\n")
	if err != nil {
		t.Fatal(err)
	}
	if n != 12 {
		t.Error("expected 12 bytes written, got", n)
	}
}

func TestDeleteFile(t *testing.T) {
	_, _ = OpenWrite("test.txt", "")
	err := DeleteFile("test.txt")
	if err != nil {
		t.Error(err)
	}
}
