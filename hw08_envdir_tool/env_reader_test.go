package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("empty dir", func(t *testing.T) {
		envs, err := ReadDir("")
		require.Nil(t, envs)
		require.EqualError(t, err, "directory path is empty")
	})

	t.Run("not found dir", func(t *testing.T) {
		envs, err := ReadDir("/not_existed_dir")
		require.Nil(t, envs)
		require.True(t, os.IsNotExist(err))
	})

	t.Run("success read", func(t *testing.T) {
		expected := Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY": EnvValue{Value: "", NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		}

		actual, err := ReadDir("testdata/env")
		require.Nil(t, err)
		require.Equal(t, expected, actual)
	})
}
