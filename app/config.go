package app

import (
	"fmt"
	"kulana/misc"
	"kulana/setup"
	"os"
	"reflect"
	"strings"
)

func processConfigArgs(application Application) Application {
	l := len(os.Args)
	if l == 2 {
		misc.Usage(CommandConfig)
		os.Exit(1)
	}

	if l == 3 && (os.Args[2] == "--help" || os.Args[2] == "-h") {
		code := 1
		if os.Args[2] == "--help" || os.Args[2] == "-h" {
			code = 0
		}

		misc.Usage(CommandConfig)
		os.Exit(code)
	}

	return application
}

func runConfig(application Application) {
	l := len(os.Args)
	minlength := 3

	switch os.Args[2] {
	case "get":
		// kulana, config, get, key
		minlength = 4
		break

	case "set":
		// kulana, config, set, key, value
		minlength = 5
		break
	}

	getConfig("foo")

	if l < minlength {
		misc.Die("Not enough arguments.")
	}

	fmt.Println(l)
}

// WIP
func getConfig(key string) any {
	config := setup.ReadConfig()
	PrintFields(config.Mail)

	return true
}

func setConfig(key string, value any) {

}

// PrintFields
// @see https://stackoverflow.com/questions/40864840/how-to-get-the-json-field-names-of-a-struct-in-golang
func PrintFields(b any) {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name

		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
			fmt.Println(fieldName)
		default:
			parts := strings.Split(jsonTag, ",")
			name := parts[0]
			if name == "" {
				name = fieldName
			}
			fmt.Println(name)
		}
	}
}
