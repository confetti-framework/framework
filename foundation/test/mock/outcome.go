package mock

import (
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/foundation/encoder"
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
