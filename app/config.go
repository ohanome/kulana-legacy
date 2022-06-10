package app

import (
	"fmt"
	"kulana/misc"
	"kulana/setup"
	"os"
	"reflect"
	"strings"
)

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

	getConfig("mail.subject")

	if l < minlength {
		misc.Die("Not enough arguments.")
	}

	fmt.Println(l)
}

// WIP
func getConfig(key string) any {
	fmt.Println(key)
	config := setup.ReadConfig()
	PrintFields(config)

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
		fmt.Printf("fieldName: %s\n", fieldName)
		fmt.Printf("t: %v\n", t)

		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
			fmt.Printf("fieldName: %s\n", fieldName)
		default:
			parts := strings.Split(jsonTag, ",")
			name := parts[0]
			if name == "" {
				name = fieldName
			}
			fmt.Printf("name: %s\n", name)
			fmt.Printf("jsonTag: %s\n", jsonTag)
		}
	}
}
