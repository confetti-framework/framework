package env

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func init() {
	defaultEnv := ".env"

	envFile := GetEnvFileByArgs(os.Args, defaultEnv)
	err := godotenv.Load(getFullPath(envFile))
	if err != nil && envFile != defaultEnv {
		println("Can't load environments: " + err.Error())
		os.Exit(1)
	}
}

// GetEnvFileByArgs returns the desired env file
func GetEnvFileByArgs(args []string, def string) string {
	const flag = "--env-file"

	for i, arg := range args {
		if arg == flag && flagHasValue(i, args) {
			return args[i+1]
		}
	}
	return def
}

// flagHasValue checks whether a value is available after the flag
func flagHasValue(i int, args []string) bool {
	return len(args) > i+1
}

func getFullPath(file string) string {
	if strings.HasPrefix(file, "/") {
		return file
	}

	root, _ := os.Getwd()
	return root + "/" + file
}
