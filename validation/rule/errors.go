package rule

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

// Validation Errors

var ValidationError = errors.WithLevel(errors.WithStatus(errors.New(""), net.StatusUnprocessableEntity), log_level.INFO)
var MustBePresentError = errors.Wrap(ValidationError, "the :attribute must be present")
var IsRequiredError = errors.Wrap(ValidationError, "the :attribute is required")
var MustBeAcceptedError = errors.Wrap(ValidationError, "the :attribute must be accepted")
var DateMustBeAfterError = errors.Wrap(ValidationError, "the :attribute must be after :date, :input given")
var DateMustBeAfterOrEqualError = errors.Wrap(ValidationError, "the :attribute must be after or equal to :date, :input given")
var DateMustBeBeforeError = errors.Wrap(ValidationError, "the :attribute must be before :date, :input given")
var DateMustBeBeforeOrEqualError = errors.Wrap(ValidationError, "the :attribute must be before or equal to :date, :input given")
var DateMustBeEqualError = errors.Wrap(ValidationError, "the :attribute must be equal to :date, :input given")
var DateNotValidFormatError = errors.Wrap(ValidationError, "the :attribute is not a valid date (example :example), :input given")
var MuseBeABooleanError = errors.Wrap(ValidationError, "the :attribute must be a boolean, :input given")
var MuseEndWithError = errors.Wrap(ValidationError, "the :attribute must end with :expect, :input given")
var MuseStartWithError = errors.Wrap(ValidationError, "the :attribute must start with :expect, :input given")
var MustHaveAValueError = errors.Wrap(ValidationError, "the :attribute field must have a value")
var SelectedIsInvalidError = errors.Wrap(ValidationError, "the selected :attribute is invalid")
var MustBeAnIntegerError = errors.Wrap(ValidationError, "the :attribute must be an integer")
var MayNotBeGreaterThanError = errors.Wrap(ValidationError, "the :attribute may not be greater than :expect, :input given")
var MayNotHaveMoreThanItemsError = errors.Wrap(ValidationError, "the :attribute may not have more than :expect items, :input items given")
var MustBeAtLeastThanError = errors.Wrap(ValidationError, "the :attribute must be at least :expect, :input given")
var MustBeAtLeastThanItemsError = errors.Wrap(ValidationError, "the :attribute must be at least :expect items, :input items given")
var MustBeError = errors.Wrap(ValidationError, "the :attribute must be :expect, :input given")
var MustBeContainItemsError = errors.Wrap(ValidationError, "the :attribute must contain :expect items, :input items given")
var MustBeASliceError = errors.Wrap(ValidationError, "the :attribute must be a slice")
var MustBeAStringError = errors.Wrap(ValidationError, "the :attribute must be a string")
var MustBeAMapError = errors.Wrap(ValidationError, "the :attribute must be a map")

// System Errors

var OptionDateIsRequiredError = errors.New("option Date is required")
var OptionWithIsRequiredError = errors.New("option With is required")
