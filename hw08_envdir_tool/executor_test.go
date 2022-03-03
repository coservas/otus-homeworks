package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("-1", func(t *testing.T) {
		var env Environment
		cmd := []string{
			"not_found_cmd",
			"-somekey",
		}

		code := RunCmd(cmd, env)
		require.Equal(t, -1, code)
	})

	t.Run("0", func(t *testing.T) {
		var env Environment
		cmd := []string{
			"ls",
			"-l",
		}

		code := RunCmd(cmd, env)
		require.Equal(t, 0, code)
	})

	t.Run("1", func(t *testing.T) {
		var env Environment
		cmd := []string{
			"git",
		}

		code := RunCmd(cmd, env)
		require.Equal(t, 1, code)
	})
}
