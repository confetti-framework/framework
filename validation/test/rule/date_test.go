package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/confetti-framework/framework/validation/val"
	"github.com/stretchr/testify/require"
	"github.com/uniplaces/carbon"
	"testing"
)

func Test_date_field_not_present(t *testing.T) {
	errs := val.Validate(nil,
		nil,
		val.Verify("start_date", rule.Date{}),
	)
	require.Len(t, errs, 0)
}

func Test_date_field_no_options_and_invalid_date(t *testing.T) {
	value := support.NewValue("2021")
	err := rule.Date{}.Verify(value)
	require.Regexp(t, "the :attribute is not a valid date \\(example \\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\), 2021 given", err.Error())
}

func Test_date_with_valid_default_date(t *testing.T) {
	value := support.NewValue("2021-11-11 15:04:05")
	err := rule.Date{}.Verify(value)
	require.Nil(t, err)
}

func Test_date_with_valid_input_format(t *testing.T) {
	input := carbon.Now()
	input.SetStringFormat("15:04")
	value := support.NewValue(input.String())

	err := rule.Date{
		Format: carbon.HourMinuteFormat,
	}.Verify(value)

	require.Nil(t, err)
}

func Test_date_with_invalid_input_format(t *testing.T) {
	value := support.NewValue("2021-11-11 13:08:05")

	err := rule.Date{
		Format: carbon.HourMinuteFormat,
	}.Verify(value)

	require.NotNil(t, err)
	require.EqualError(t, err, "the :attribute is not a valid date (example 15:04), 2021-11-11 13:08:05 given")
}
