package filemode

import (
	"os"
	"testing"
)

const (
	TEST_FILE = "testfile.txt"
	TEST_DIR  = "testdir"
)

func TestMode(t *testing.T) {

	mode := Set(0, S_IFDIR)

	if !IsDir(mode) {
		t.Fatalf("Must be dir\n")
	}

	if mode = Set(mode, S_IRUSR); !IsSet(mode, S_IRUSR) {
		t.Fatalf("Mode must have S_IRUSR\n")
	}

	if mode = Set(mode, S_IRUSR); !IsSet(mode, S_IRUSR) {
		t.Fatalf("Mode must have S_IRUSR\n")
	}
	t.Logf("Mode: %s (%o)\n", mode, mode)

	if mode = Set(mode, S_IRUSR); !IsSet(mode, S_IRUSR) {
		t.Fatalf("Mode must have S_IRUSR\n")
	}

	if mode = Unset(mode, S_IFDIR); IsDir(mode) {
		t.Fatalf("Should not be dir\n")
	}
	t.Logf("Mode: %s (%o)\n", mode, mode)

	if mode = Unset(mode, S_IFDIR); IsDir(mode) {
		t.Fatalf("Should not be dir\n")
	}
	t.Logf("Mode: %s (%o)\n", mode, mode)

}

func TestFile(t *testing.T) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	defer file.Close()

	isDir, err := IsDirFile(file)
	if err != nil {
		t.Fatalf("Failed to check dir: %s\n", err)
	}
	if isDir {
		t.Fatalf("Must not be a dir")
	}

	if err = SetFile(file, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	}

	if isSet, err := IsSetFile(file, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	} else if !isSet {
		t.Fatalf("S_IROTH must be set\n")
	}

	if err = SetFile(file, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	}

	if isSet, err := IsSetFile(file, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	} else if !isSet {
		t.Fatalf("S_IROTH must be set\n")
	}

}

func TestPath(t *testing.T) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	file.Close()

	isDir, err := IsDirPath(TEST_FILE)
	if err != nil {
		t.Fatalf("Failed to check is dir: %s\n", err)
	}
	if isDir {
		t.Fatalf("Must not be a dir")
	}

	if err = SetPath(TEST_FILE, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	}

	if isSet, err := IsSetPath(TEST_FILE, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	} else if !isSet {
		t.Fatalf("S_IROTH must be set\n")
	}

	if err = SetPath(TEST_FILE, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	}

	if isSet, err := IsSetPath(TEST_FILE, S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	} else if !isSet {
		t.Fatalf("S_IROTH must be set\n")
	}

}

func BenchmarkGetFile(b *testing.B) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	defer file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetFile(file)
	}

}

func BenchmarkSetFile(b *testing.B) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	defer file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SetFile(file, S_IXOTH)
	}

}

func BenchmarkGetPath(b *testing.B) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetPath(TEST_FILE)
	}
}

func BenchmarkSetPath(b *testing.B) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SetPath(TEST_FILE, S_IXOTH)
	}
}
