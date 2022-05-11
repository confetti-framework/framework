package console

import "github.com/confetti-framework/framework/inter"

var FoundationCommands = []inter.Command{
	AppInfo{},
	AppServe{},
	Baker{},
	LogClear{},
	RouteList{},
}
