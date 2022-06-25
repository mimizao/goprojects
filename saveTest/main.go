package main

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"
)

func Save(w io.Writer, data []byte) error {
	return errors.New("text string")
}

func TestSave(t *testing.T) {
	b := make([]byte, 0, 128)
	buf := bytes.NewBuffer(b)
	data := []byte("hello world")
	err := Save(buf, data)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	saved := buf.Bytes()
	if !reflect.DeepEqual(saved, data) {
		t.Errorf("want %s, actual %s", string(data), string(saved))
	}
}

func main() {

}
