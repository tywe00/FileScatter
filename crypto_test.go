package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewEncryptionKey(t *testing.T) {
	key := newEncryptionKey()
	for i := 0; i < len(key); i++ {
		if key[i] == 0x0 {
			t.Errorf("0 bytes")
		}
	}
}

func TestCopyEncrypt(t *testing.T) {
	src := bytes.NewReader([]byte("Foo not Bar"))
	dst := new(bytes.Buffer)
	key:= newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(dst.Bytes())
}