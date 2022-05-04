package env

import (
	"golang.org/x/text/language"
	"os"
)

func Lang(search string) language.Tag {
	env, ok := os.LookupEnv(search)
	if !ok {
		panic("Environment '" + search + "' not found")
	}

	return language.Make(env)
}

func LangOr(search string, def language.Tag) language.Tag {
	env, ok := os.LookupEnv(search)
	if !ok {
		return def
	}

	return language.Make(env)
}
