package setup

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"kulana/misc"
	"os"
)

type Config struct {
	Mail MailConfig `json:"mail"`
}

type MailConfig struct {
	StatusCodes  string `json:"status_codes"`
	Subject      string `json:"subject"`
	TemplateFile string `json:"template_file"`
}

var setupDir = os.Getenv("HOME") + "/.kulana"
var envFile = setupDir + "/.env"
var configFile = setupDir + "/config.json"

var defaultConfig = Config{
	Mail: MailConfig{
		StatusCodes:  "4xx,5xx",
		Subject:      "Host %s is %s",
		TemplateFile: setupDir + "/mail.html",
	},
}

func getSetupDir() string {
	return setupDir
}

func GetEnvFile() string {
	return envFile
}

func EnsureEnvironmentIsReady() {
	createSetupDirIfNotExists()
	ensureEnvFileIsReady()
	ensureConfigFileIsReady()
}

func ensureConfigFileIsReady() {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		defaultConfig := getDefaultConfig()
		defaultConfigJson, jErr := json.Marshal(defaultConfig)
		misc.Check(jErr)

		fErr := os.WriteFile(configFile, defaultConfigJson, 0644)
		misc.Check(fErr)
	}
}

func ensureEnvFileIsReady() {
	// Check if environment file ~/.kulana/.env exists
	_, err := os.Stat(envFile)
	if os.IsNotExist(err) {
		defaultEnv := getDefaultEnv()

		fErr := os.WriteFile(envFile, defaultEnv, 0644)
		misc.Check(fErr)
	}

	err = godotenv.Load(envFile)
	misc.Check(err)
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

func getDefaultConfig() Config {
	return defaultConfig
}

func ReadConfig() Config {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		// Config file should exist at this point.
		panic(err)
	}

	configJsonFile, fErr := os.Open(configFile)
	misc.Check(fErr)
	defer func(configJsonFile *os.File) {
		err := configJsonFile.Close()
		misc.Check(err)
	}(configJsonFile)

	byteValue, _ := ioutil.ReadAll(configJsonFile)
	var config Config

	jErr := json.Unmarshal(byteValue, &config)
	misc.Check(jErr)

	return config
}

func WriteConfig(config Config) {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		// Config file should exist at this point.
		panic(err)
	}

	configJson, jErr := json.Marshal(config)
	misc.Check(jErr)
	
	fErr := os.WriteFile(configFile, configJson, 0644)
	misc.Check(fErr)
}
