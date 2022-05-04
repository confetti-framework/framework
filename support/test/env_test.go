package test

import (
	"github.com/confetti-framework/support/env"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_env_file_without_args(t *testing.T) {
	require.Equal(t, "", env.GetEnvFileByArgs([]string{}, ""))
}

func Test_env_file_with_env_file(t *testing.T) {
	require.Equal(t, ".env.testing", env.GetEnvFileByArgs([]string{"--env-file", ".env.testing"}, ""))
}

func Test_env_file_without_correct_flag(t *testing.T) {
	require.Equal(t, "", env.GetEnvFileByArgs([]string{"--env", ".env.testing"}, ""))
}

func Test_env_file_in_as_second_flag(t *testing.T) {
	require.Equal(t, ".env.testing", env.GetEnvFileByArgs([]string{"--dry-run", "--env-file", ".env.testing"}, ""))
}

func Test_env_file_in_as_second_flag_without_value(t *testing.T) {
	require.Equal(t, "", env.GetEnvFileByArgs([]string{"--dry-run", "--env-file"}, ""))
}

func Test_env_file_with_default_value(t *testing.T) {
	require.Equal(t, ".env", env.GetEnvFileByArgs([]string{}, ".env"))
}