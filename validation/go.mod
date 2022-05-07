module github.com/confetti-framework/framework/validation

go 1.17

require (
	github.com/confetti-framework/framework/contract v0.29.0
	github.com/confetti-framework/errors v0.11.0-rc.1
	github.com/confetti-framework/framework/support v0.29.0
	github.com/confetti-framework/syslog v0.1.0-rc
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.6.1
	github.com/uniplaces/carbon v0.1.6
	github.com/vigneshuvi/GoDateFormat v0.0.0-20190923034126-379ee8a8c45f
)

replace (
	github.com/confetti-framework/framework/contract v0.29.0 => ../contract
    github.com/confetti-framework/framework/support v0.29.0 => ../support
)