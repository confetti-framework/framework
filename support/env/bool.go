package env

import (
	"github.com/confetti-framework/support"
	"os"
)

func Bool(search string) bool {
	env, ok := os.LookupEnv(search)
	if !ok {
		panic("Environment '" + search + "' not found")
	}

	return support.NewValue(env).Bool()
}

func BoolOr(search string, def bool) bool {
	env, ok := os.LookupEnv(search)
	if !ok {
		return def
	}

	return support.NewValue(env).Bool()
}
