package rule

import (
	"errors"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/confetti-framework/framework/validation/val"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_accepted_field_not_present(t *testing.T) {
	errs := val.Validate(nil,
		nil,
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 1)
	require.EqualError(t, errs[0], "the title must be present")
}

func Test_accepted_field_present_but_empty_string(t *testing.T) {
	errs := val.Validate(nil,
		map[string]string{"title": ""},
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 1)
	require.True(t, errors.Is(errs[0], rule.MustBeAcceptedError))
}

func Test_accepted_field_present_with_string_yes(t *testing.T) {
	errs := val.Validate(nil,
		map[string]string{"title": "yes"},
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 0)
}

func Test_accepted_field_present_with_string_on(t *testing.T) {
	errs := val.Validate(nil,
		map[string]string{"title": "on"},
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 0)
}
