package rule

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/val_errors"
	"github.com/uniplaces/carbon"
)

type DateEqual struct {
	Date     *carbon.Carbon
	Format   string
	TimeZone string
}

func (b DateEqual) Verify(value support.Value) error {
	format := normalizeFormat(b.Format)
	zone := normalizeZone(b.TimeZone)
	compareTo, err := getComparableDate(b.Date, format, zone)
	if err != nil {
		return errors.Wrap(err, "can't validate rule.DateEqual")
	}
	input, err := generateDate(value.String(), format, zone)
	if err != nil {
		return err
	}

	if !input.EqualTo(compareTo) {
		return val_errors.WithAttributes(
			DateMustBeEqualError,
			map[string]string{"date": compareTo.String(), "input": input.String()},
		)
	}

	return nil
}
