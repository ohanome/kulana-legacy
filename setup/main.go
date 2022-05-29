package setup

import (
	"github.com/joho/godotenv"
	"kulana/misc"
	"os"
)

var setupDir = os.Getenv("HOME") + "/.kulana"
var envFile = setupDir + "/.env"

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

	if os.Getenv("SMTP_HOST") == "" {
		misc.Die("Missing SMTP_HOST config. Edit the environment file under " + envFile + " and try again.")
	}
	if os.Getenv("SMTP_USERNAME") == "" {
		misc.Die("Missing SMTP_USERNAME config. Edit the environment file under " + envFile + " and try again.")
	}
	if os.Getenv("SMTP_PASSWORD") == "" {
		misc.Die("Missing SMTP_PASSWORD config. Edit the environment file under " + envFile + " and try again.")
	}
	if os.Getenv("SMTP_PORT") == "" {
		misc.Die("Missing SMTP_PORT config. Edit the environment file under " + envFile + " and try again.")
	}
	if os.Getenv("SMTP_ENCRYPTION") == "" {
		misc.Die("Missing SMTP_ENCRYPTION config. Edit the environment file under " + envFile + " and try again.")
	}
	if os.Getenv("SMTP_ADDRESS") == "" {
		misc.Die("Missing SMTP_ADDRESS config. Edit the environment file under " + envFile + " and try again.")
	}
}
