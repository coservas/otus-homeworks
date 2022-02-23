package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	fromPath := "testdata/input.txt"
	toPath := "/tmp/output.txt"

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", toPath, 0, 0)
		require.EqualError(t, err, ErrUnsupportedFile.Error())
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		err := Copy(fromPath, toPath, 9999, 0)
		require.EqualError(t, err, ErrOffsetExceedsFileSize.Error())
	})
}
