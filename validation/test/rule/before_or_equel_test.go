package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/confetti-framework/framework/validation/val"
	"github.com/stretchr/testify/require"
	"github.com/uniplaces/carbon"
	"testing"
)

func Test_before_or_equal_or_equal_field_not_present(t *testing.T) {
	errs := val.Validate(nil,
		nil,
		val.Verify("start_date", rule.BeforeOrEqual{}),
	)
	require.Len(t, errs, 0)
}

func Test_before_or_equal_field_no_options(t *testing.T) {
	value := support.NewValue("2021")
	err := rule.BeforeOrEqual{}.Verify(value)
	require.EqualError(t, err, "can't validate rule.BeforeOrEqual: option Date is required")
}

func Test_before_or_equal_tomorrow(t *testing.T) {
	value := support.NewValue(carbon.Now().SubDay().String())
	err := rule.BeforeOrEqual{Date: carbon.Now()}.Verify(value)
	require.Nil(t, err)
}

func Test_before_or_equal_with_equal_date(t *testing.T) {
	date := carbon.Now()
	date.SetStringFormat(carbon.DefaultFormat)
	value := support.NewValue(date.String())
	err := rule.BeforeOrEqual{Date: date, Format: carbon.DefaultFormat}.Verify(value)
	require.Nil(t, err)
}

func Test_before_or_equal_not_before(t *testing.T) {
	value := support.NewValue(carbon.Now().AddDay().String())
	err := rule.BeforeOrEqual{Date: carbon.Now()}.Verify(value)
	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be before or equal to \d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, err.Error())
}

func Test_before_or_equal_with_input_format(t *testing.T) {
	input := carbon.Now().AddDay()
	input.SetStringFormat(carbon.DateFormat)
	value := support.NewValue(input.String())

	err := rule.BeforeOrEqual{
		Date:   carbon.Now(),
		Format: carbon.DateFormat,
	}.Verify(value)

	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be before or equal to \d{4}-\d{2}-\d{2}, \d{4}-\d{2}-\d{2} given`, err.Error())
}

func Test_before_or_equal_with_clear_format(t *testing.T) {
	input := carbon.Now().AddMinutes(2)
	input.SetStringFormat(carbon.HourMinuteFormat)
	value := support.NewValue(input.String())

	err := rule.BeforeOrEqual{
		Date:     carbon.Now(),
		Format:   "HH:MM",
		TimeZone: "UTC",
	}.Verify(value)

	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be before or equal to \d{2}:\d{2}, \d{2}:\d{2} given`, err.Error())
}

func Test_before_or_equal_with_timezone(t *testing.T) {
	input := carbon.Now().AddSeconds(5)
	_ = input.SetTimeZone("UTC")
	value := support.NewValue(input.String())

	err := rule.BeforeOrEqual{
		Date:     carbon.Now(),
		TimeZone: "UTC",
	}.Verify(value)

	require.NotNil(t, err)
	require.Contains(t, err.Error(), `the :attribute must be before or equal`)
}
