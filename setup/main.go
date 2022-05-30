package setup

import (
	"github.com/joho/godotenv"
	"kulana/misc"
	"os"
)

var setupDir = os.Getenv("HOME") + "/.kulana"
var envFile = setupDir + "/.env"

func getSetupDir() string {
	return setupDir
}

func GetEnvFile() string {
	return envFile
}

func EnsureEnvironmentIsReady() {
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

	// Check if environment file ~/.kulana/.env exists
	_, err = os.Stat(envFile)
	if os.IsNotExist(err) {
		defaultEnv := []byte("" +
			"SMTP_HOST=\n" +
			"SMTP_USERNAME=\n" +
			"SMTP_PASSWORD=\n" +
			"SMTP_PORT=\n" +
			"SMTP_ADDRESS=\n" +
			"")

		fErr := os.WriteFile(envFile, defaultEnv, 0644)
		misc.Check(fErr)
	}

	err = godotenv.Load(envFile)
	misc.Check(err)
}
