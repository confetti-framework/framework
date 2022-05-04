package rule

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

// Validation Error
var ValidationError = errors.New("").Status(net.StatusUnprocessableEntity).Level(log_level.INFO)
var MustBePresentError = ValidationError.Wrap("the :attribute must be present")
var IsRequiredError = ValidationError.Wrap("the :attribute is required")
var MustBeAcceptedError = ValidationError.Wrap("the :attribute must be accepted")
var DateMustBeAfterError = ValidationError.Wrap("the :attribute must be after :date, :input given")
var DateMustBeAfterOrEqualError = ValidationError.Wrap("the :attribute must be after or equal to :date, :input given")
var DateMustBeBeforeError = ValidationError.Wrap("the :attribute must be before :date, :input given")
var DateMustBeBeforeOrEqualError = ValidationError.Wrap("the :attribute must be before or equal to :date, :input given")
var DateMustBeEqualError = ValidationError.Wrap("the :attribute must be equal to :date, :input given")
var DateNotValidFormatError = ValidationError.Wrap("the :attribute is not a valid date (example :example), :input given")
var MuseBeABooleanError = ValidationError.Wrap("the :attribute must be a boolean, :input given")
var MuseEndWithError = ValidationError.Wrap("the :attribute must end with :expect, :input given")
var MuseStartWithError = ValidationError.Wrap("the :attribute must start with :expect, :input given")
var MustHaveAValueError = ValidationError.Wrap("the :attribute field must have a value")
var SelectedIsInvalidError = ValidationError.Wrap("the selected :attribute is invalid")
var MustBeAnIntegerError = ValidationError.Wrap("the :attribute must be an integer")
var MayNotBeGreaterThanError = ValidationError.Wrap("the :attribute may not be greater than :expect, :input given")
var MayNotHaveMoreThanItemsError = ValidationError.Wrap("the :attribute may not have more than :expect items, :input items given")
var MustBeAtLeastThanError = ValidationError.Wrap("the :attribute must be at least :expect, :input given")
var MustBeAtLeastThanItemsError = ValidationError.Wrap("the :attribute must be at least :expect items, :input items given")
var MustBeError = ValidationError.Wrap("the :attribute must be :expect, :input given")
var MustBeContainItemsError = ValidationError.Wrap("the :attribute must contain :expect items, :input items given")
var MustBeASliceError = ValidationError.Wrap("the :attribute must be a slice")
var MustBeAStringError = ValidationError.Wrap("the :attribute must be a string")
var MustBeAMapError = ValidationError.Wrap("the :attribute must be a map")

// System Error
var OptionDateIsRequiredError = errors.New("option Date is required")
var OptionWithIsRequiredError = errors.New("option With is required")
