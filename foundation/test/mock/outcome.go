package mock

import (
	"github.com/confetti-framework/framework/foundation/encoder"
	"github.com/confetti-framework/framework/inter"
)

var JsonEncoders = []inter.Encoder{
	encoder.JsonReaderToJson{},
	encoder.RawToJson{},
	encoder.JsonToJson{},
	encoder.ErrorsToJson{},
	encoder.InterfaceToJson{},
}

var HtmlEncoders = []inter.Encoder{
	encoder.StringerToHtml{},
	encoder.RawToHtml{},
	encoder.InterfaceToHtml{},
}
