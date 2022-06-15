package setup

import (
	"github.com/joho/godotenv"
	"kulana/_legacy/_misc"
	"os"
)

var setupDir = os.Getenv("HOME") + "/.kulana"
var envFile = setupDir + "/.env"
var configFile = setupDir + "/config.json"
var logFile = setupDir + "/kulana.log"

func GetSetupDir() string {
	return setupDir
}

func GetEnvFile() string {
	return envFile
}

func GetConfigFile() string {
	return configFile
}

func GetLogFile() string {
	return logFile
}

func EnsureEnvironmentIsReady() {
	createSetupDirIfNotExists()
	ensureEnvFileIsReady()
}

func ensureEnvFileIsReady() {
	// Check if environment file ~/.kulana/.env exists
	_, err := os.Stat(envFile)
	if os.IsNotExist(err) {
		defaultEnv := getDefaultEnv()

		fErr := os.WriteFile(envFile, defaultEnv, 0644)
		_misc.Check(fErr)
	}

	err = godotenv.Load(envFile)
	_misc.Check(err)
}

func createSetupDirIfNotExists() {
	// Check if dir ~/.kulana exists
	_, err := os.Stat(setupDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(setupDir, 0755)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
}

func getDefaultEnv() []byte {
	return []byte("" +
		"SMTP_HOST=\n" +
		"SMTP_USERNAME=\n" +
		"SMTP_PASSWORD=\n" +
		"SMTP_PORT=\n" +
		"SMTP_ADDRESS=\n" +
		"")
}
