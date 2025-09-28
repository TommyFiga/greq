package output

import (
	"os"
	"testing"
)

func TestWriteOutput_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "greq_test_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := []byte("Hello File")
	err = WriteResponseContentToFile(string(content), tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error writing file: %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	if string(data) != string(content) {
		t.Errorf("expected %q, got %q", string(content), string(data))
	}
}

func TestWriteOutput_EmptyContent(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "greq_test_empty_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := []byte("")
	err = WriteResponseContentToFile(string(content), tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error writing empty file: %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	if len(data) != 0 {
		t.Errorf("expected empty file, got %d bytes", len(data))
	}
}

func TestWriteOutput_InvalidPath(t *testing.T) {
	content := []byte("Should fail")
	err := WriteResponseContentToFile(string(content), "invalid/dir/test.txt")
	if err == nil {
		t.Fatalf("expected error writing to invalid path, got nil")
	}
}

func TestWriteOutput_OverwriteFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "greq_test_overwrite_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	initial := []byte("Old content")
	if err := os.WriteFile(tmpFile.Name(), initial, 0644); err != nil {
		t.Fatalf("failed to write initial content: %v", err)
	}

	newContent := []byte("New content")
	err = WriteResponseContentToFile(string(newContent), tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error writing file: %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	if string(data) != string(newContent) {
		t.Errorf("expected %q, got %q", string(newContent), string(data))
	}
}