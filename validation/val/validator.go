package val

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/support"
	rules "github.com/confetti-framework/framework/validation/rule"
	"github.com/confetti-framework/framework/validation/val_errors"
)

func Validate(app inter.AppReader, input interface{}, verifications ...Verification) []error {
	result := []error{}
	value := support.NewValue(input)
	if value.Error() != nil {
		return append(result, value.Error())
	}
	for _, verification := range verifications {
		verification.app = app
		result = append(result, verifyVerification(value, verification)...)
	}
	return result

}

func verifyVerification(input support.Value, verification Verification) []error {
	var result []error

	// If the key contains an asterisk, then there are more keys we need to verify
	fields := support.GetSearchableKeysByOneKey(verification.Field, input)
	for _, field := range fields {
		value, err := input.GetE(field)
		present := err == nil

		for _, rule := range getAllRequiredRules(verification) {
			err := verifyRule(present, value, rule)
			if err != nil {
				result = append(result, decorateErr(err, field, rule))
				break
			}
		}
	}

	return result
}

func decorateErr(err error, field support.Key, rule inter.Rule) error {
	if !errors.Is(err, rules.ValidationError) {
		err = errors.WithMessage(err, "failed to validate :attribute with %s", support.Name(rule))
	}

	err = errors.WithStack(err)
	return val_errors.WithAttribute(err, "attribute", field)
}

func verifyRule(present bool, value support.Value, rule inter.Rule) error {
	if !needToVerify(present, rule) {
		return nil
	}

	return rule.Verify(value)
}

func getAllRequiredRules(verification Verification) []inter.Rule {
	var result []inter.Rule
	for _, baseRule := range verification.Rules {
		if baseRule, ok := baseRule.(inter.RuleWithRequirements); ok {
			result = append(result, baseRule.Requirements()...)
		}

		result = append(result, baseRule)
	}
	result = setAppOnRules(result, verification.app)

	return result
}

func setAppOnRules(rules []inter.Rule, app inter.AppReader) []inter.Rule {
	var result []inter.Rule
	for _, rule := range rules {
		if withApp, ok := rule.(inter.RuleWithApp); ok {
			rule = withApp.SetApp(app)
		}
		result = append(result, rule)
	}

	return result
}

func needToVerify(present bool, rule inter.Rule) bool {
	_, isPresentRule := rule.(rules.Present)

	// If the field is neither present nor required to be present,
	// we do not need to validate further
	if !isPresentRule && !present {
		return false
	}

	// If we only need to check if the value is present, and the value
	// is present, we do not need to validate further
	if isPresentRule && present {
		return false
	}

	return true
}
