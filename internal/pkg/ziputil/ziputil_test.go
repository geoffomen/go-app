package ziputil

import (
	"os"
	"testing"
)

func TestCompressToBuffer(t *testing.T) {
	var files = []string{"../ziputil"}
	b, err := CompressToBuffer(files)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("/tmp/test1.zip", b, 0777)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCompressToFile(t *testing.T) {
	var files = []string{"../ziputil"}
	dest := "/tmp/test2.zip"
	err := CompressToFile(files, dest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeCompress(t *testing.T) {
	err := DeCompress("/tmp/test.zip", "/tmp/de")
	if err != nil {
		t.Fatal(err)
	}
}
