package rule

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/val_errors"
	"github.com/uniplaces/carbon"
)

type BeforeOrEqual struct {
	Date     *carbon.Carbon
	Format   string
	TimeZone string
}

func (b BeforeOrEqual) Verify(value support.Value) error {
	format := normalizeFormat(b.Format)
	zone := normalizeZone(b.TimeZone)
	compareTo, err := getComparableDate(b.Date, format, zone)
	if err != nil {
		return errors.Wrap(err, "can't validate rule.BeforeOrEqual")
	}
	input, err := generateDate(value.String(), format, zone)
	if err != nil {
		return err
	}

	if !input.LessThanOrEqualTo(compareTo) {
		return val_errors.WithAttributes(
			DateMustBeBeforeOrEqualError,
			map[string]string{"date": compareTo.String(), "input": input.String()},
		)
	}

	return nil
}
