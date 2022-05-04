package rule

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/val_errors"
	"github.com/uniplaces/carbon"
)

type After struct {
	Date     *carbon.Carbon
	Format   string
	TimeZone string
}

func (a After) Verify(value support.Value) error {
	format := normalizeFormat(a.Format)
	zone := normalizeZone(a.TimeZone)
	compareTo, err := getComparableDate(a.Date, format, zone)
	if err != nil {
		return errors.Wrap(err, "can't validate rule.After")
	}
	input, err := generateDate(value.String(), format, zone)
	if err != nil {
		return err
	}

	if !input.GreaterThan(compareTo) {
		return val_errors.WithAttributes(
			DateMustBeAfterError,
			map[string]string{"date": compareTo.String(), "input": input.String()},
		)
	}

	return nil
}
