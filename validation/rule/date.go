package rule

import (
	"github.com/confetti-framework/framework/support"
	"github.com/confetti-framework/framework/validation/val_errors"
	"github.com/uniplaces/carbon"
)

type Date struct {
	Format string
}

func (d Date) Verify(value support.Value) error {
	format := normalizeFormat(d.Format)
	_, err := carbon.CreateFromFormat(format, value.String(), "")
	if err != nil {
		return val_errors.WithAttributes(
			DateNotValidFormatError,
			map[string]string{
				"example": format,
				"input":   value.String(),
			},
		)
	}

	return nil
}
