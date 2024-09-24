package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpictures"
	pathKey := CASPathTransformFunc(key)

	expectedPathName := "b2f8f/1dd50/fdeec/113ac/1b506/6d1d3/a10f7/0f1dc"
	expectedOriginalKey := "b2f8f1dd50fdeec113ac1b5066d1d3a10f70f1dc"

	if pathKey.Pathname != expectedPathName {
		t.Errorf("have %s, want %s", pathKey.Pathname, expectedPathName)
	}
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("have %s, want %s", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	key := "myspecialpicture"
	data := []byte("some jpg bytes")
	
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	
	key := "myspecialpicture"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}

	s.Delete(key)
}
