// Copyright 2016 Sam Whited.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package internal

import (
	"errors"
	"testing"
)

func BenchmarkRandomIDEven(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomID(8)
	}
}

func BenchmarkRandomIDOdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomID(9)
	}
}

func TestRandomIDLength(t *testing.T) {
	for i := 0; i <= 15; i++ {
		if s := RandomID(i); len(s) != i {
			t.Logf("Expected length %d got %d", i, len(s))
			t.Fail()
		}
	}
}

type errorReader struct{}

func (errorReader) Read(p []byte) (int, error) {
	return 0, errors.New("Expected error from error reader")
}

type nopReader struct{}

func (nopReader) Read(p []byte) (int, error) {
	return 0, nil
}

func TestRandomPanicsIfRandReadFails(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected randomID to panic if reading random bytes failed")
		}
	}()
	randomID(16, errorReader{})
}

func TestRandomPanicsIfRandReadWrongLen(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected randomID to panic if no random bytes were read")
		}
	}()
	randomID(16, nopReader{})
}
