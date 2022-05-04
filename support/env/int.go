package env

import (
	"github.com/confetti-framework/support"
	"os"
)

func Int(search string) int {
	env, ok := os.LookupEnv(search)
	if !ok {
		panic("env " + search + " not found")
	}

	return support.NewValue(env).Int()
}

func IntOr(search string, def int) int {
	env, ok := os.LookupEnv(search)
	if !ok {
		return def
	}

	result, err := support.NewValue(env).IntE()
	if err != nil {
		return def
	}

	return result
}
