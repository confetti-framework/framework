package test

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_mime_from_unknown_extension(t *testing.T) {
	mime, ok := support.MimeByExtension("unknown")
	require.Equal(t, "", mime)
	require.False(t, ok)
}

func Test_mime_from_empty_extension(t *testing.T) {
	mime, ok := support.MimeByExtension("")
	require.Equal(t, "", mime)
	require.False(t, ok)
}

func Test_mime_from_valid_extension(t *testing.T) {
	mime, ok := support.MimeByExtension(".pdf")
	require.Equal(t, "application/pdf", mime)
	require.True(t, ok)
}

func Test_mime_from_valid_filename(t *testing.T) {
	mime, ok := support.MimeByExtension("1234.pdf")
	require.Equal(t, "application/pdf", mime)
	require.True(t, ok)
}
