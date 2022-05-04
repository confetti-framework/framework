package test

import (
	"fmt"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/rule"
	"github.com/confetti-framework/framework/validation/val"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/stretchr/testify/require"
	net "net/http"
	"testing"
)

func Test_validate_nothing(t *testing.T) {
	errs := val.Validate(
		nil, support.NewValue(nil))
	require.Equal(t, []error{}, errs)
}

func Test_validate_nothing_with_empty_verification(t *testing.T) {
	errs := val.Validate(
		nil,
		support.NewValue(nil),
		val.Verify("title"),
	)
	require.Empty(t, errs)
}

func Test_validate_nothing_with_empty_verifications(t *testing.T) {
	errs := val.Validate(
		nil,
		support.NewValue(nil),
		val.Verify("title"),
		val.Verify("description"),
	)
	require.Empty(t, errs)
}

func Test_validate_with_multiple_verifications(t *testing.T) {
	errs := val.Validate(
		nil,
		support.NewValue(map[string]string{"title": "Horse", "description": "Big animal"}),
		val.Verify("title"),
		val.Verify("description"),
	)
	require.Empty(t, errs)
}

func Test_validate_with_multiple_invalid_keys(t *testing.T) {
	errs := val.Validate(
		nil,
		support.NewValue(map[string]string{}),
		val.Verify("title", rule.Required{}),
		val.Verify("description", rule.Required{}),
	)
	require.Len(t, errs, 2)
}

func Test_validate_invalid_values_with_multiple_rules(t *testing.T) {
	errs := val.Validate(
		nil,
		support.NewValue(nil),
		val.Verify("title", rule.Present{}, rule.Required{}),
	)
	require.Len(t, errs, 1)
	require.EqualError(t, errs[0], "the title must be present")
}

func Test_validate_nested_key_error(t *testing.T) {
	errs := val.Validate(
		nil,
		support.NewValue(map[string]string{}),
		val.Verify("user.title", rule.Present{}),
	)
	require.EqualError(t, errs[0], "the user.title must be present")
}

func Test_validate_map(t *testing.T) {
	errs := val.Validate(
		nil,
		map[string]string{},
		val.Verify("user.title", rule.Present{}),
	)
	require.EqualError(t, errs[0], "the user.title must be present")
}

func Test_validate_gives_non_validation_error(t *testing.T) {
	errs := val.Validate(
		nil,
		map[string]string{"user": "Jip"},
		val.Verify("user", mockRuleWithNonValidationError{}),
	)
	require.EqualError(t, errs[0], "failed to validate user with test.mockRuleWithNonValidationError: option Date is required")
}

func Test_error_has_stack_trace(t *testing.T) {
	errs := val.Validate(
		nil,
		map[string]string{},
		val.Verify("user.title", rule.Present{}),
	)
	stack, ok := errors.FindStack(errs[0])
	require.True(t, ok)
	require.Contains(t, fmt.Sprintf("%+v", stack), "validator_test.go")
}

func Test_normal_rule_not_required(t *testing.T) {
	errs := val.Validate(
		nil,
		nil,
		val.Verify("title", mockRuleNotRequired{}),
	)
	require.Empty(t, errs)
}

func Test_validation_error_status(t *testing.T) {
	errs := val.Validate(
		nil,
		map[string]string{},
		val.Verify("user.title", rule.Present{}),
	)
	status, _ := errors.FindStatus(errs[0])
	require.Equal(t, net.StatusUnprocessableEntity, status)
}

func Test_validation_log_level(t *testing.T) {
	errs := val.Validate(
		nil,
		map[string]string{},
		val.Verify("user.title", rule.Present{}),
	)
	level, _ := errors.FindLevel(errs[0])
	require.Equal(t, log_level.INFO, level)
}

func Test_validation_with_application(t *testing.T) {
	errs := val.Validate(
		mockApp{},
		map[string]int{"title": 12},
		val.Verify("title", mockRuleApplicationNeeded{}),
	)
	require.Empty(t, errs)
}

// Mock rule not required
type mockRuleNotRequired struct{}

func (m mockRuleNotRequired) Verify(value support.Value) error {
	return errors.New("don't show this error if value not present")
}

// Mock rule with non validation error
type mockRuleWithNonValidationError struct{}

func (m mockRuleWithNonValidationError) Verify(value support.Value) error {
	return errors.New("option Date is required")
}

// Mock rule where application is needed
type mockRuleApplicationNeeded struct {
	app inter.AppReader
}

func (m mockRuleApplicationNeeded) SetApp(app inter.AppReader) inter.Rule {
	m.app = app
	return m
}

func (m mockRuleApplicationNeeded) Verify(value support.Value) error {
	_, err := m.app.MakeE("the_value")
	return err
}

// A mocked application reader
type mockApp struct{}

func (a mockApp) Make(abstract interface{}) interface{} {
	return "The horse"
}

func (a mockApp) MakeE(_ interface{}) (interface{}, error) {
	return "The horse", nil
}
