package test

import (
	"github.com/confetti-framework/support/caller"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_path(t *testing.T) {
	require.Contains(t, caller.Path(), "caller_test.go")
}

func Test_current_dir(t *testing.T) {
	require.Contains(t, caller.CurrentDir(), "/test")
}
