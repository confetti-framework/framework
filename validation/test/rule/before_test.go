package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/confetti-framework/framework/validation/val"
	"github.com/stretchr/testify/require"
	"github.com/uniplaces/carbon"
	"testing"
)

func Test_before_field_not_present(t *testing.T) {
	errs := val.Validate(nil,
		nil,
		val.Verify("start_date", rule.Before{}),
	)
	require.Len(t, errs, 0)
}

func Test_before_field_no_options(t *testing.T) {
	value := support.NewValue("2021")
	err := rule.Before{}.Verify(value)
	require.EqualError(t, err, "can't validate rule.Before: option Date is required")
}

func Test_before_tomorrow(t *testing.T) {
	value := support.NewValue(carbon.Now().String())
	err := rule.Before{Date: carbon.Now().AddDay()}.Verify(value)
	require.Nil(t, err)
}

func Test_before_but_equal(t *testing.T) {
	date := carbon.Now()
	value := support.NewValue(date.String())
	err := rule.Before{Date: date}.Verify(value)
	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be before \d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, err.Error())
}

func Test_before_not_before(t *testing.T) {
	value := support.NewValue(carbon.Now().String())
	err := rule.Before{Date: carbon.Now().SubDay()}.Verify(value)
	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be before \d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, err.Error())
}

func Test_before_with_input_format(t *testing.T) {
	input := carbon.Now()
	input.SetStringFormat(carbon.HourMinuteFormat)
	value := support.NewValue(input.String())

	err := rule.Before{
		Date:   carbon.Now().AddDay(),
		Format: carbon.HourMinuteFormat,
	}.Verify(value)

	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be before \d{2}:\d{2}`, err.Error())
}

func Test_before_with_clear_format(t *testing.T) {
	input := carbon.Now()
	input.SetStringFormat(carbon.HourMinuteFormat)
	value := support.NewValue(input.String())

	err := rule.Before{
		Date:     carbon.Now().AddDay(),
		Format:   "HH:MM",
		TimeZone: "UTC",
	}.Verify(value)

	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be before \d{2}:\d{2}`, err.Error())
}

func Test_before_with_timezone(t *testing.T) {
	input := carbon.Now()
	_ = input.SetTimeZone("UTC")
	value := support.NewValue(input.String())

	err := rule.Before{
		Date:     carbon.Now().SubSeconds(5),
		TimeZone: "UTC",
	}.Verify(value)

	require.NotNil(t, err)
	require.Contains(t, err.Error(), `the :attribute must be before`)
}
